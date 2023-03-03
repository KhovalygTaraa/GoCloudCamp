.PHONY: all test build docker proto

all: build test

build:
	go get -d ./...
	go build -o ./bin/music_service

test:
	go test ./...

docker:
	docker-compose up -d --build
	docker logs -tf music_service
	
proto:
	protoc --go_out=. --go_opt=paths=import \
    --go-grpc_out=. --go-grpc_opt=paths=import \
    proto/main.proto