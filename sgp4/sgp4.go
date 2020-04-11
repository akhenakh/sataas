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

type Observation struct {
	SatLat      float64
	SatLng      float64
	SatAltitude float64
	Azimuth     float64
	Elevation   float64
	Range       float64
	RangeRate   float64
}

func (p *SGP4) ObservationFromLocation(lat, lng, alt float64) Observation {
	obs := cppsgp4.NewObserver(lat, lng, alt)
	now := cppsgp4.DateTimeNow(true)

	// calculate satellite position
	eci := p.csgp4.FindPosition(now)

	// get look angle for observer to satellite
	topo := obs.GetLookAngle(eci)

	// convert satellite position to geodetic coordinates
	geo := eci.ToGeodetic()

	slat := geo.GetLatitude() * (180 / math.Pi)
	slng := geo.GetLongitude() * (180 / math.Pi)
	salti := geo.GetAltitude() * 1000

	return Observation{
		SatLat:      slat,
		SatLng:      slng,
		SatAltitude: salti,
		Azimuth:     topo.GetAzimuth() * (180 / math.Pi),
		Elevation:   topo.GetElevation() * (180 / math.Pi),
		Range:       topo.GetXrange(),
		RangeRate:   topo.GetRange_rate(),
	}
}
