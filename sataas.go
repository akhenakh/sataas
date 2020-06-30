package sataas

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-retryablehttp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/status"

	"github.com/akhenakh/sataas/satsvc"
	"github.com/akhenakh/sataas/sgp4"
)

type Service struct {
	logger log.Logger
	Health *health.Server

	tleURL      string
	categoryURL string
	sats        *ActiveSats
	categories  *ActiveCategories

	gCtx context.Context
}

// New returns a new Sataas service manager.
func New(ctx context.Context, logger log.Logger, health *health.Server, tleURL, categoryURL string) *Service {
	return &Service{
		logger:      log.With(logger, "components", "service"),
		Health:      health,
		tleURL:      tleURL,
		categoryURL: categoryURL,
		sats:        NewActiveSats(),
		categories:  NewActiveCategories(),
		gCtx:        ctx,
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

// UpdateCategories fetch categories and update them.
func (s *Service) UpdateCategories() error {
	resp, err := retryablehttp.Get(s.categoryURL)
	if err != nil {
		return fmt.Errorf("failed while fetching categories: %w", err)
	}
	defer resp.Body.Close()

	var categories []Category
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return fmt.Errorf("failed unmarshaling categories: %w", err)
	}

	for _, cat := range categories {
		var validSats []int32
		for _, satid := range cat.Sats {
			if _, ok := s.sats.sats[satid]; !ok {
				level.Warn(s.logger).Log("msg", "missing sat from category", "category", cat.ID, "sat_id", satid)

				continue
			}
			validSats = append(validSats, satid)
		}
		s.categories.Set(cat.ID, cat.Name, validSats)
	}

	level.Info(s.logger).Log("msg", "updated Categories", "count", len(categories))

	return nil
}

// SatInfos gRPC exposed to query satellites infos.
func (s *Service) SatsInfos(ctx context.Context, req *satsvc.SatsRequest) (*satsvc.SatsInfosResponse, error) {
	if len(req.NoradNumbers) == 0 && req.Category == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Category != 0 {
		cat, ok := s.categories.categories[req.Category]
		if !ok {
			return nil, status.Error(codes.NotFound, "invalid category")
		}
		req.NoradNumbers = cat.Sats
	}

	resp := &satsvc.SatsInfosResponse{
		SatInfos: make([]*satsvc.SatInfos, len(req.NoradNumbers)),
	}

	for i, satid := range req.NoradNumbers {
		sat, ok := s.sats.Get(satid)
		if !ok {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("non existing norad id %d", satid))
		}

		ut, err := ptypes.TimestampProto(sat.updateTime)
		if !ok {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't serialize time %v", err))
		}

		resp.SatInfos[i] = &satsvc.SatInfos{
			NoradNumber: int32(sat.TLE.NoradNumber()),
			Name:        sat.TLE.Name(),
			Tle1:        sat.TLE.Line1(),
			Tle2:        sat.TLE.Line2(),
			UpdateTime:  ut,
		}
	}

	return resp, nil
}

