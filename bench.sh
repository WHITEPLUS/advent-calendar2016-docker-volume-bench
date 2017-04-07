#! /bin/bash -Ceu

[ ! -f ./gophercolor.png ] && \
curl -O https://golang.org/doc/gopher/gophercolor.png

echo \#consistent
docker run -v `pwd`:/go/src/bench-fio:consistent `docker build --no-cache -q .`
echo \#cached
docker run -v `pwd`:/go/src/bench-fio:cached `docker build --no-cache -q .`
echo \#delegated
docker run -v `pwd`:/go/src/bench-fio:delegated `docker build --no-cache -q .`
