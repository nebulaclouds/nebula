package main

import (
	"github.com/golang/glog"

	"github.com/nebulaclouds/nebula/nebulaadmin/cmd/entrypoints"
	"github.com/nebulaclouds/nebula/nebulaadmin/plugins"
)

func main() {
	glog.V(2).Info("Beginning Nebula Controller")
	err := entrypoints.Execute(plugins.NewRegistry())
	if err != nil {
		panic(err)
	}
}
