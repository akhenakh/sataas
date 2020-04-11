package main

import (
	"fmt"
	"log"
	"time"

	"github.com/akhenakh/sataas/sgp4"
)

func main() {
	const l1 = "1 25544U 98067A   20101.90690972 -.00000449  00000-0  00000+0 0  9993"
	const l2 = "2 25544  51.6446 321.4198 0003848 108.5166  84.0719 15.48680394221581"
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
}
