package main

import (
	"context"
	"encoding/json"
	"html/template"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gobuffalo/packr/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/r3labs/sse"

	"github.com/akhenakh/sataas/satsvc"
)

const starlinkChannel = "starlink"

var (
	pathTpl = []string{"index.html"}
)

type Server struct {
	logger        log.Logger
	client        satsvc.PredictionClient
	fileHandler   http.Handler
	box           *packr.Box
	tilesKey      string
	selfHostedMap bool
	tilesURL      string
	SSE           *sse.Server
}

func NewServer(
	logger log.Logger,
	client satsvc.PredictionClient,
	box *packr.Box, selfHostedMap bool,
	tilesURL,
	tilesKey string,
) *Server {
	logger = log.With(logger, "component", "web")

	s := &Server{
		logger:        logger,
		client:        client,
		box:           box,
		fileHandler:   http.FileServer(box),
		selfHostedMap: selfHostedMap,
		tilesKey:      tilesKey,
		tilesURL:      tilesURL,
		SSE:           sse.New(),
	}
	s.SSE.CreateStream(starlinkChannel)

	return s
}

func (s *Server) Start(ctx context.Context) error {
	cats, err := s.client.Categories(ctx, &empty.Empty{})
	if err != nil {
		level.Error(s.logger).Log("msg", "can't fetch categories from API", "error", err)
		os.Exit(2)
	}

	var found int32
	for _, cat := range cats.Categories {
		if cat.Name == "STARLINK" {
			found = cat.Id
			level.Info(s.logger).Log("msg", "found starlink category", "cat", cat)
		}
	}
	if found == 0 {
		level.Error(s.logger).Log("msg", "can't find starlink category")
		return err
	}

	req := &satsvc.SatsRequest{
		NoradNumbers: nil,
		Category:     found,
	}

	stream, err := s.client.SatsLocations(context.Background(), req)
	if err != nil {
		return err
	}

	level.Info(s.logger).Log("msg", "starting publishing to sse")

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		locations, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		level.Info(s.logger).Log("msg", "publishing to sse", "count", len(locations.SatLocations))

		res, err := json.Marshal(locations.SatLocations)
		if err != nil {
			return err
		}
		s.SSE.Publish(starlinkChannel, &sse.Event{
			Data: res,
		})
	}
	return nil
}

func (s *Server) TLEHandler(w http.ResponseWriter, r *http.Request) {
	cats, err := s.client.Categories(r.Context(), &empty.Empty{})
	if err != nil {
		level.Error(s.logger).Log("msg", "can't fetch categories from API", "error", err)
		http.Error(w, err.Error(), 500)
		return
	}

	var found int32
	for _, cat := range cats.Categories {
		if cat.Name == "STARLINK" {
			found = cat.Id
			level.Info(s.logger).Log("msg", "found starlink category", "cat", cat)
		}
	}
	if found == 0 {
		level.Error(s.logger).Log("msg", "can't find starlink category")
		http.Error(w, "can't find starlink category", 500)
		return
	}

	req := &satsvc.SatsRequest{
		NoradNumbers: nil,
		Category:     found,
	}

	sats, err := s.client.SatsInfos(r.Context(), req)
	if err != nil {
		level.Error(s.logger).Log("msg", "can't fetch tle from API", "error", err)
		http.Error(w, err.Error(), 500)
		return
	}

	b, err := json.Marshal(sats.SatInfos)
	if err != nil {
		level.Error(s.logger).Log("msg", "can't marshal tle", "error", err)
		http.Error(w, err.Error(), 500)
		return
	}

	_, _ = w.Write(b)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	if path == "" {
		path = "index.html"
	}

	p := map[string]interface{}{
		"TilesURL":      s.tilesURL,
		"TilesKey":      s.tilesKey,
		"Lat":           48.864716,
		"Lng":           2.349014,
		"SelfHostedMap": s.selfHostedMap,
	}

	// serve file normally
	if !isTpl(path) {
		s.fileHandler.ServeHTTP(w, r)
		return
	}

	tmplt := template.New(path)

	sf, err := s.box.FindString(path)
	if err != nil {
		level.Error(s.logger).Log("msg", "can't open template", "error", err)
		http.Error(w, err.Error(), 500)
		return
	}

	tmplt, err = tmplt.Parse(sf)
	if err != nil {
		http.Error(w, err.Error(), 500)
		level.Error(s.logger).Log("msg", "can't parse template", "error", err)
		return
	}

	ctype := mime.TypeByExtension(filepath.Ext(path))
	w.Header().Set("Content-Type", ctype)

	if err := tmplt.Execute(w, p); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func isTpl(path string) bool {
	for _, p := range pathTpl {
		if p == path {
			return true
		}
	}
	return false
}
