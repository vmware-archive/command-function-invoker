# Command Function Invoker [![Build Status](https://travis-ci.com/projectriff/command-function-invoker.svg?branch=master)](https://travis-ci.com/projectriff/command-function-invoker)

## Purpose

The *command function invoker* provides a host for functions implemented
as a single executable command (be it a shell script or a binary).
It accepts HTTP requests and invokes the command for each request.

Communication with the function is done via `stdin` and `stdout`.
Functions can log by writing to `stderr`. 


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

To run tests:

```bash
make test
```
