package main

import (
	"context"
	"fmt"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"

	"github.com/akhenakh/sataas"
	"github.com/akhenakh/sataas/satsvc"
)

const appName = "sataas"

var (
	version = "no version from LDFLAGS"

	grpcPort        = flag.Int("grpcPort", 9200, "gRPC API port")
	healthPort      = flag.Int("healthPort", 6666, "grpc health port")
	httpMetricsPort = flag.Int("httpMetricsPort", 8088, "http port")

	tleURL = flag.String(
		"tleURL",
		"https://celestrak.com/NORAD/elements/active.txt",
		"default URL to fetch")
	logLevel = flag.String("logLevel", "INFO", "DEBUG|INFO|WARN|ERROR")

	grpcServer        *grpc.Server
	grpcHealthServer  *grpc.Server
	httpMetricsServer *http.Server
)

func main() {
	flag.Parse()

	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "caller", log.Caller(5), "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "app", appName)
	logger = NewLevelFilterFromString(logger, *logLevel)

	stdlog.SetOutput(log.NewStdlibAdapter(logger))

	level.Info(logger).Log("msg", "Starting app", "version", version)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// catch termination
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	// gRPC Health Server
	healthServer := health.NewServer()
	g.Go(func() error {
		grpcHealthServer = grpc.NewServer()

		healthpb.RegisterHealthServer(grpcHealthServer, healthServer)

		haddr := fmt.Sprintf(":%d", *healthPort)
		hln, err := net.Listen("tcp", haddr)
		if err != nil {
			level.Error(logger).Log("msg", "gRPC Health server: failed to listen", "error", err)
			os.Exit(2)
		}
		level.Info(logger).Log("msg", fmt.Sprintf("gRPC health server listening at %s", haddr))
		return grpcHealthServer.Serve(hln)
	})

	// web server metrics
	g.Go(func() error {
		httpMetricsServer = &http.Server{
			Addr:         fmt.Sprintf(":%d", *httpMetricsPort),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		level.Info(logger).Log("msg", fmt.Sprintf("HTTP Metrics server listening at :%d", *httpMetricsPort))

		versionGauge.WithLabelValues(version).Add(1)

		// Register Prometheus metrics handler.
		http.Handle("/metrics", promhttp.Handler())

		if err := httpMetricsServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	// our sataas service
	s := sataas.New(ctx, logger, healthServer, *tleURL)

	// gRPC Server
	g.Go(func() error {
		addr := fmt.Sprintf(":%d", *grpcPort)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			level.Error(logger).Log("msg", "gRPC server: failed to listen", "error", err)
			os.Exit(2)
		}

		grpcServer = grpc.NewServer(
			// MaxConnectionAge is just to avoid long connection, to facilitate load balancing
			// MaxConnectionAgeGrace will torn them, default to infinity
			grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_opentracing.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
			)),
		)
		satsvc.RegisterPredictionServer(grpcServer, s)
		level.Info(logger).Log("msg", fmt.Sprintf("gRPC server serving at %s", addr))

		healthServer.SetServingStatus(fmt.Sprintf("grpc.health.v1.%s", appName), healthpb.HealthCheckResponse_SERVING)

		return grpcServer.Serve(ln)
	})

	err := s.UpdateTLEs()
	if err != nil {
		level.Error(logger).Log("msg", "can't update TLEs", "error", err)
	}

	healthServer.SetServingStatus(fmt.Sprintf("grpc.health.v1.%s", appName), healthpb.HealthCheckResponse_SERVING)
	level.Info(logger).Log("msg", "serving status to SERVING")

	select {
	case <-interrupt:
		cancel()
		break
	case <-ctx.Done():
		break
	}

	level.Warn(logger).Log("msg", "received shutdown signal")

	healthServer.SetServingStatus(fmt.Sprintf("grpc.health.v1.%s", appName), healthpb.HealthCheckResponse_NOT_SERVING)

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	if httpMetricsServer != nil {
		_ = httpMetricsServer.Shutdown(shutdownCtx)
	}

	if grpcHealthServer != nil {
		grpcHealthServer.GracefulStop()
	}

	err = g.Wait()
	if err != nil {
		level.Error(logger).Log("msg", "server returning an error", "error", err)
		os.Exit(2)
	}
}

// NewLevelFilterFromString filter the log level using the string "DEBUG|INFO|WARN|ERROR".
func NewLevelFilterFromString(next log.Logger, ls string) log.Logger {
	switch strings.ToLower(ls) {
	case "debug":
		return level.NewFilter(next, level.AllowDebug())
	case "info":
		return level.NewFilter(next, level.AllowInfo())
	case "warn", "warning":
		return level.NewFilter(next, level.AllowWarn())
	case "error", "err":
		return level.NewFilter(next, level.AllowError())
	}

	return level.NewFilter(next, level.AllowAll())
}
