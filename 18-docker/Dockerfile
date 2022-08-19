FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /my-app


WORKDIR /my-app

COPY ./* /my-app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o my-app .

CMD ./my-app