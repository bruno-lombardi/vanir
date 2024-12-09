FROM golang:1.20-alpine AS build_base

RUN apk add --no-cache git  git gcc g++

# Set the Current Working Directory inside the container
WORKDIR /build

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./main.go
RUN go get -d github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make docs

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /build/out/app /app/vanir

# This container exposes port 3334 to the outside world
EXPOSE 3334

# Run the binary program produced by `go install`
CMD ["/app/vanir"]
