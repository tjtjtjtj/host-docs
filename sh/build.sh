#!/usr/bin/env bash

HASH=$(git rev-parse --verify HEAD)
BUILDDATE=$(date '+%Y/%m/%d %H:%M:%S %Z')
GOVERSION=$(go version)

go build -ldflags "-s -w -X main.hash=${HASH} -X \"main.builddate=${BUILDDATE}\" -X \"main.goversion=${GOVERSION}\""

pwd

ls -lR
