#! /usr/bin/env bash

[ ! -f ./gophercolor.png ] && \
curl -O https://golang.org/doc/gopher/gophercolor.png

docker run -v `pwd`:/go/src/bench-fio `docker build --no-cache -q .`
