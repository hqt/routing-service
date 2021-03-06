FROM golang:1.14.0
LABEL maintainer="huynhquangthao@gmail.com"

ARG BUILD_DIRECTORY=/app

RUN apt-get update

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.23.6

WORKDIR $BUILD_DIRECTORY
COPY go.mod go.sum $BUILD_DIRECTORY/
RUN GO111MODULE=on go mod download

COPY . $BUILD_DIRECTORY/
