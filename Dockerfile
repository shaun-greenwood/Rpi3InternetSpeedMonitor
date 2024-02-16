# syntax=docker/dockerfile:1

FROM golang:1.22.0-alpine3.19

WORKDIR /app

COPY golang/go.mod golang/go.sum ./

RUN go mod download

COPY golang/main.go ./

RUN go build main.go

EXPOSE 8081

CMD ["./main"]


