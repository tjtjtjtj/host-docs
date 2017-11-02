#!/usr/bin/env bash

HASH=$(git rev-parse --verify HEAD)
BUILDDATE=$(date '+%Y/%m/%d %H:%M:%S %Z')
GOVERSION=$(go version)
BASE_DIR=$(cd $(dirname $0)/.. && pwd)

go build -ldflags "-s -w -X main.hash=${HASH} -X \"main.builddate=${BUILDDATE}\" -X \"main.goversion=${GOVERSION}\""
