##### Translate to: [简体中文](README_zh.md)

# About bedrock

## Bedrock
Bedrock is a simple tool for creating go microservice.

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

