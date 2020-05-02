package sataas

import (
	"context"
	"fmt"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/golang/protobuf/ptypes"
	"github.com/hashicorp/go-retryablehttp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/akhenakh/sataas/satsvc"
	"github.com/akhenakh/sataas/sgp4"
)

type Service struct {
	logger log.Logger

	tleURL string
	sats   *ActiveSats
}

// New returns a new Sataas service manager.
func New(logger log.Logger, tleURL string) *Service {
	return &Service{
		logger: log.With(logger, "components", "service"),
		tleURL: tleURL,
		sats:   NewActiveSats(),
	}
}

// UpdateTLEs fetch all TLEs and update them.
func (s *Service) UpdateTLEs() error {
	resp, err := retryablehttp.Get(s.tleURL)
	if err != nil {
		return fmt.Errorf("failed while fetching TLEs: %w", err)
	}
	defer resp.Body.Close()

	r := sgp4.NewTLEReader(resp.Body)
	tles, err := r.ReadAllTLE()
	if err != nil {
		return fmt.Errorf("failed reading TLEs: %w", err)
	}

	for _, tle := range tles {
		sat, err := NewSatFromTLE(tle)
		if err != nil {
			level.Warn(s.logger).Log("msg", "can't create sat from tle", "error", err, "tle_name", tle.Name())
			continue
		}

		s.sats.Set(int32(tle.NoradNumber()), sat)
	}

	level.Info(s.logger).Log("msg", "updated TLEs", "count", len(tles))

	return nil
}

// SatInfos gRPC exposed to query satellites infos.
func (s *Service) SatInfos(ctx context.Context, req *satsvc.SatRequest) (*satsvc.SatInfosResponse, error) {
	sat, ok := s.sats.Get(req.NoradNumber)
	if !ok {
		return nil, status.Error(codes.NotFound, "non existing norad id")
	}
	return &satsvc.SatInfosResponse{
		NoradNumber: int32(sat.TLE.NoradNumber()),
		Name:        sat.TLE.Name(),
		Tle1:        sat.TLE.Line1(),
		Tle2:        sat.TLE.Line2(),
	}, nil
}

// SatLocation gRPC exposed satellites position.
func (s *Service) SatLocation(ctx context.Context, req *satsvc.SatLocationRequest) (*satsvc.Location, error) {
	sat, ok := s.sats.Get(req.NoradNumber)
	if !ok {
		return nil, status.Error(codes.NotFound, "non existing norad id")
	}
	t := time.Now()
	if req.Time != nil {
		pt, err := ptypes.Timestamp(req.Time)
		if err != nil {
			return nil, err
		}
		t = pt
	}

	lat, lng, alt, err := sat.Position(t)
	if err != nil {
		return nil, err
	}

	return &satsvc.Location{
		Latitude:  lat,
		Longitude: lng,
		Altitude:  alt,
	}, nil
}
