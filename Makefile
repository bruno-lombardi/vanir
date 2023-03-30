run:
	ENV_FILE=.env go run cmd/app/main.go

build:
	go build -o bin/vanir cmd/app/main.go

dependencies:
	go mod download

tests:
	ENV_FILE=$(shell pwd)/.env.test go test -v ./test/... -race -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic

exclude-coverage:
	./.github/scripts/exclude-from-coverage.sh
