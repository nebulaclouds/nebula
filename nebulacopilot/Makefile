export REPOSITORY=nebulacopilot
include boilerplate/nebula/docker_build/Makefile
include boilerplate/nebula/golang_test_targets/Makefile

.PHONY: update_boilerplate
update_boilerplate:
	@curl https://raw.githubusercontent.com/nebulaclouds/boilerplate/master/boilerplate/update.sh -o boilerplate/update.sh
	@boilerplate/update.sh

clean:
	rm -rf bin

.PHONY: linux_compile
linux_compile: export CGO_ENABLED ?= 0
linux_compile: export GOOS ?= linux
linux_compile:
	go build -o /artifacts/nebula-copilot .

.PHONY: compile
compile:
	mkdir -p ./artifacts
	go build -o ./artifacts/nebula-copilot .

cross_compile:
	@mkdir -p ./artifacts/cross
	GOOS=linux GOARCH=amd64 go build -o ./artifacts/nebula-copilot .

.PHONY: go-tidy
go-tidy:
	go mod tidy
