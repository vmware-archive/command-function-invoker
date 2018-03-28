# Command Function Invoker [![Build Status](https://travis-ci.org/projectriff/command-function-invoker.svg?branch=master)](https://travis-ci.org/projectriff/command-function-invoker)

## Purpose

The *command function invoker* provides a Docker base layer for a function consisting of a single command.
It accepts gRPC requests, invokes the command for each request in the input stream,
and sends the command's output to the stream of gRPC responses.

## Install as a riff invoker

```bash
riff invokers apply -f command-invoker.yaml
```

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

If you need to re-compile the protobuf protocol, install:

- Google's [protocol compiler](https://github.com/google/protobuf)

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

To attach a [delve capable](https://github.com/derekparker/delve/blob/master/Documentation/EditorIntegration.md) debugger (such as Goland)
to a `command-function-invoker` running _inside_ k8s:

```bash
make debug-dockerize
```

Then expose the `2345` port as a service, using `riff/config/delve/command-function-invoker-debug.yaml`:

```bash
riff invokers apply -f riff/config/delve/command-function-invoker-debug.yaml
```

Finally, update the function you would like to debug so that it picks up the new base layer.
Then you can connect the debugger through port `30110`.

### Compiling the Protocol

The gRPC protocol for the command function invoker is defined in [function.proto](https://github.com/projectriff/function-proto/blob/master/function.proto).

Clone https://github.com/projectriff/function-proto and set `$FN_PROTO_PATH` to point at the cloned directory. Then issue:

```bash
make gen-proto
```

## riff Commands

- [riff init command](#riff-init-command)
- [riff create command](#riff-create-command)

<!-- riff-init -->

### riff init command

Initialize a command function

#### Synopsis

Generate the function based on the executable command specified as the filename, using the name
and version specified for the function image repository and tag.

For example, from a directory named 'echo' containing a function 'echo.sh', you can simply type :

    riff init command -f echo

  or

    riff init command

to generate the resource definitions using sensible defaults.


```
riff init command [flags]
```

#### Options

```
  -h, --help                     help for command
      --invoker-version string   the version of invoker to use when building containers (default "0.0.6-snapshot")
```

#### Options inherited from parent commands

```
  -a, --artifact string      path to the function artifact, source code or jar file
      --config string        config file (default is $HOME/.riff.yaml)
      --dry-run              print generated function artifacts content to stdout only
  -f, --filepath string      path or directory used for the function resources (defaults to the current directory)
      --force                overwrite existing functions artifacts
  -i, --input string         the name of the input topic (defaults to function name)
  -n, --name string          the name of the function (defaults to the name of the current directory)
  -o, --output string        the name of the output topic (optional)
  -u, --useraccount string   the Docker user account to be used for the image repository (default "current OS user")
  -v, --version string       the version of the function image (default "0.0.1")
```

#### SEE ALSO

* [riff init](https://github.com/projectriff/riff/blob/master/riff-cli/docs/riff_init.md)	 - Initialize a function


<!-- /riff-init -->

<!-- riff-create -->

### riff create command

Create a command function

#### Synopsis

Create the function based on the executable command specified as the filename, using the name
and version specified for the function image repository and tag.

For example, from a directory named 'echo' containing a function 'echo.sh', you can simply type :

    riff create command -f echo

  or

    riff create command

to create the resource definitions, and apply the resources, using sensible defaults.


```
riff create command [flags]
```

#### Options

```
  -h, --help                     help for command
      --invoker-version string   the version of invoker to use when building containers (default "0.0.6-snapshot")
      --namespace string         the namespace used for the deployed resources (defaults to kubectl's default)
      --push                     push the image to Docker registry
```

#### Options inherited from parent commands

```
  -a, --artifact string      path to the function artifact, source code or jar file
      --config string        config file (default is $HOME/.riff.yaml)
      --dry-run              print generated function artifacts content to stdout only
  -f, --filepath string      path or directory used for the function resources (defaults to the current directory)
      --force                overwrite existing functions artifacts
  -i, --input string         the name of the input topic (defaults to function name)
  -n, --name string          the name of the function (defaults to the name of the current directory)
  -o, --output string        the name of the output topic (optional)
  -u, --useraccount string   the Docker user account to be used for the image repository (default "current OS user")
  -v, --version string       the version of the function image (default "0.0.1")
```

#### SEE ALSO

* [riff create](https://github.com/projectriff/riff/blob/master/riff-cli/docs/riff_create.md)	 - Create a function (equivalent to init, build, apply)


<!-- /riff-create -->
