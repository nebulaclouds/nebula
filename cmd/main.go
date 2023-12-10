package main

import (
	"github.com/nebulaclouds/nebula/cmd/single"
	_ "github.com/nebulaclouds/nebula/nebulapropeller/plugins"
	"github.com/golang/glog"
)

func main() {
	glog.V(2).Info("Starting Nebula")
	err := single.Execute()
	if err != nil {
		panic(err)
	}
}
