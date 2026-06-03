FROM golang:1.20.0-alpine@sha256:0d145ecb3cb3772ee54d3a97ae2774aa4f8a179f28f9d4ea67b9cb38b58acebd AS server-build
RUN apk add --update --no-cache git
WORKDIR /go/src/github.com/traPtitech/booQ
COPY ./go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build

FROM alpine:3.12.0
WORKDIR /app
RUN apk --update add tzdata ca-certificates && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
COPY --from=server-build /go/src/github.com/traPtitech/booQ ./
ENTRYPOINT ["/app/booQ"]
