package sataas

import (
	"errors"
	"sync"
	"time"

	"github.com/akhenakh/sataas/sgp4"
)

type Sat struct {
	*sgp4.SGP4
	*sgp4.TLE
	updateTime time.Time
}

// ActiveSats for the sataas
type ActiveSats struct {
	*sync.RWMutex
	sats map[int32]*Sat
}

// NewSatFromTLE returns a Sat created with a name and TLEs.
func NewSatFromTLE(tle *sgp4.TLE) (sat *Sat, err error) {
	s, err := sgp4.NewSGP4(tle)
	if err != nil {
		return nil, err
	}

	return &Sat{
		SGP4:       s,
		TLE:        tle,
		updateTime: time.Now(),
	}, nil
}

// NewActiveSats create a new list of active sats.
func NewActiveSats() *ActiveSats {
	return &ActiveSats{
		RWMutex: &sync.RWMutex{},
		sats:    make(map[int32]*Sat),
	}
}

// Get one sat.
func (as *ActiveSats) Get(norad int32) (*Sat, bool) {
	as.RLock()
	defer as.RUnlock()
	sat, ok := as.sats[norad]
	return sat, ok
}

func (as *ActiveSats) Set(norad int32, sat *Sat) {
	as.Lock()
	defer as.Unlock()
	as.sats[norad] = sat
}

// Position returns the position of a sat (only from now to future).
func (s *Sat) Position(lt time.Time) (lat, lng, alt float64, err error) {
	if lt.Add(1 * time.Second).Before(time.Now()) {
		return 0, 0, 0, errors.New("time in the past, unsupported")
	}
	return s.SGP4.Position(lt)
}
