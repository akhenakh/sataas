package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
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

	cats, err := c.Categories(ctx, &empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	var found int32
	for _, cat := range cats.Categories {
		if cat.Name == "STARLINK" {
			found = cat.Id
			log.Printf("found cat: %+v", cat)
		}
	}
	if found == 0 {
		log.Fatal("can't find valid cateogry for starlink")
	}

	req := &satsvc.SatsLocationsRequest{
		NoradNumbers: nil,
		Category:     found,
	}

	stream, err := c.SatsLocations(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		locations, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, loc := range locations.SatLocations {
			log.Printf("ID: %d Latitude : %.02f Longitude : %.02f Altitude: %.01fkm\n",
				loc.NoradNumber,
				loc.Latitude,
				loc.Longitude,
				loc.Altitude)
		}

	}
}
