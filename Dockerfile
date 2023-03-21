FROM golang:1.20-alpine AS build_base

RUN apk add --no-cache git  git gcc g++

# Set the Current Working Directory inside the container
WORKDIR /tmp

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./cmd/app/main.go

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/out/app /app/vanir

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the binary program produced by `go install`
CMD ["/app/vanir"]
