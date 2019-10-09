#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

version=$(cat VERSION)
git_sha=$(git rev-parse HEAD)

gsutil cp -a public-read command-function-invoker.tgz gs://projectriff/command-function-invoker/command-function-invoker-linux-amd64-${version}/snapshot/${git_sha}.tgz
gsutil cp -a public-read command-function-invoker.tgz gs://projectriff/command-function-invoker/command-function-invoker-linux-amd64-${version}.tgz
gsutil cp -a public-read command-function-invoker.tgz gs://projectriff/command-function-invoker/command-function-invoker-linux-amd64-latest.tgz
