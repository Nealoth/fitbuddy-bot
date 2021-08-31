.PHONY: clean build run lint local-run update container-run

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

update:
	git pull

container-run: update clean build run
local-run: lint build run

.DEFAULT_GOAL := local-run
