FROM node:lts-alpine AS client-build
RUN apk add python make gcc g++
WORKDIR /github.com/traPtitech/booQ/client
COPY ./client/package*.json ./
RUN npm ci
COPY ./client .
RUN npm run build

FROM golang:1.15.8-alpine AS server-build
RUN apk add --update --no-cache git
WORKDIR /go/src/github.com/traPtitech/booQ
COPY ./go.* ./
RUN go mod download
COPY . .
RUN go build

FROM alpine:3.12.0
WORKDIR /app
RUN apk --update add tzdata ca-certificates && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
COPY --from=server-build /go/src/github.com/traPtitech/booQ ./
COPY --from=client-build /github.com/traPtitech/booQ/client/dist ./client/dist
ENTRYPOINT ./booQ
