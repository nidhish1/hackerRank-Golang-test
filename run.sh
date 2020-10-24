#! /bin/sh
set -e


export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED=0


## DB Configs

export DB_NAME="g97.nidhish@gmail.com"
export DB_PASS="g97.nidhish@gmail.com"
export DB_HOST="143.110.190.177"
export DB_PORT="3306"

go build -v -o dist/go-mysql-crud
