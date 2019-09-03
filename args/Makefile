SHELL = bash
SEMVER = 1.0.0
ROOT_DIR := $(shell pwd)
OUTPUT_DIR := ${ROOT_DIR}/dist
$(shell mkdir -p dist/bin)

test:
	go test -v ./...

build:
	go build -o dist/bin/main

tags:
	@echo -n ${SEMVER}\,latest > .tags
	@echo -n ${SEMVER} > .version
