#!/usr/bin/env bash

pwd=$(pwd)

cd ..

env GOOS=linux \
GOARCH=amd64  \
go build -o docker/faketsdb .
