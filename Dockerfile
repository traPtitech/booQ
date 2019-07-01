FROM golang:1.12.5-alpine AS build
ENV GO111MODULE=on
RUN apk add --update --no-cache git && \
  apk --update add tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apk del tzdata && \
  rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/traPtitech/booQ
COPY ./go.* ./
RUN go mod download \
  && go get github.com/pilu/fresh

COPY . .
ENTRYPOINT fresh -c fresh.conf