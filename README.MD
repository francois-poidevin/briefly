# Briefly

## Introduction
This repository is part of the project briefly.
Briefly project is a URL shortcoder.

This repository concern the gRPC server to encode/decode URL 

## Compilation

### Dependancies
```shell
$ go mod vendor
```

### Protobuf
#### Prototool (UBER)
Prototool is neccessary to compile
##### Installation
```shell
$ GO111MODULE=on go get github.com/uber/prototool/cmd/prototool@dev
```
##### Code generation
The following command allow to generate (based on a prototool yaml file) the golang generated code interface 
```shell
$ prototool all pkg/protocol
```
### Bin generation
```shell
$ go build -o bin/briefly
```
## Usage
### Run server
```shell
$ ./bin/briefly startGrpc
```
### Call endpoints
#### GetShortCodeHash endpoint
```shell
$ prototool grpc --address localhost:5556 --method briefly.schema.v1.SchemaAPI/GetShortCodeHash --data '{"url":"test"}'
```
```json
{"hash":"13606376228"}
```
#### GetShortCodeHash endpoint
```shell
$ prototool grpc --address localhost:5556 --method briefly.schema.v1.SchemaAPI/GetUnShortCodedURL --data '{"hash":"test"}'
```
```json
{"url":"https://www.google.com"}
```
