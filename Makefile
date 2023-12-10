export REPOSITORY=nebula
include boilerplate/nebula/end2end/Makefile
include boilerplate/nebula/golang_test_targets/Makefile

define PIP_COMPILE
pip-compile $(1) --upgrade --verbose --resolver=backtracking --annotation-style=line
endef

GIT_VERSION := $(shell git describe --always --tags)
GIT_HASH := $(shell git rev-parse --short HEAD)
TIMESTAMP := $(shell date '+%Y-%m-%d')
PACKAGE ?=github.com/nebulaclouds/nebulastdlib
LD_FLAGS="-s -w -X $(PACKAGE)/version.Version=$(GIT_VERSION) -X $(PACKAGE)/version.Build=$(GIT_HASH) -X $(PACKAGE)/version.BuildTime=$(TIMESTAMP)"

.PHONY: cmd/single/dist
cmd/single/dist: export NEBULACONSOLE_VERSION ?= latest
cmd/single/dist:
	script/get_nebulaconsole_dist.sh

.PHONY: compile
compile: cmd/single/dist
	go build -tags console -v -o nebula -ldflags=$(LD_FLAGS) ./cmd/
	mv ./nebula ${GOPATH}/bin || echo "Skipped copying 'nebula' to ${GOPATH}/bin"

.PHONY: linux_compile
linux_compile: cmd/single/dist
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -tags console -v -o /artifacts/nebula -ldflags=$(LD_FLAGS) ./cmd/

.PHONY: update_boilerplate
update_boilerplate:
	@boilerplate/update.sh

.PHONY: kustomize
kustomize:
	KUSTOMIZE_VERSION=3.9.2 bash script/generate_kustomize.sh

.PHONY: helm
helm: ## Generate K8s Manifest from Helm Charts.
	bash script/generate_helm.sh

.PHONY: release_automation
release_automation:
	mkdir -p release
	bash script/release.sh
	bash script/generate_config_docs.sh
	$(MAKE) -C docker/sandbox-bundled manifests

.PHONY: deploy_sandbox
deploy_sandbox: 
	bash script/deploy.sh

.PHONY: install-piptools
install-piptools: ## Install pip-tools
	pip install -U pip-tools

.PHONY: doc-requirements.txt
doc-requirements.txt: doc-requirements.in install-piptools
	$(call PIP_COMPILE,doc-requirements.in)

.PHONY: install-conda-lock
install-conda-lock:
	pip install conda-lock

.PHONY: conda-lock
conda-lock: install-conda-lock
	conda-lock -f monodocs-environment.yaml --without-cuda --lockfile monodocs-environment.lock.yaml

.PHONY: stats
stats:
	@generate-dashboard -o deployment/stats/prometheus/nebulapropeller-dashboard.json stats/nebulapropeller.dashboard.py
	@generate-dashboard -o deployment/stats/prometheus/nebulaadmin-dashboard.json stats/nebulaadmin.dashboard.py
	@generate-dashboard -o deployment/stats/prometheus/nebulauser-dashboard.json stats/nebulauser.dashboard.py

.PHONY: prepare_artifacts
prepare_artifacts:
	bash script/prepare_artifacts.sh

.PHONE: helm_update
helm_update: ## Update helm charts' dependencies.
	helm dep update ./charts/nebula/

.PHONY: helm_install
helm_install: ## Install helm charts
	helm install nebula --debug ./charts/nebula -f ./charts/nebula/values.yaml --create-namespace --namespace=nebula

.PHONY: helm_upgrade
helm_upgrade: ## Upgrade helm charts
	helm upgrade nebula --debug ./charts/nebula -f ./charts/nebula/values.yaml --create-namespace --namespace=nebula

.PHONY: docs
docs:
	make -C docs clean html SPHINXOPTS=-W

.PHONY: help
help: SHELL := /bin/sh
help: ## List available commands and their usage
	@awk 'BEGIN {FS = ":.*?##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } ' $(MAKEFILE_LIST)

.PHONY: setup_local_dev
setup_local_dev: ## Sets up k3d cluster with Nebula dependencies for local development
	@bash script/setup_local_dev.sh

# This builds the nebula binary image for whatever architecture you're on. Don't push this image to the official
# registry because we need those to be multi-arch.
.PHONY: build_native_nebula
build_native_nebula: NEBULACONSOLE_VERSION := latest
build_native_nebula:
	docker build \
	--build-arg NEBULACONSOLE_VERSION=$(NEBULACONSOLE_VERSION) \
	--tag nebula-binary:native .

.PHONY: go-tidy
go-tidy:
	go mod tidy
	make -C datacatalog go-tidy
	make -C nebulaadmin go-tidy
	make -C nebulaidl go-tidy
	make -C nebulapropeller go-tidy
	make -C nebulaplugins go-tidy
	make -C nebulastdlib go-tidy
	make -C nebulacopilot go-tidy
