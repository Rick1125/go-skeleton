PROG ?= app
VERSION = $(shell git rev-parse --short HEAD)
REPO = $(PROG)
DOCKER_PATH = docker.senal.club/rhizen

all:

install:
	CGO_ENABLED=0 go install -a -installsuffix cgo -ldflags "-X main.Version=${VERSION}" ./cmd/...

test:
	go test -v -bench=. ./test

build/$(REPO):
	git clone git@bitbucket.org:rhizen/${REPO}.git $@

image: build/$(REPO)
	docker build --build-arg program=${PROG} -t ${DOCKER_PATH}/${REPO}/${PROG}:${VERSION} ./build
	docker tag ${DOCKER_PATH}/${REPO}/${PROG}:${VERSION} ${DOCKER_PATH}/${REPO}/${PROG}:latest
	@rm -fr build/${REPO}

.PHONY: test image

