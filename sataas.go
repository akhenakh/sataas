package sataas

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/akhenakh/sataas/sgp4"
	"github.com/akhenakh/sataas/sgp4svc"
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
	resp, err := http.Get(s.tleURL)
	if err != nil {
		level.Error(s.logger).Log("msg", "failed while fetching TLEs", "error", err, "tle_url", s.tleURL)
		return fmt.Errorf("failed while fetching TLEs %v", err)
	}
	defer resp.Body.Close()

	r := sgp4.NewTLEReader(resp.Body)
	tles, err := r.ReadAllTLE()
	if err != nil {
		level.Error(s.logger).Log("msg", "failed reading TLEs", "error", err)
		return fmt.Errorf("failed reading TLEs %v", err)
	}

	for _, tle := range tles {
		sat, err := NewSatFromTLE(tle)
		if err != nil {
			level.Warn(s.logger).Log("msg", "can't create sat from tle", "error", err, "tle_name", tle.Name())
			continue
		}

		s.sats.Set(int32(tle.NoradNumber()), sat)
	}

	return nil
}

// SatInfos gRPC exposed to query satellites infos.
func (s *Service) SatInfos(ctx context.Context, req *sgp4svc.SatRequest) (*sgp4svc.SatInfosResponse, error) {
	sat, ok := s.sats.Get(req.NoradNumber)
	if !ok {
		return nil, status.Error(codes.NotFound, "non existing norad id")
	}
	return &sgp4svc.SatInfosResponse{
		NoradNumber: int32(sat.TLE.NoradNumber()),
		Name:        sat.TLE.Name(),
		Tle1:        sat.TLE.Line1(),
		Tle2:        sat.TLE.Line2(),
	}, nil
}

// SatLocation gRPC exposed satellites position.
func (s *Service) SatLocation(ctx context.Context, req *sgp4svc.SatLocationRequest) (*sgp4svc.Location, error) {
	sat, ok := s.sats.Get(req.NoradNumber)
	if !ok {
		return nil, status.Error(codes.NotFound, "non existing norad id")
	}

	t, err := ptypes.Timestamp(req.Time)
	if err != nil {
		return nil, err
	}

	lat, lng, alt, err := sat.Position(t)
	if err != nil {
		return nil, err
	}

	return &sgp4svc.Location{
		Latitude:  lat,
		Longitude: lng,
		Altitude:  alt,
	}, nil
}
