run:
	go run cmd/app/main.go

build:
	go build -o bin/vanir cmd/app/main.go

dependencies:
	go mod download

test:
	go test -v ./...