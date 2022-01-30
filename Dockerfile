FROM golang:1.16
WORKDIR /src
ADD . /src

RUN go mod download
