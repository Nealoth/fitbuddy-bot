.PHONY: clean build run lint default

clean:
	go clean -v
	rm -rf ./bin/

build:
	go build -o ./bin/bot -race -v ./cmd/fitbuddy
	cp -R conf bin/conf

run:
	./bin/bot

lint:
	golangci-lint run -c .golangci.yml

default: lint clean build run

.DEFAULT_GOAL := default
