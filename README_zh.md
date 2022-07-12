##### Translate to: [English](README.md)

## Bedrock
[![Go Report Card](https://goreportcard.com/badge/github.com/bob-zou/bedrock)](https://goreportcard.com/report/github.com/bob-zou/bedrock)

Bedrock是一个简单的新建golang微服务项目的工具，可以快速的新建一个新的服务项目。

## Getting Started
### Required
- [go](https://go.dev)
- [wire](https://github.com/google/wire)

### Installing
```shell
go install github.com/bob-zou/bedrock@latest
bedrock upgrade
```

### Create a service
```shell
bedrock new helloworld
cd helloworld
go run cmd/main.go
```
