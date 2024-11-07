FROM golang:1.23.0-alpine3.20

RUN apk update && apk add --no-cache git build-base

WORKDIR /workspace/cqrs-example-go

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

ENTRYPOINT ["air", "-c", ".air.toml"]
