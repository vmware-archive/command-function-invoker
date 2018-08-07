# Command Function Invoker [![Build Status](https://travis-ci.org/projectriff/command-function-invoker.svg?branch=master)](https://travis-ci.org/projectriff/command-function-invoker)

## Purpose

The *command function invoker* provides a Docker base layer for a function consisting of a single command.
It accepts HTTP requests, invokes the command for each request.


## Development

### Prerequisites

The following tools are required to build this project:

- `make`
- Docker
- [Glide](https://github.com/Masterminds/glide#install) for dependency management

If you intend to re-generate mocks for testing, install:

- [Mockery](https://github.com/vektra/mockery#installation)

If you would like to run tests using the `ginkgo` command, install:

- [Ginkgo](http://onsi.github.io/ginkgo/)

### Get the source

```bash
cd $(go env GOPATH)   #defaults to ~/go
git clone -o upstream https://github.com/projectriff/command-function-invoker src/github.com/projectriff/command-function-invoker
```

### Building

To build locally (this will produce a binary named `command-function-invoker` on _your_ machine):

```bash
make build
```

To build the Docker base layer:

```bash
make dockerize
```

This assumes that your docker client is correctly configured to target the daemon where you want the image built.

To run tests:

```bash
make test
```
