package main

import (
	"fmt"
	"log"
	"time"

	"github.com/akhenakh/sataas/sgp4"
)

func main() {
	const l1 = "1 25544U 98067A   20102.72807343  .00002441  00000-0  53009-4 0  9991"
	const l2 = "2 25544  51.6439 317.3578 0003912 108.5059 345.3148 15.48688334221718"
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
	obv := p.ObservationFromLocation(46.829853, -71.254028, 0)

	fmt.Printf("Observation %+v\n", obv)

	passes := p.GeneratePasses(46.829853, -71.254028, 0, time.Now(), time.Now().Add(24*10*time.Hour), 60)
	fmt.Printf("Passes for the next 10 days:\n%+v\n", passes)
}
