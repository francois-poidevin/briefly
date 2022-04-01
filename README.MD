# Briefly

## Introduction
This repository is part of the project briefly.
Briefly project is a URL shortcoder.

This repository concern the gRPC server to generate ID of a received URL and a answer an URL link to an ID (previously associated to an ID) 

## Compilation
### Protobuf
#### Prototool (UBER)
[Prototool](https://github.com/uber/prototool) is neccessary to generate protobuf

Info: Into the pipeline to migrate to [buf](https://github.com/bufbuild/buf)
##### Installation
```shell
$ GO111MODULE=on go get github.com/uber/prototool/cmd/prototool@dev
```
### Mage compilation
#### Pre-requisite
Please install [mage](https://github.com/magefile/mage)

#### Dependancies / Generate artifact
```shell
$ mage
```
An artifact named 'briefly' is generated in './bin/' folder

## Usage
### Generate
```shell
$ ./bin/briefly config new
```

will generate a default configuration file

```toml
###############################
# Briefly Settings 
###############################
[Briefly]

  ###############################
  # GRPC API settings 
  ###############################
  [Briefly.GRPC]

    # On which port REST HTTP service will listen
    listenPort = 5556

###############################
# Logs Settings 
###############################
[Log]

  # Allow to display logs in Json format if true
  jsonformatter = false

  # Log level: trace, debug, info, warning, error, panic, and fatal
  level = "warning"

```

store it as toml file, in example in 'conf' folder onto './conf/briefly.conf.toml' file

### Run server
```shell
$ ./bin/briefly startGrpc --config ./conf/briefly.conf.toml
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
