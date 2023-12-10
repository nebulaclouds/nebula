package main

import (
	"github.com/nebulaclouds/nebula/nebulapropeller/cmd/controller/cmd"
	_ "github.com/nebulaclouds/nebula/nebulapropeller/plugins"
)

func main() {
	cmd.Execute()
}
