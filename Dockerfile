FROM golang:1.16-alpine

WORKDIR /src

ADD . /src

RUN go mod tidy
