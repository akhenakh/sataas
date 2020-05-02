package main

import (
	"context"
	"log"
	"time"

	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"

	"github.com/akhenakh/sataas/satsvc"
)

var (
	sataasURI = flag.String("sataasURI", "localhost:9200", "sataas grpc URI")

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
}
