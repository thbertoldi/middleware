native_run:
	go run middleware.go

native_build:
	go build middleware.go

docker_build:
	docker build . -t middleware:latest

docker_run:
	docker-compose up
