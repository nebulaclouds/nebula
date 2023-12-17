#!/bin/bash

set -e

tools=(
  "github.com/EngHabu/mockery/cmd/mockery"
  "github.com/nebulaclouds/nebulastdlib/cli/pflags@latest"
  "github.com/golangci/golangci-lint/cmd/golangci-lint"
  "github.com/daixiang0/gci"
  "github.com/alvaroloes/enumer"
  "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)

tmp_dir=$(mktemp -d -t gotooling-XXX)
echo "Using temp directory ${tmp_dir}"
cp -R boilerplate/nebula/golang_support_tools/* $tmp_dir
pushd "$tmp_dir"

for tool in "${tools[@]}"
do
    echo "Installing ${tool}"
    GO111MODULE=on go install $tool
done

popd
