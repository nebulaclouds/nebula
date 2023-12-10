// Generates a Register method to automatically add pflags to a pflagSet for all fields in a given type.
package main

import (
	"log"

	"github.com/nebulaclouds/nebula/nebulastdlib/cli/pflags/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
