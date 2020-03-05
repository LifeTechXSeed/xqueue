-include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build test run

run: 
	go run main.go

build: 
	go build -o ./bin/xqueue