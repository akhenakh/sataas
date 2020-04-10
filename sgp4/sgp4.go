package sgp4

import (
	"fmt"
	"math"
	"time"

	"github.com/akhenakh/sataas/cppsgp4"
)

type TLE struct {
	ctle cppsgp4.Tle
}
type SGP4 struct {
	csgp4 cppsgp4.SGP4
}

func NewTLE(tle1, tle2 string) (tle *TLE, err error) {
	defer catch(&err)
	ctle := cppsgp4.NewTle(tle1, tle2)

	tle = &TLE{ctle: ctle}
	return tle, nil
}

func NewSGP4(tle *TLE) (p *SGP4, err error) {
	defer catch(&err)
	cp := cppsgp4.NewSGP4(tle.ctle)

	p = &SGP4{csgp4: cp}
	return p, nil

}

func (p *SGP4) FindPosition(t time.Time) (lat, lng, alt float64) {
	// TODO: pass real time
	dt := cppsgp4.DateTimeNow(false)
	eci := p.csgp4.FindPosition(dt)
	geo := eci.ToGeodetic()
	return geo.GetLatitude() * (180 / math.Pi), geo.GetLongitude() * (180 / math.Pi), geo.GetAltitude()
}

func catch(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%v", r)
	}
}
