SHELL = bash

SEM_VER = 1.0.0
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_SHA_SHORT := $(shell git rev-parse --short HEAD)
COMMIT_TIME := $(shell git log -1 --format=%ci)
GO_VER := $(shell go version)

ROOT_DIR := $(shell pwd)
OUTPUT_DIR := ${ROOT_DIR}/dist
$(shell mkdir -p dist/bin)
APP := main

LDFLAGS = -ldflags \
	"-X build.name=${APP} \
	-X build.version=${SEM_VER} \
	-X build.gitBranch=${BRANCH} \
	-X build.lastCommitSHA=${COMMIT_SHA_SHORT} \
	-X 'build.lastCommitTime=${COMMIT_TIME}' \
	-X 'build.goVersion=${GO_VER}'"
BUILD_TAGS = -tags helloworld

build:
	go build ${LDFLAGS} ${BUILD_TAGS} -o ${OUTPUT_DIR}/bin/${APP}

run:
	# go run main.go
	${OUTPUT_DIR}/bin/${APP}

test:
	go test ${LDFLAGS} ${BUILD_TAGS} -v ./...

clean:
	rm -rf ${OUTPUT_DIR}

tags:
	@echo -n ${SEM_VER}\,latest > .tags

.PHONY: build clean