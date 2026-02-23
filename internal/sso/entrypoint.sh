#!/bin/sh
#TODO replace to base shell script to all containers
alias migrations='go run /common/migrations/main.go'

go build -o main ./main.go

exec ./main