.EXPORT_ALL_VARIABLES:

ifndef VERSION
VERSION := $(shell git describe --always --tags)
endif

DATE := $(shell date -u +%Y%m%d.%H%M%S)

LDFLAGS = -trimpath -ldflags "-X=main.version=$(VERSION)-$(DATE)"
CGO_ENABLED=1

targets = sataas 

.PHONY: all lint test clean

all: lint test $(targets)

test: 
	go test -race ./...

lint:
	golangci-lint run

sataas:
	cd cmd/sataas && go build -a $(LDFLAGS)

swig:
	cd cppsgp4 && swig -c++ -intgosize 64 -go SGP4.i

cmd/sataas/grpc_health_probe: GRPC_HEALTH_PROBE_VERSION=v0.3.2
cmd/sataas/grpc_health_probe:
	wget -qOcmd/sataas/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
		chmod +x cmd/sataas/grpc_health_probe

grpc_health_probe: cmd/sataas/grpc_health_probe

docker-image: sataas grpc_health_probe
	cd ./cmd/sataas/ && docker build . -t sataas:${VERSION}
	docker tag sataas:${VERSION} akhenakh/sataas:latest

clean:
	rm -f cmd/sataas/sataas
	rm -r cmd/sataas/grpc_health_probe
