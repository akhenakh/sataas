.PHONY: all

TASK :=  $(shell task 2> /dev/null)
GOLANGCI :=  $(shell golangci-lint 2> /dev/null)

all: 
ifndef TASK
	$(error "task not available")
endif
	task -v

cicd-build:
ifndef TASK
	curl -sL https://taskfile.dev/install.sh | sh
endif
ifndef GOLANGCI
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
endif
	./bin/task -v
