
.PHONY: docker_build
docker_build:
	IMAGE_NAME=$$REPOSITORY ./boilerplate/nebula/docker_build/docker_build.sh

.PHONY: dockerhub_push
dockerhub_push:
	IMAGE_NAME=nebulaclouds/$$REPOSITORY REGISTRY=docker.io ./boilerplate/nebula/docker_build/docker_build.sh
