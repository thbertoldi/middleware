native_run:
	go run middleware.go

native_build:
	go build -ldflags "-X main.version=$(shell git describe --tags)" middleware.go

docker_build: native_build
	docker build . -t middleware:latest

docker_run:
	docker-compose up
