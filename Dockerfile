FROM golang:1.6

MAINTAINER Akash

ADD . /snabar_core

WORKDIR /snabar_core/src/shopper

RUN pwd

ENV GOPATH=/snabar_core/

RUN go get

RUN go test

