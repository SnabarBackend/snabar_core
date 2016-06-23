FROM golang:1.6

MAINTAINER Akash

ADD . /snabar_core

WORKDIR /snabar_core/src/shopper

RUN pwd

ENV GOPATH=/snabar_core/

RUN go get

RUN go get -u github.com/axw/gocov

RUN go get -u github.com/AlekSi/gocov-xml

RUN cd /snabar_core/bin && ./gocov test ../src/shopper | ./gocov-xml > ../coverage.xml