package main

import (
	"fmt"

	"github.com/akhenakh/sataas/sgp4"
)

func main() {
	const l1 = "1 25544U 98067A   20099.88933163  .00000859  00000-0  23991-4 0  9991"
	const l2 = "2 25544  51.6455 331.3904 0003968  99.3691 357.2762 15.48691931221277"
	tle := sgp4.NewTle(l1, l2)
	p := sgp4.NewSGP4(tle)

	dt := sgp4.DateTimeNow(false)
	eci := p.FindPosition(dt)
	geo := eci.ToGeodetic()
	fmt.Println(geo.GetLatitude(), geo.GetLongitude(), geo.GetAltitude())
}