// SatLocations gRPC exposed satellites position.
func (s *Service) SatsLocations(req *satsvc.SatsRequest, stream satsvc.Prediction_SatsLocationsServer) error {
	if len(req.NoradNumbers) == 0 && req.Category == 0 {
		return status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Category != 0 {
		cat, ok := s.categories.categories[req.Category]
		if !ok {
			return status.Error(codes.NotFound, "invalid category")
		}
		req.NoradNumbers = cat.Sats
	}

	for _, satid := range req.NoradNumbers {
		_, ok := s.sats.Get(satid)
		if !ok {
			return status.Error(codes.NotFound, fmt.Sprintf("non existing norad id %d", satid))
		}
	}

	ticker := time.NewTicker(1 * time.Second)

	resp := &satsvc.SatsLocationsResponse{
		SatLocations: make([]*satsvc.SatLocation, len(req.NoradNumbers)),
	}

	for {
		select {
		case <-s.gCtx.Done():
			ticker.Stop()
			return nil
		case <-stream.Context().Done():
			ticker.Stop()
			return nil
		case <-ticker.C:

			for i, nid := range req.NoradNumbers {
				sat, ok := s.sats.Get(nid)
				if !ok {
					return status.Error(codes.NotFound, fmt.Sprintf("non existing norad id %d", nid))
				}
				t := time.Now()
				lat, lng, alt, err := sat.Position(t)
				if err != nil {
					return err
				}
				resp.SatLocations[i] = &satsvc.SatLocation{
					NoradNumber: nid,
					Latitude:    lat,
					Longitude:   lng,
					Altitude:    alt,
				}
			}

			if err := stream.Send(resp); err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
		}
	}
}

// SatLocationFromObs gRPC exposed stream Live Sat observation.
func (s *Service) SatLocationFromObs(req *satsvc.SatLocationFromObsRequest,
	stream satsvc.Prediction_SatLocationFromObsServer) error {
	if req.ObserverLocation == nil {
		return status.Error(codes.InvalidArgument, "invalid observer location")
	}

	ticker := time.NewTicker(time.Duration(req.StepsMs) * time.Millisecond)

	for {
		select {
		case <-s.gCtx.Done():
			ticker.Stop()
			return nil
		case <-stream.Context().Done():
			ticker.Stop()
			return nil
		case <-ticker.C:
			for _, nid := range req.NoradNumbers {
				sat, ok := s.sats.Get(nid)
				if !ok {
					return status.Error(codes.NotFound, fmt.Sprintf("non existing norad id %d", nid))
				}

				sobs := sat.SGP4.ObservationFromLocation(
					req.ObserverLocation.Latitude,
					req.ObserverLocation.Longitude,
					req.ObserverLocation.Altitude,
				)
				obs := &satsvc.Observation{
					NoradNumber: nid,
					SatLocation: &satsvc.Location{
						Latitude:  sobs.SatLat,
						Longitude: sobs.SatLng,
						Altitude:  sobs.SatAltitude,
					},
					Azimuth:   sobs.Azimuth,
					Elevation: sobs.Elevation,
					Range:     sobs.Range,
					RangeRate: sobs.RangeRate,
				}

				if err := stream.Send(obs); err != nil {
					if err == io.EOF {
						return nil
					}
					return err
				}
			}
		}
	}
}

// GenLocations gRPC exposed to generate positions.
func (s *Service) GenLocations(
	ctx context.Context,
	req *satsvc.GenLocationsRequest,
) (*satsvc.GenLocationsResponse, error) {
	if len(req.NoradNumbers) == 0 && req.Category == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Category != 0 {
		cat, ok := s.categories.categories[req.Category]
		if !ok {
			return nil, status.Error(codes.NotFound, "invalid category")
		}
		req.NoradNumbers = cat.Sats
	}

	resp := &satsvc.GenLocationsResponse{
		Locations:   make([]*satsvc.GenLocations, len(req.NoradNumbers)),
		StepSeconds: req.StepSeconds,
	}

	startt, err := ptypes.Timestamp(req.StartTime)
	if err != nil {
		return nil, err
	}

	stopp, err := ptypes.Timestamp(req.StopTime)
	if err != nil {
		return nil, err
	}

	for i, satid := range req.NoradNumbers {
		sat, ok := s.sats.Get(satid)
		if !ok {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("non existing norad id %d", satid))
		}

		locs := sat.GenerateLocations(startt, stopp, int(req.StepSeconds))
		resp.Locations[i] = &satsvc.GenLocations{
			NoradNumber: satid,
			Location:    make([]*satsvc.Location, len(locs)),
		}
		for j, loc := range locs {
			resp.Locations[i].Location[j] = &satsvc.Location{
				Latitude:  loc.SatLng,
				Longitude: loc.SatLng,
				Altitude:  loc.SatAltitude,
			}
		}
	}

	return resp, nil
}

// GenPasses gRPC exposed to generate satellites passes.
func (s *Service) GenPasses(ctx context.Context, req *satsvc.GenPassesRequest) (*satsvc.Passes, error) {
	sat, ok := s.sats.Get(req.NoradNumber)
	if !ok {
		return nil, status.Error(codes.NotFound, "non existing norad id")
	}

	if req.ObserverLocation == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid observer location")
	}

	if req.StartTime == nil || req.StopTime == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid time")
	}

	startt, err := ptypes.Timestamp(req.StartTime)
	if err != nil {
		return nil, err
	}

	stopp, err := ptypes.Timestamp(req.StopTime)
	if err != nil {
		return nil, err
	}

	passesDetails := sat.GeneratePasses(
		req.ObserverLocation.Latitude,
		req.ObserverLocation.Latitude,
		req.ObserverLocation.Altitude,
		startt,
		stopp,
		int(req.StepSeconds),
	)

	var passes []*satsvc.Pass

	for _, pd := range passesDetails {
		aos, err := ptypes.TimestampProto(pd.AOS)
		if !ok {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't serialize aos time %v", err))
		}

		los, err := ptypes.TimestampProto(pd.LOS)
		if !ok {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't serialize los time %v", err))
		}

		// filter out passes with low elevation
		if req.MinElevation != 0.0 && pd.MaxElevation < req.MinElevation {
			continue
		}

		passes = append(passes, &satsvc.Pass{
			Aos:          aos,
			Los:          los,
			AosAzimuth:   pd.AOSAzimuth,
			LosAzimuth:   pd.LOSAzimuth,
			MaxElevation: pd.MaxElevation,
			AosRangeRate: pd.AOSRangeRate,
			LosRangeRate: pd.LOSRangeRate,
		})
	}
	return &satsvc.Passes{Passes: passes}, nil
}

// Categories gRPC exposed to get category list.
func (s *Service) Categories(ctx context.Context, empty *empty.Empty) (*satsvc.CategoriesResponse, error) {
	return s.categories.grpcCategories, nil
}
