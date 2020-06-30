.PHONY: all

TASK :=  $(shell task 2> /dev/null)

all: 
ifndef TASK
	$(error "task not available")
endif
	task -v

cicd-build:
ifndef TASK
	curl -sL https://taskfile.dev/install.sh | sh
endif
	task -v
