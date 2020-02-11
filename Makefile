APP?=kubert
PORT?=8080
PROJECT?=github.com/bmsandoval/kubert
CONTAINER_IMAGE?=docker.io/bmsandoval/${APP}

RELEASE?=0.0.3

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CURDIR?=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

GOOS?=linux
GOARCH?=amd64

# Should only ever need to be run once
push-dev:
	docker build -f ./deployment/Dockerfile-dev -t $(CONTAINER_IMAGE):latest-dev .
	docker push $(CONTAINER_IMAGE):latest-dev

push-staging:
	docker build -f ./deployment/Dockerfile-staging -t $(CONTAINER_IMAGE):$(RELEASE)-staging .
	docker push $(CONTAINER_IMAGE):$(RELEASE)-staging

local:
	helm upgrade --install dev-${APP} ./chart/kubert

remove:
	helm delete dev-${APP}

stop:
	minikube stop

start:

	minikube start --mount-string ${CURDIR}:${CURDIR} --mount

test:
	go test -v -race ./...

#.PHONY: charts
#all: charts
#
#charts:
#	cd chart && helm package kubert/
#	mv chart/*.tgz docs/
##	helm repo index docs --url https://alexellis.github.io/kubert/ --merge ./docs/index.yaml

