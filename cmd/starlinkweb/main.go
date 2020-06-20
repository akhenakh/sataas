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
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/akhenakh/sataas/satsvc"
)

const appName = "starlinkweb"

var (
	version = "no version from LDFLAGS"

	sataasURI     = flag.String("sataasURI", "localhost:9200", "sataas grpc URI")
	selfHostedMap = flag.Bool("selfHostedMap", false, "Use a self hosted map rather than MapBox")
	tilesKey      = flag.String("tilesKey", "", "The key that will passed in the queries to the tiles server")
	tilesURL      = flag.String(
		"tilesURL",
		"http://127.0.0.1:8081",
		"the URL where to point to get tiles",
	)

	logLevel = flag.String("logLevel", "INFO", "DEBUG|INFO|WARN|ERROR")

	httpPort        = flag.Int("httpPort", 8090, "http port")
	httpMetricsPort = flag.Int("httpMetricsPort", 8898, "http port")
	healthPort      = flag.Int("healthPort", 6696, "grpc health port")

	httpServer        *http.Server
	grpcHealthServer  *grpc.Server
	httpMetricsServer *http.Server
)

func main() {
	flag.Parse()

	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "caller", log.DefaultCaller, "ts", log.DefaultTimestampUTC)
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

	conn, err := grpc.Dial(*sataasURI,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name), // nolint:staticcheck
	)
	if err != nil {
		level.Error(logger).Log("msg", "can't reach sataas API", "error", err)
		os.Exit(2)
	}

	c := satsvc.NewPredictionClient(conn)

	// box html templates
	box := packr.New("Root box", "./templates")
	// web server
	s := NewServer(logger, c, box, *selfHostedMap, *tilesURL, *tilesKey)

	// web server
	g.Go(func() error {
		r := mux.NewRouter()
		r.HandleFunc("/api/tles", s.TLEHandler)
		//r.HandleFunc("/api/events", s.SSE.HTTPHandler)
		r.PathPrefix("/").Handler(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}))(s))

		httpServer = &http.Server{
			Addr:         fmt.Sprintf(":%d", *httpPort),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      handlers.CompressHandler(r),
		}
		level.Info(logger).Log("msg", fmt.Sprintf("HTTP server serving at :%d", *httpPort))

		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	// process streaming from api for /api/events endpoint
	// g.Go(func() error {
	// 	return s.Start(ctx)
	// })

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

	if httpMetricsServer != nil {
		_ = httpMetricsServer.Shutdown(shutdownCtx)
	}

	if httpServer != nil {
		_ = httpServer.Shutdown(shutdownCtx)
	}

	if grpcHealthServer != nil {
		grpcHealthServer.GracefulStop()
	}

	err = g.Wait()
	if err != nil {
		level.Debug(logger).Log("msg", "server returning an error", "error", err)
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
