FROM node:21-alpine AS builder

WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm ci
COPY . .
RUN npm run build

FROM node:21-alpine AS production

WORKDIR /app
COPY --from=builder /app/build ./build
COPY --from=builder /app/node_modules ./node_modules
COPY package.json .
CMD ["node", "./build"]

#
# server-builder (CGO is required)
#

FROM golang:alpine as server-builder

ENV CGO_ENABLED=1
RUN apk add build-base

WORKDIR /app

COPY backend /app/backend
COPY go.* /app
COPY *.go /app

RUN go build -buildvcs=false -mod=readonly -v