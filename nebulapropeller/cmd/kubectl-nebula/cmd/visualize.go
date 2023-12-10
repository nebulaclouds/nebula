package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/visualize"
)

type VisualizeOpts struct {
	*RootOptions
}

func NewVisualizeCommand(opts *RootOptions) *cobra.Command {

	vizOpts := &VisualizeOpts{
		RootOptions: opts,
	}

	visualizeCmd := &cobra.Command{
		Use:   "visualize <workflow_name>",
		Short: "Get GraphViz dot-formatted output.",
		Long:  `Generates GraphViz dot-formatted output for the workflow.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			w, err := vizOpts.nebulaClient.NebulaworkflowV1alpha1().NebulaWorkflows(vizOpts.ConfigOverrides.Context.Namespace).Get(context.TODO(), name, v1.GetOptions{})
			if err != nil {
				return err
			}

			fmt.Printf("Dot-formatted: %v\n", visualize.WorkflowToGraphViz(w))
			return nil
		},
	}

	return visualizeCmd
}
