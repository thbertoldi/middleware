native_run:
	go run middleware.go

native_build:
	go build middleware.go

docker_build: native_build
	docker build . -t middleware:latest

docker_run: native_build
	docker-compose up
