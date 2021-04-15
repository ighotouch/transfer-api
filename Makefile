PWD = $(shell pwd -L)
IMAGE_NAME = ighotouch/go-bank-transfer
DOCKER_RUN = docker run --rm -it -w /app -v ${PWD}:/app -v ${GOPATH}/pkg/mod/cache:/go/pkg/mod/cache golang:1.14-stretch

init:
	cp .env.example .env



run:
	docker-compose up

down:
	docker-compose down


.PHONY:ci coverage-report enter-container logs down up vet fmt init code-review test test-local build build-image