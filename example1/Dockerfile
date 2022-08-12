ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

COPY . /example1

WORKDIR /example1

RUN go mod download

RUN go build -o ./example1 ./main.go

EXPOSE 8080

ENTRYPOINT ["./example1"]