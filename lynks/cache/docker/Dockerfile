FROM golang:1.20-alpine

WORKDIR /cache
ENV GO111MODULE=on
COPY go.mod go.mod
RUN go mod tidy
#RUN go build -o /go-docker-demo
EXPOSE 8081