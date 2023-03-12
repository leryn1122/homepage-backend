# Project
SHELL := /usr/bin/env sh
NAME := homepage
VERSION := $(shell git describe --long --all)
BUILD_DATE := $(shell date +%Y%m%d)
GIT_VERSION := $(shell git describe --long --all)
SHA := $(shell git rev-parse --short=8 HEAD)

# Main
BINARY_NAME := exe
MAIN := ./cmd/app/main.go

# Toolchain
GO := GO111MODULE=on GOPROXY=https://goproxy.cn,direct go
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')

# Container
DOCKER := docker
DOCKERFILE := ci/docker/Dockerfile
REGISTRY := harbor.leryn.top/leryn
IMAGE_NAME := homepage-backend
FULL_IMAGE_NAME = $(REGISTRY)/$(IMAGE_NAME):$(VERSION)

##@ General

.PHONY: help
help: ## Print help info.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Developement

.PHONY: install
install: ## Install dependencies.
	$(GO) get -d -v ./...

.PHONY: check
check: ## Check
	$(GO) vet ./...

.PHONY: fmt
fmt: ## Format against code.
	$(GO) fmt ./...

.PHONY: clean
clean: ## Clean target artifact.
	$(GO) clean -r -x

.PHONY: unittest
unittest: ## Run all the unit tests.
	$(GO) test ./...

##@ Build

.PHONY: build
build: ## Build target artifact.
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -ldflags '-extldflags "-static"' -o target/$(BINARY_NAME) $(MAIN)

.PHONY: docker-build
docker-build: ## Build docker image.
	./ci/docker/docker-build.sh
