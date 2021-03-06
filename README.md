##### Translate to: [简体中文](README_zh.md)

[![Go Report Card](https://goreportcard.com/badge/github.com/bob-zou/bedrock)](https://goreportcard.com/report/github.com/bob-zou/bedrock)
[![GoDoc](https://godoc.org/github.com/bob-zou/bedrock?status.svg)](https://godoc.org/github.com/bob-zou/bedrock)

## Bedrock
Bedrock is a simple tool for creating go microservice.

## Getting Started
### Required
- [go](https://go.dev)
- [wire](https://github.com/google/wire)

### Installing
![installing.gif](https://cdn.jsdelivr.net/gh/bob-zou/bedrock/assets/images/installing.gif)
```shell
go install github.com/bob-zou/bedrock@latest
bedrock upgrade
```

### Create a service
```shell
bedrock new helloworld
cd helloworld

# update swagger docs
bedrock docs

# check if database is exist and password is right
cat configs/db.json
cat configs/redis.json

# start service
bedrock run
```
