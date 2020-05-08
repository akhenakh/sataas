package main

import (
	"context"
	"io"
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

	noradNumber = flag.Uint("noradNumber", 25544, "norad number sat to query")
	stepsMs     = flag.Uint("stepsMs", 1000, "recompute position every stepsMs (in ms)")

	minElevation = flag.Float64("minElevation", 0, "filter passes with max elevations lower than minElevation")
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

	obsLoc := &satsvc.Location{
		Latitude:  *lat,
		Longitude: *lng,
		Altitude:  *alt,
	}

	starttp := ptypes.TimestampNow()

	stopt := time.Now().Add(*duration)

	stoptp, _ := ptypes.TimestampProto(stopt)

	passes, err := c.GenPasses(ctx, &satsvc.GenPassesRequest{
		NoradNumber:      int32(*noradNumber),
		ObserverLocation: obsLoc,
		StartTime:        starttp,
		StopTime:         stoptp,
		StepSeconds:      30,
		MinElevation:     *minElevation,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Computed passes %d\n", len(passes.Passes))
	for _, p := range passes.Passes {
		log.Printf("Pass %s\n", p.String())
	}

	req := &satsvc.SatLocationFromObsRequest{
		NoradNumbers:     []int32{int32(*noradNumber)},
		ObserverLocation: obsLoc,
		StepsMs:          int32(*stepsMs),
	}

	stream, err := c.SatLocationFromObs(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		obs, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Latitude : %.02f Longitude : %.02f Altitude: %.01fkm\n"+
			"Azimuth : %.0f Elevation %.01f Range: %.01fkm RangeRate: %f\n",
			obs.SatLocation.Latitude,
			obs.SatLocation.Longitude,
			obs.SatLocation.Altitude,
			obs.Azimuth,
			obs.Elevation,
			obs.Range,
			obs.RangeRate)
	}
}
