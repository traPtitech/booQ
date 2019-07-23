FROM golang:1.12.5-alpine AS build
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV DOCKERIZE_VERSION v0.6.1
RUN apk add --update --no-cache git && \
  apk --update add tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apk del tzdata && \
  rm -rf /var/cache/apk/* && \
  wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz && \
  tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz && \
  rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /go/src/github.com/traPtitech/booQ
COPY ./go.* ./
RUN go mod download \
  && go get github.com/pilu/fresh

COPY . .
