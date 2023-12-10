#!/usr/bin/env bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

GIT_SHA=$(git rev-parse HEAD)

if ! docker images | grep "${GIT_SHA}-builder"; then
  echo "ERROR:"
  echo "nebulaadmin:${GIT_SHA}-builder image does not exist."
  echo "build the image first with BUILD_PHASE=builder make docker_build"
fi

docker tag "nebulaadmin:${GIT_SHA}-builder" "nebulaadmin:test"

docker save -o "/tmp/nebulaadmin" "nebulaadmin:test"

# stop any existing test container that might be running
docker kill dockernetes || true

# The container must start with systemd (/sbin/init) as PID 1

docker run \
  --detach \
  --rm \
  --privileged \
  --volume /var/lib/docker \
  --volume /lib/modules:/lib/modules \
  --volume ${DIR}/../..:/nebulaadmin \
  --volume /tmp/nebulaadmin:/images/nebulaadmin \
  --name dockernetes \
  --env "DOCKER_REGISTRY_USERNAME=${DOCKER_REGISTRY_USERNAME}" \
  --env "DOCKER_REGISTRY_PASSWORD=${DOCKER_REGISTRY_PASSWORD}" \
  --env "DOCKERNETES_DEBUG=${DOCKERNETES_DEBUG}" \
  lyft/dockernetes:1.10.1-v0.1 /sbin/init

# wait for the system to initialize, then run execute.sh
docker exec \
  -it \
  dockernetes /nebulaadmin/script/integration/k8s/main.sh
