run:
	ENV_FILE=.env go run main.go

build:
	go build -o bin/vanir main.go

dependencies:
	go mod download

tests:
	ENV_FILE=$(shell pwd)/.env.test go test -v ./... -race -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic

docs:
	swag init --parseInternal --parseDependency --output internal/app/docs