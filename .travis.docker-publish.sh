#!/bin/bash		
		
set -o errexit
set -o nounset
set -o pipefail

version=`cat VERSION`

TAG="${version}" make dockerize
docker tag "projectriff/shell-function-invoker:latest" "projectriff/shell-function-invoker:${version}-ci-${TRAVIS_COMMIT}"

docker login -u "${DOCKER_USERNAME}" -p "${DOCKER_PASSWORD}"
docker push "projectriff/shell-function-invoker"