FROM node:lts-alpine
COPY package*.json ./
RUN npm ci
COPY ./client .
RUN npm run build

FROM golang:1.14-alpine AS build
RUN apk add --update --no-cache git
WORKDIR /go/src/github.com/traPtitech/booQ
COPY ./go.* ./
RUN go mod download
COPY . .
RUN go build

FROM alpine:3.11
WORKDIR /app
RUN apk --update add tzdata ca-certificates && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
COPY --from=build /go/src/github.com/traPtitech/booQ ./
ENTRYPOINT ./booQ
