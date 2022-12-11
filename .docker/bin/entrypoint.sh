#! /bin/bash

sleep 10
cd /go/src/app
go build
make migrate
./expenses