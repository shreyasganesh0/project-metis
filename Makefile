
BINARY_NAME=metisctl
BINARY_PATH=bin/$(BINARY_NAME)
VERSION_PKG_PATH=github.com/shreyasganesh0/project-metis/cmd/metisctl/cmd

VERSION ?= 0.1.0
COMMIT = $(shell git rev-parse --short HEAD)
DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS= -ldflags "-X '$(VERSION_PKG_PATH).version=$(VERSION)'\
				   -X '$(VERSION_PKG_PATH).commit=$(COMMIT)'\
				   -X '$(VERSION_PKG_PATH).date=$(DATE)'"

.PHONY: all build run

all: build

build:
	@echo "Building started..."
	@go build $(LDFLAGS) -o $(BINARY_PATH) ./cmd/metisctl

run: build
	@echo "Running binary..."
	@./$(BINARY_PATH)
