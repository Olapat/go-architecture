ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /go-api-servive

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./go-api-servive-app .

FROM alpine:latest

RUN apk update && apk add ca-certificates && apk add --no-cache tzdata && rm -rf /var/cache/apk/*

RUN mkdir -p /go-api-servive
WORKDIR /go-api-servive
COPY --from=builder /go-api-servive/go-api-servive-app .
COPY .env .

ENV TZ=Asia/Bangkok

ENTRYPOINT ["./go-api-servive-app"]
