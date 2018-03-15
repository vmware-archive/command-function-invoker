#!/bin/bash		
		
set -o errexit
set -o nounset
set -o pipefail

version=`cat VERSION`

TAG="${version}" make dockerize
docker tag "projectriff/command-function-invoker:latest" "projectriff/command-function-invoker:${version}-ci-${TRAVIS_COMMIT}"

docker login -u "${DOCKER_USERNAME}" -p "${DOCKER_PASSWORD}"
docker push "projectriff/command-function-invoker"