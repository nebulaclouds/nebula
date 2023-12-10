export REPOSITORY=nebulaadmin
export NEBULA_SCHEDULER_REPOSITORY=nebulascheduler
include boilerplate/nebula/docker_build/Makefile
include boilerplate/nebula/golang_test_targets/Makefile
include boilerplate/nebula/end2end/Makefile

GIT_VERSION := $(shell git describe --always --tags)
GIT_HASH := $(shell git rev-parse --short HEAD)
TIMESTAMP := $(shell date '+%Y-%m-%d')
# TODO(monorepo): Do we need to change this? This is used in the service that provides a version.
PACKAGE ?=github.com/nebulaclouds/nebulastdlib

LD_FLAGS="-s -w -X $(PACKAGE)/version.Version=$(GIT_VERSION) -X $(PACKAGE)/version.Build=$(GIT_HASH) -X $(PACKAGE)/version.BuildTime=$(TIMESTAMP)"

.PHONY: update_boilerplate
update_boilerplate:
	@curl https://raw.githubusercontent.com/nebulaclouds/boilerplate/master/boilerplate/update.sh -o boilerplate/update.sh
	@boilerplate/update.sh

.PHONY: docker_build_scheduler
docker_build_scheduler:
	docker build -t $$NEBULA_SCHEDULER_REPOSITORY:$(GIT_HASH) -f Dockerfile.scheduler .

.PHONY: integration
integration: export CGO_ENABLED ?= 0
integration: export GOFLAGS ?= -count=1
integration:
	go test -v -tags=integration ./tests/...

.PHONY: k8s_integration
k8s_integration:
	@script/integration/launch.sh

.PHONY: k8s_integration_execute
k8s_integration_execute:
	@script/integration/execute.sh

.PHONY: compile
compile:
	mkdir -p ./bin
	go build -o bin/nebulaadmin -ldflags=$(LD_FLAGS) ./cmd/

.PHONY: compile_debug
compile_debug:
	go build -o nebulaadmin -gcflags='all=-N -l' ./cmd/ && mv ./nebulaadmin ${GOPATH}/bin


.PHONY: compile_scheduler
compile_scheduler:
	mkdir -p ./bin
	go build -o bin/nebulascheduler -ldflags=$(LD_FLAGS) ./cmd/scheduler/

.PHONY: compile_scheduler_debug
compile_scheduler_debug:
	go build -o nebulascheduler -gcflags='all=-N -l' ./cmd/scheduler/ && mv ./nebulascheduler ${GOPATH}/bin


.PHONY: linux_compile
linux_compile: export CGO_ENABLED ?= 0
linux_compile: export GOOS ?= linux
linux_compile:
	go build -o /artifacts/nebulaadmin -ldflags=$(LD_FLAGS) ./cmd/

.PHONY: linux_compile_scheduler
linux_compile_scheduler: export CGO_ENABLED ?= 0
linux_compile_scheduler: export GOOS ?= linux
linux_compile_scheduler:
	go build -o /artifacts/nebulascheduler -ldflags=$(LD_FLAGS) ./cmd/scheduler/


.PHONY: server
server:
	go run cmd/main.go serve  --server.kube-config ~/.kube/config  --config nebulaadmin_config.yaml

.PHONY: scheduler
scheduler:
	go run scheduler/main.go run  --server.kube-config ~/.kube/config  --config nebulaadmin_config.yaml


.PHONY: migrate
migrate:
	go run cmd/main.go migrate run --server.kube-config ~/.kube/config  --config nebulaadmin_config.yaml

.PHONY: seed_projects
seed_projects:
	go run cmd/main.go migrate seed-projects project admintests nebulakit --server.kube-config ~/.kube/config  --config nebulaadmin_config.yaml

all: compile

.PHONY: go-tidy
go-tidy:
	go mod tidy
