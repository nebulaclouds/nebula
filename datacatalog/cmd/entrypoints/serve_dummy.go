package entrypoints

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/nebulaclouds/nebula/datacatalog/pkg/config"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/rpc/datacatalogservice"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
)

var serveDummyCmd = &cobra.Command{
	Use:   "serve-dummy",
	Short: "Launches the Data Catalog server without any connections",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		cfg := config.GetConfig()

		// serve a http healthcheck endpoint
		go func() {
			err := datacatalogservice.ServeHTTPHealthCheck(ctx, cfg)
			if err != nil {
				logger.Errorf(ctx, "Unable to serve http", cfg.GetGrpcHostAddress(), err)
			}
		}()

		return datacatalogservice.Serve(ctx, cfg)
	},
}

func init() {
	RootCmd.AddCommand(serveDummyCmd)
}
