#!/bin/bash

#go test -coverpkg=./... -coverprofile=golang.out ./...
go test -coverprofile=coverage.out ./...
go tool cover -func coverage.out
#go tool cover -html golang.out

