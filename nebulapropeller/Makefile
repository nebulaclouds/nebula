export REPOSITORY=nebulapropeller
include boilerplate/nebula/docker_build/Makefile
include boilerplate/nebula/golang_test_targets/Makefile
include boilerplate/nebula/end2end/Makefile

.PHONY: update_boilerplate
update_boilerplate:
	@curl https://raw.githubusercontent.com/nebulaclouds/boilerplate/master/boilerplate/update.sh -o boilerplate/update.sh
	@boilerplate/update.sh

.PHONY: linux_compile
linux_compile: export CGO_ENABLED ?= 0
linux_compile: export GOOS ?= linux
linux_compile:
	go build -o /artifacts/nebulapropeller ./cmd/controller/main.go
	go build -o /artifacts/nebulapropeller-manager ./cmd/manager/main.go
	go build -o /artifacts/kubectl-nebula ./cmd/kubectl-nebula/main.go

.PHONY: compile
compile:
	mkdir -p ./bin
	go build -o bin/nebulapropeller ./cmd/controller/main.go
	go build -o bin/nebulapropeller-manager ./cmd/manager/main.go
	go build -o bin/kubectl-nebula ./cmd/kubectl-nebula/main.go

.PHONY: compile_nebulapropeller
compile_nebulapropeller:
	mkdir -p ./bin
	go build -o bin/nebulapropeller ./cmd/controller/main.go

cross_compile:
	@glide install
	@mkdir -p ./bin/cross
	go build -o bin/cross/nebulapropeller ./cmd/controller/main.go
	go build -o bin/cross/nebulapropeller-manager ./cmd/manager/main.go
	go build -o bin/cross/kubectl-nebula ./cmd/kubectl-nebula/main.go

op_code_generate:
	@RESOURCE_NAME=nebulaworkflow OPERATOR_PKG=github.com/nebulaclouds/nebulapropeller ./hack/update-codegen.sh

benchmark:
	mkdir -p ./bin/benchmark
	@go test -run=^$ -bench=. -cpuprofile=cpu.out -memprofile=mem.out ./pkg/controller/nodes/. && mv *.out ./bin/benchmark/ && mv *.test ./bin/benchmark/

# server starts the service in development mode
.PHONY: server
server:
	@go run ./cmd/controller/main.go --alsologtostderr --propeller.kube-config=$(HOME)/.kube/config

# manager starts the manager service in development mode
.PHONY: manager
manager:
	@go run ./cmd/manager/main.go --alsologtostderr --propeller.kube-config=$(HOME)/.kube/config

clean:
	rm -rf bin

# Generate golden files. Add test packages that generate golden files here.
golden:
	go test ./cmd/kubectl-nebula/cmd -update
	go test ./pkg/compiler/test -update

.PHONY: go-tidy
go-tidy:
	go mod tidy
