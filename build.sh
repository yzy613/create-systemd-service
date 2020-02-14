#!/bin/bash


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o create-systemd-service ./main.go
tar -czvf linux_amd64.tar.gz create-systemd-service
rm -f create-systemd-service

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o create-systemd-service.exe ./main.go
tar -czvf windows_amd64.tar.gz create-systemd-service.exe
rm -f create-systemd-service.exe