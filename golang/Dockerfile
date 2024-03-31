# syntax=docker/dockerfile:1

FROM golang:1.22.0-alpine3.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN go build main.go

EXPOSE 8081

CMD ["./main"]


