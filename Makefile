SHELL = bash
SEMVER = 1.0.0
ROOT_DIR := $(shell pwd)
OUTPUT_DIR := ${ROOT_DIR}/dist
APP := main
$(shell mkdir -p dist/bin)

test:
	go test -v ./...

build:
	go build -o dist/bin/${APP}

tags:
	@echo -n ${SEMVER}\,latest > .tags
	@echo -n ${SEMVER} > .version

clean:
	rm -rf dist