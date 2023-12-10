package main

import (
	"fmt"
	"os"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/nebulaclouds/nebula/nebulapropeller/cmd/kubectl-nebula/cmd"
	"github.com/nebulaclouds/nebula/nebulastdlib/contextutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils/labeled"
)

func init() {
	labeled.SetMetricKeys(contextutils.ProjectKey, contextutils.DomainKey, contextutils.WorkflowIDKey,
		contextutils.TaskIDKey)
}

func main() {

	rootCmd := cmd.NewNebulaCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
