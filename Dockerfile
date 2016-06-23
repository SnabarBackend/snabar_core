FROM golang:1.6

MAINTAINER Akash

EXPOSE 8081 8081

ADD . /snabar_service

WORKDIR /snabar_service/src

ENV GOPATH=/snabar_service/

RUN cd /snabar_service/src/main/

RUN go get

RUN go test ../..

CMD ["go","run", "/snabar_service/src/main/route.go"]