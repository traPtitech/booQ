FROM node:11.11.0-alpine AS build
WORKDIR /app
RUN apk add --no-cache git
COPY package*.json ./
RUN npm ci
COPY . .
ENTRYPOINT npm run serve