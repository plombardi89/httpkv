SHELL := /usr/bin/env bash

GIT_COMMIT=$(shell git rev-parse --verify HEAD)

GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GOBUILD = go build -o bin/$(BINARY_BASENAME)-$(GOOS)-$(GOARCH)

BINARY_BASENAME=httpkv

all: clean fmt test build

build:
	$(GOBUILD) main.go
	ln -sf $(BINARY_BASENAME)-$(GOOS)-$(GOARCH) bin/$(BINARY_BASENAME)

build.image:
	docker build \
	-t plombardi89/httpkv \
	-t plombardi89/httpkv:$(GIT_COMMIT) \
	-t quay.io/plombardi89/httpkv \
	-t quay.io/plombardi89/httpkv:$(GIT_COMMIT) \
	-f Dockerfile \
	.

clean:
	rm -rf bin

fmt:
	go fmt ./...

test:
	go test -v ./...
