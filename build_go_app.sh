#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -a -installsuffix cgo -o gs_auth gs_auth.go
cp ./gs_auth ./cgo/gs_auth
echo "gs_auth build completed and can be found in cgo folder"
