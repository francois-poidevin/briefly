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
```shell
$ prototool all pkg/protocol
```
## Usage

