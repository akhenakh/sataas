package sataas

import (
	"errors"
	"sync"
	"time"

	"github.com/akhenakh/sataas/satsvc"
	"github.com/akhenakh/sataas/sgp4"
)

// Sat is holding live propagation obj to compute predictions.
type Sat struct {
	*sgp4.SGP4
	*sgp4.TLE
	updateTime time.Time
}

// Category is holding categories of sats
type Category struct {
	ID   int32
	Name string
	Sats []int32 // norad numbers
}

// ActiveSats for the sataas.
type ActiveSats struct {
	*sync.RWMutex
	sats map[int32]*Sat
}

// ActiveCategories for the sataas.
type ActiveCategories struct {
	*sync.RWMutex
	categories     map[int32]*Category
	grpcCategories *satsvc.CategoriesResponse
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

// Get one sat, read thread safe.
func (as *ActiveSats) Get(norad int32) (*Sat, bool) {
	as.RLock()
	defer as.RUnlock()
	sat, ok := as.sats[norad]
	return sat, ok
}

// Set one sat, read thread safe.
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

// NewActiveCategories create a new list of categories.
func NewActiveCategories() *ActiveCategories {
	return &ActiveCategories{
		RWMutex:    &sync.RWMutex{},
		categories: make(map[int32]*Category),
	}
}

// Set one cat, read thread safe.
func (ac *ActiveCategories) Set(id int32, name string, sats []int32) {
	ac.Lock()
	defer ac.Unlock()
	ac.categories[id] = &Category{
		ID:   id,
		Name: name,
		Sats: sats,
	}
	grpcCat := &satsvc.CategoriesResponse{
		Categories: make([]*satsvc.Category, len(ac.categories)),
	}
	i := 0
	for _, v := range ac.categories {
		grpcCat.Categories[i] = &satsvc.Category{
			Id:   v.ID,
			Name: v.Name,
			Sats: v.Sats,
		}
		i++
	}
	ac.grpcCategories = grpcCat
}
