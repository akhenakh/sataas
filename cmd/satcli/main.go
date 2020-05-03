package main

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"

	"github.com/akhenakh/sataas/satsvc"
)

var (
	sataasURI = flag.String("sataasURI", "localhost:9200", "sataas grpc URI")

	lat = flag.Float64("lat", 46.83, "latitude of observer")
	lng = flag.Float64("lng", -71.25, "longitude of observer")
	alt = flag.Float64("alt", 0, "altitude of observer")

	duration = flag.Duration("duration", 24*time.Hour, "compute passes from now to duration")

	noradNumber = flag.Uint("noradNumber", 25544, "Norad number sat to query")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*sataasURI,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name), // nolint:staticcheck
	)
	if err != nil {
		log.Fatal(err)
	}

	c := satsvc.NewPredictionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := c.SatInfos(ctx, &satsvc.SatRequest{
		NoradNumber: int32(*noradNumber),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sat %+v", resp)

	ctx, ccancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer ccancel()

	loc, err := c.SatLocation(ctx, &satsvc.SatLocationRequest{
		NoradNumber: int32(*noradNumber),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Location %+v", loc)

	starttp := ptypes.TimestampNow()

	stopt := time.Now().Add(*duration)

	stoptp, _ := ptypes.TimestampProto(stopt)

	passes, err := c.GenPasses(ctx, &satsvc.GenPassesRequest{
		NoradNumber: int32(*noradNumber),
		Location: &satsvc.Location{
			Latitude:  *lat,
			Longitude: *lng,
			Altitude:  *alt,
		},
		StartTime:   starttp,
		StopTime:    stoptp,
		StepSeconds: 30,
	})

	log.Printf("Passes to %s:\n%#v", stopt, passes.String())
}
