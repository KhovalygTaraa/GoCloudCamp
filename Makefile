all: build test

build:
	GOPROXY=direct go get -d ./...
	go build -o ./bin/music_service
test:
	go test ./...
docker:
	docker-compose up -d --build
proto:
	protoc -go_out=. --go_opt=paths=import \
    --go-grpc_out=. --go-grpc_opt=paths=import \
    proto/main.proto