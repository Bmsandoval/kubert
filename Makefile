APP?=kubert
PORT?=8080
PROJECT?=github.com/bmsandoval/kubert
CONTAINER_IMAGE?=docker.io/bmsandoval/${APP}

RELEASE?=0.0.3

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOOS?=linux
GOARCH?=amd64

container:
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

local: push
	helm upgrade --install dev-${APP} ./chart/kubert

remove:
	helm delete dev-${APP}

test:
	go test -v -race ./...

#.PHONY: charts
#all: charts
#
#charts:
#	cd chart && helm package kubert/
#	mv chart/*.tgz docs/
##	helm repo index docs --url https://alexellis.github.io/kubert/ --merge ./docs/index.yaml

