package sgp4

import (
	"fmt"
	"math"
	"time"

	"github.com/akhenakh/sataas/cppsgp4"
)

type SGP4 struct {
	csgp4 cppsgp4.SGP4
}

func NewSGP4(tle *TLE) (p *SGP4, err error) {
	defer catch(&err)
	cp := cppsgp4.NewSGP4(tle.ctle)

	p = &SGP4{csgp4: cp}
	return p, nil
}

func (p *SGP4) FindPosition(lt time.Time) (lat, lng, alt float64) {
	t := lt.UTC()
	dt := cppsgp4.NewDateTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
	eci := p.csgp4.FindPosition(dt)
	geo := eci.ToGeodetic()
	lat = geo.GetLatitude() * (180 / math.Pi)
	lng = geo.GetLongitude() * (180 / math.Pi)
	alt = geo.GetAltitude()
	return
}

func catch(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%v", r)
	}
}
