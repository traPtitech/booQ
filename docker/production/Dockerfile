FROM golang:1.16.5-alpine AS server-build
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
