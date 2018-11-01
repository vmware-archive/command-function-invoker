#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

version=`cat VERSION`

gcloud auth activate-service-account --key-file <(echo $GCLOUD_CLIENT_SECRET | base64 --decode)

gsutil cp -a public-read command-function-invoker.tgz gs://projectriff/command-function-invoker/command-function-invoker-linux-amd64-${version/snapshot/$TRAVIS_COMMIT}.tgz
gsutil cp -a public-read command-function-invoker.tgz gs://projectriff/command-function-invoker/command-function-invoker-linux-amd64-latest.tgz
