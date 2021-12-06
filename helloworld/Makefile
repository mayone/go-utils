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

RESET  := $(shell tput -Txterm sgr0)
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
CYAN   := $(shell tput -Txterm setaf 6)
WHITE  := $(shell tput -Txterm setaf 7)

.PHONY: all build clean

all: build

## Build:

build: ## Build app.
	go build ${LDFLAGS} ${BUILD_TAGS} -o ${OUTPUT_DIR}/bin/${APP}

run: ## Run app.
	${OUTPUT_DIR}/bin/${APP}

tags: ## Create tag file.
	@echo -n ${SEM_VER}\,latest > .tags

clean: ## Remove build related files.
	rm -rf ${OUTPUT_DIR}

## Test:
test: ## Run tests.
	go test ${LDFLAGS} ${BUILD_TAGS} -v ./...

bench: ## Run benchmark.
	go test ${LDFLAGS} ${BUILD_TAGS} -bench=BenchmarkMemory -benchmem -run=^$ -count=10 ./...

## Help:
help: Makefile ## Show this help message.
	@echo
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)