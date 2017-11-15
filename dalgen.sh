#!/bin/bash
go run dalgen.go --xml=$1
gofmt -w ./samples/dal/dao/mysql_dao/*.go
gofmt -w ./samples/dal/dataobject/*.go

