package main

import (
	"github.com/go-kit/kit/log"

	"github.com/akhenakh/sataas/satsvc"
)

type Server struct {
	logger log.Logger
	client *satsvc.PredictionClient
}
