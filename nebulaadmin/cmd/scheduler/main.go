package main

import (
	"github.com/golang/glog"

	"github.com/nebulaclouds/nebula/nebulaadmin/cmd/scheduler/entrypoints"
)

func main() {
	glog.V(2).Info("Beginning Nebula Scheduler")
	err := entrypoints.Execute()
	if err != nil {
		panic(err)
	}
}
