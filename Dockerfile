FROM golang:1.6

MAINTAINER Akash

ADD . /snabar_core

WORKDIR /snabar_core/src/shopper

RUN pwd

ENV GOPATH=/snabar_core/

RUN go get

RUN cd /snabar_core && go get github.com/axw/gocov && go get github.com/AlekSi/gocov-xml

RUN cd /snabar_core/bin && ./gocov test ../src/shopper | ./gocov-xml > ../coverage.xml