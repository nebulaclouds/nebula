export REPOSITORY=nebulaplugins
include boilerplate/nebula/docker_build/Makefile
include boilerplate/nebula/golang_test_targets/Makefile

.PHONY: update_boilerplate
update_boilerplate:
	@curl https://raw.githubusercontent.com/nebulaclouds/boilerplate/master/boilerplate/update.sh -o boilerplate/update.sh
	@boilerplate/update.sh

.PHONY: go-tidy
go-tidy:
	go mod tidy
