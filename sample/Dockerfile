# syntax = docker/dockerfile:1.0-experimental

FROM golang:1.14.4-alpine AS build

RUN --mount=type=cache,target=/var/cache/apk apk add --update git

WORKDIR /go/src/github.com/titech-cpp/sqlboiler/sample
COPY go.* ./
RUN go mod download

COPY ./ ./
RUN --mount=type=cache,target=/root/.cache/go-build \
  go build -o main -ldflags "-s -w"

FROM alpine:3.11.6 AS runtime

ENV DOCKERIZE_VERSION v0.6.1
RUN apk add --update tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /Tokyo && \
  apk del tzdata && \
  rm -rf /var/cache/apk/* && \
  mkdir -p /usr/share/zoneinfo/Asia && \
  mv /Tokyo /usr/share/zoneinfo/Asia
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz && \
  tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz && \
  rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY --from=build /go/src/github.com/titech-cpp/sqlboiler/sample/main ./

ENTRYPOINT dockerize -wait tcp://mariadb:3306 ./main