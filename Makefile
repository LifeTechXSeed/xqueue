-include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build test run

run: 
	go run main.go

build: 
	go build -o ./bin/xqueue

test: 
	ENV=Test go test -v ./test -parallel 1