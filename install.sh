#!/usr/bin/env bash
echo 'install packages...'

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

go get -u -v github.com/garyburd/redigo/redis
go get -u -v github.com/gorilla/sessions
go get -u -v github.com/joho/godotenv/autoload
go get -u -v github.com/justinas/nosurf
go get -u -v github.com/lib/pq
go get -u -v github.com/nu7hatch/gouuid
go get -u -v gopkg.in/boj/redistore.v1
go get -u -v golang.org/x/net/context
go get -u -v gopkg.in/boj/redistore.v1

export GOPATH="$OLDGOPATH"

echo 'finished.'