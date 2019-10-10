FROM golang:1.12.5-alpine AS build
ENV GO111MODULE=on
RUN apk add --update --no-cache git nodejs-npm

WORKDIR /go/src/github.com/traPtitech/booQ/client
COPY ./client .
RUN npm ci
ARG api_endpoint
ARG api_client_id
ARG activity_channel_id
ARG equipment_channel_id
ENV VUE_APP_API_ENDPOINT=$api_endpoint \
  VUE_APP_API_CLIENT_ID=$api_client_id \
  VUE_APP_ACTIVITY_CHANNEL_ID=$activity_channel_id \
  VUE_APP_EQUIPMENT_CHANNEL_ID=$equipment_channel_id
RUN npm run build

WORKDIR /go/src/github.com/traPtitech/booQ
COPY ./go.* ./
RUN go mod download
COPY . .
RUN go build

FROM alpine:3.9
WORKDIR /app
RUN apk --update add tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
COPY --from=build /go/src/github.com/traPtitech/booQ ./
ENTRYPOINT ./booQ
