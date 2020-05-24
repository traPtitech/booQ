FROM golang:1.14-alpine AS build
RUN apk add --update --no-cache git nodejs-npm

WORKDIR /go/src/github.com/traPtitech/booQ/client
COPY ./client .
RUN npm ci
RUN npm run build

WORKDIR /go/src/github.com/traPtitech/booQ
COPY ./go.* ./
RUN go mod download
COPY . .
RUN go build

FROM alpine:3.9
WORKDIR /app
RUN apk --update add tzdata ca-certificates && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
COPY --from=build /go/src/github.com/traPtitech/booQ ./
ENTRYPOINT ./booQ
