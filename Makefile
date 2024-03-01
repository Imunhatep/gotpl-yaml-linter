# set the shell to bash in case some environments use sh
SHELL:=/bin/bash

# VERSION is the version of the binary.
VERSION:=$(shell git describe --tags --always)

# commit ref
REPO = $(shell sh -c "git ls-remote --get-url origin | cut -f 2 -d @" | awk -F ".git" '{print $$1}' | sed 's/:/\//')

# user
USER_UID = $(shell id -u)
USER_GID = $(shell id -g)

# os
UNAME_S := $(shell uname -s)
GOOS = linux
ifeq ($(UNAME_S),Windows)
	GOOS = windows
endif
ifeq ($(UNAME_S),Darwin)
	GOOS = darwin
endif
export GOOS

# arch
UNAME_P := $(shell uname -p)
GOARCH = amd64
ifneq ($(filter %86,$(UNAME_P)),)
	GOARCH = 386
endif
ifneq ($(filter arm%,$(UNAME_P)),)
	GOARCH = arm64
endif
export GOARCH

# docker vs podman
ifeq ($(shell command -v podman 2> /dev/null),)
    DOCKER=docker
else
    DOCKER=podman
endif

IMAGE_PATH = imunhatep/yamltpl_linter
ifeq (${IMAGE_ORG}, )
  IMAGE_ORG = docker.io
  export IMAGE_ORG
endif

# Specify the date of build
DBUILD_DATE=$(shell date -u +'%Y%m%dT%H%M%SZ')
export DBUILD_ARGS=--build-arg DBUILD_DATE=${DBUILD_DATE}


# -composite: avoid "literal copies lock value from fakePtr"
.PHONY: vet
vet:
	go list ./... | grep -v "./vendor/*" | xargs go vet -composites

.PHONY: fmt
fmt:
	find . -type f -name "*.go" | grep -v "./vendor/*" | xargs gofmt -s -w -l

.PHONY: lint
lint:
	golangci-lint run

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: build.common
build.common: version

.PHONY: clean
clean:
	@echo '--> Cleaning directory...'
	rm -rf bin
	@echo '--> Done cleaning.'

# Compile binaries and build docker images
.PHONY: docker
docker:
	$(DOCKER) build -t ${IMAGE_ORG}/${IMAGE_PATH}:${DBUILD_DATE} ${DBUILD_ARGS} -f ./build/cli/Dockerfile .

.PHONY: build
build: clean build.common build-cli

.PHONY: build-cli
build-cli:
	mkdir -p ./bin && env CGO_ENABLED=0 go build -o bin/yamltpl_lint_${GOOS}-${GOARCH} -ldflags "-s -w" ./cmd/cli/main.go

.PHONY: test
test:
	$(DOCKER) build -t ${IMAGE_ORG}:test -f build/test/Dockerfile .

.PHONY: update-deps
update-deps:
	@echo ">> updating Go dependencies"
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done
	go mod tidy
ifneq (,$(wildcard vendor))
	go mod vendor
endif