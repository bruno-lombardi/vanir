export ENV_FILE := $(pwd)/.env

run:
	go run cmd/app/main.go

build:
	go build -o bin/vanir cmd/app/main.go

dependencies:
	go mod download

tests:
	ENV_FILE=$(shell pwd)/.env.test go test -v ./...
