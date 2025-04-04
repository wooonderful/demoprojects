#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fileserver_linux  main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o fileserver_windows.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o fileserver_darwin_amd64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o fileserver_darwin_arm64 main.go

CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o fileserver_linux_arm main.go
