SHELL = bash

APP := main
SEM_VER = 1.0.0
PROJ_NAME := $(shell basename "$(PWD)")
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_SHA_SHORT := $(shell git rev-parse --short HEAD)
COMMIT_TIME := $(shell git log -1 --format=%ci)
GO_VER := $(shell go version)

REPO := github.com/mayone/Go-Utilities/helloworld
ROOT_DIR := $(shell pwd)
OUTPUT_DIR := ${ROOT_DIR}/dist
$(shell mkdir -p dist/bin)

LDFLAGS = -ldflags \
	"-X ${REPO}/buildinfo.name=${APP} \
	-X ${REPO}/buildinfo.version=${SEM_VER} \
	-X ${REPO}/buildinfo.gitBranch=${BRANCH} \
	-X ${REPO}/buildinfo.lastCommitSHA=${COMMIT_SHA_SHORT} \
	-X '${REPO}/buildinfo.lastCommitTime=${COMMIT_TIME}' \
	-X '${REPO}/buildinfo.goVersion=${GO_VER}'"
BUILD_TAGS = -tags dev

## [build]: Build app
all: build
build:
	go build ${LDFLAGS} ${BUILD_TAGS} -o ${OUTPUT_DIR}/bin/${APP}

## run: Run app
run:
	${OUTPUT_DIR}/bin/${APP}

## test: Run tests
test:
	go test ${LDFLAGS} ${BUILD_TAGS} -v ./...

## clean: Clean build files.
clean:
	rm -rf ${OUTPUT_DIR}

tags:
	@echo -n ${SEM_VER}\,latest > .tags

.PHONY: build clean

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJ_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'