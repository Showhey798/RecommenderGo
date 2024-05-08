FROM golang:1.22.3-bullseye

ENV PROTBUF_VERSION=3.19.4

RUN apt-get update \
    && apt-get install -y protobuf-compiler unzip --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* 

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.58.0

WORKDIR /app
COPY . /app

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/cosmtrek/air@latest

