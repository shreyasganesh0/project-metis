.PHONY: all build run

BINARY_NAME=metisctl
BINARY_PATH=bin/$(BINARY_NAME)

all: build

build:
	@echo "Building started..."
	@go build -o $(BINARY_PATH) ./cmd/metisctl

run: build
	@echo "Running binary..."
	@./$(BINARY_PATH)
