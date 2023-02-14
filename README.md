# FME Server CLI

A command line interface for FME Server.

## Description

This is a command line interface that utilizes the FME Server REST API to interact with a running FME Server. It is meant to ease the pain of using the REST API by providing intuitive commands and flags for various operations on an FME Server.

## Getting Started

### Installing

* Simply download the binary for your system from the [releases](https://github.com/safesoftware/fmeserver-cli/releases) page.

### Executing program

* Execute the program to get a high level overview of each command
```
fmeserver
```
* Log in to an existing FME Server. It is recommended to generate an API token using the FME Server Web UI initially and use that to log in.
```
fmeserver login https://my-fmeserver.com --token my-token-here
```
* Your token and URL will be saved to a config file located in $HOME/.fmeserver-cli.yaml. Config file location can be overridden with the `--config` flag
* Test your credentials work
```
fmeserver info
```

For full documentation of all commands, see the [Documentation](doc/fmeserver.md).


## Development

* Run while coding:
```
go run main.go
```
* Build binary
```
go build -o fmeserver
```

A great resource for adding new structs to represent JSON returned from FME Server is this [JSON to Go converter](https://mholt.github.io/json-to-go/) which will create a Go struct for you from a JSON sample.

## Releasing a new version

There is a github action that will run when a new release is created that will build the binary for 5 different platforms and automatically add them to the release as assets.

## Acknowledgments

* Created using [cobra](https://github.com/spf13/cobra)
