FROM golang:1.20.0-alpine AS build
ENV CGO_ENABLED=0
ENV DOCKERIZE_VERSION v0.6.1
RUN apk add --update --no-cache git && \
  apk --update add tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apk del tzdata && \
  rm -rf /var/cache/apk/*

WORKDIR /app
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]
