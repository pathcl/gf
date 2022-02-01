# Lots of good stuff here, use them for ideas
# - https://github.com/themotion/ladder/blob/master/Makefile
# - https://sohlich.github.io/post/go_makefile/

# Shell to use for running scripts
SHELL := $(shell which bash)

# Get docker path or an empty string
DOCKER := $(shell command -v docker)

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get

-include .config/.makerc
-include .config/.makerc.dist

# Application variables
BINARY_NAME ?= main
BINARY_NAME_UNIX ?= $(BINARY_NAME)_unix

# Build time variables
GIT_BRANCH      ?= $(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT      ?= $(shell git rev-parse --short HEAD)
GIT_TAG         ?= $(shell (git describe --tags --exact-match 2>/dev/null) || echo "dev-$(GIT_COMMIT)")
GO_VERSION      ?= $(shell go version | sed 's/go//g' | tr ' ' '\n' | grep '\d\+.\d\+.\d\+')
BUILD_DATE      ?= $(shell date)
BUILD_OS_HOST   ?= $(shell go env GOHOSTOS)
BUILD_ARCH_HOST ?= $(shell go env GOHOSTARCH)
BUILD_OS_UNIX   ?= linux
BUILD_ARCH_UNIX ?= $(shell go env GOHOSTARCH)

# Environment defaults
APPLICATION_NAME  ?= $(BINARY_NAME)
SHORT_DESCRIPTION ?= unknown
LONG_DESCRIPTION  ?= unknown

# Module information
GO_MODULE_NAME ?= $(shell go mod edit -json | jq -r .Module.Path)

# Injected variables
GO_LDFLAGS_SHARED = \
-X '$(GO_MODULE_NAME)/version.applicationName=$(APPLICATION_NAME)' \
-X '$(GO_MODULE_NAME)/version.shortDescription=$(SHORT_DESCRIPTION)' \
-X '$(GO_MODULE_NAME)/version.longDescription=$(LONG_DESCRIPTION)' \
-X '$(GO_MODULE_NAME)/version.applicationVersion=$(GIT_TAG)' \
-X '$(GO_MODULE_NAME)/version.gitVersion=$(GIT_COMMIT)' \
-X '$(GO_MODULE_NAME)/version.goVersion=$(GO_VERSION)' \
-X '$(GO_MODULE_NAME)/version.buildDate=$(BUILD_DATE)'

GO_LDFLAGS_HOST = \
$(GO_LDFLAGS_SHARED) \
-X '$(GO_MODULE_NAME)/version.buildOS=$(BUILD_OS_HOST)' \
-X '$(GO_MODULE_NAME)/version.buildArch=$(BUILD_ARCH_HOST)'

GO_LDFLAGS_UNIX = \
$(GO_LDFLAGS_SHARED) \
-X '$(GO_MODULE_NAME)/version.buildOS=$(BUILD_OS_UNIX)' \
-X '$(GO_MODULE_NAME)/version.buildArch=$(BUILD_ARCH_UNIX)'

default: build
all: test build

build:
		$(GOBUILD) -o bin/$(BINARY_NAME) -v -ldflags "$(GO_LDFLAGS_HOST)" main.go
		CGO_ENABLED=0 GOOS=$(BUILD_OS_UNIX) GOARCH=$(BUILD_ARCH_UNIX) $(GOBUILD) -o bin/$(BINARY_NAME_UNIX) -v -ldflags "$(GO_LDFLAGS_UNIX)" main.go

test:
		$(GOTEST) -v ./...

clean:
		$(GOCLEAN)
		rm -rf bin/$(BINARY_NAME)
		rm -rf bin/$(BINARY_NAME_UNIX)

run:
		$(GORUN) -ldflags "$(GO_LDFLAGS_HOST)" main.go $(args)
