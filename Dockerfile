FROM golang:1.7.4-alpine

VOLUME /go/src/bench-fio
WORKDIR /go/src/bench-fio

ADD gophercolor.png /tmp

ENTRYPOINT go test -bench .
