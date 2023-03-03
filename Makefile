all: build test

build:
	GOPROXY=direct go get -d ./...
	go build -o ./bin/music_service
test:
	go test ./...
docker:
	docker-compose up -d --build