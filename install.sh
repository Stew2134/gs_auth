#!/bin/sh

mkdir cgo
docker build -t gs_auth .
docker run -v $(pwd)/cgo:/app/cgo gs_auth
