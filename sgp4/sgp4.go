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

func (p *SGP4) Position(lt time.Time) (lat, lng, alt float64, err error) {
	defer catch(&err)
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
		*err = fmt.Errorf("%v", r) // nolint:goerr113
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

type PassDetails struct {
	AOS          time.Time
	LOS          time.Time
	AOSAzimuth   float64
	LOSAzimuth   float64
	MaxElevation float64
	AOSRangeRate float64
	LOSRangeRate float64
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

	slat := geo.GetLatitude() * (180.0 / math.Pi)
	slng := geo.GetLongitude() * (180.0 / math.Pi)
	salti := geo.GetAltitude()

	return Observation{
		SatLat:      slat,
		SatLng:      slng,
		SatAltitude: salti,
		Azimuth:     topo.GetAzimuth() * (180.0 / math.Pi),
		Elevation:   topo.GetElevation() * (180.0 / math.Pi),
		Range:       topo.GetXrange(),
		RangeRate:   topo.GetRange_rate(),
	}
}

func timeFromJulian(julian float64) time.Time {
	const epoch = 2440587.5
	seconds := (julian - epoch) * 86400
	return time.Unix(int64(seconds), 0)
}

func (p *SGP4) GeneratePasses(lat, lng, alt float64, start, stop time.Time, step int) []PassDetails {
	startdt := cppsgp4.NewDateTime(
		start.Year(), int(start.Month()), start.Day(),
		start.Hour(), start.Minute(), start.Second())

	stopdt := cppsgp4.NewDateTime(
		stop.Year(), int(stop.Month()), stop.Day(),
		stop.Hour(), stop.Minute(), stop.Second())

	cdetails := cppsgp4.GeneratePassList(lat, lng, alt, p.csgp4, startdt, stopdt, step)
	details := make([]PassDetails, cdetails.Capacity())

	for i := 0; i < int(cdetails.Capacity()); i++ {
		cpd := cdetails.Get(i)

		aos := timeFromJulian(cpd.GetAos().ToJulian())
		los := timeFromJulian(cpd.GetAos().ToJulian())

		details[i] = PassDetails{
			AOS:          aos,
			LOS:          los,
			AOSAzimuth:   cpd.GetAos_azimuth() * (180.0 / math.Pi),
			LOSAzimuth:   cpd.GetLos_azimuth() * (180.0 / math.Pi),
			MaxElevation: cpd.GetMax_elevation() * (180.0 / math.Pi),
			AOSRangeRate: cpd.GetAos_range_rate(),
			LOSRangeRate: cpd.GetLos_range_rate(),
		}
	}
	return details
}
