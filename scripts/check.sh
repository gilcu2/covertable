#!/bin/bash

staticcheck ./...
golangci-lint run ./...

