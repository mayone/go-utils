SHELL = bash
SEMVER = 1.0.0
ROOT_DIR := $(shell pwd)
OUTPUT_DIR := ${ROOT_DIR}/dist
$(shell mkdir -p dist/bin)
EXEC := swap

build:
	go build -o $(OUTPUT_DIR)/bin/$(EXEC)

run:
	go run swap.go

test:
	go test -v ./...

tags:
	@echo -n ${SEMVER}\,latest > .tags
	@echo -n ${SEMVER} > .version