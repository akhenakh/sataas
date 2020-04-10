package main

import (
	"fmt"
	"log"
	"time"

	"github.com/akhenakh/sataas/sgp4"
)

func main() {
	const l1 = "1 25544U 98067A   20099.88933163  .00000859  00000-0  23991-4 0  9991"
	const l2 = "2 25544  51.6455 331.3904 0003968  99.3691 357.2762 15.48691931221277"
	tle, err := sgp4.NewTLE(l1, l2)
	if err != nil {
		log.Fatal(err)
	}
	p, err := sgp4.NewSGP4(tle)
	if err != nil {
		log.Fatal(err)
	}
	lat, lng, alt := p.FindPosition(time.Now())
	fmt.Println("ISS position", lat, lng, alt)

	// excepting error and no panic
	tle, err = sgp4.NewTLE("", "")
	if err != nil {
		log.Fatal(err)
	}

}
