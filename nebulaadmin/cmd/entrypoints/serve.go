package entrypoints

import (
	"context"
	_ "net/http/pprof" // Required to serve application.

	"github.com/spf13/cobra"

	runtimeConfig "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/server"
	"github.com/nebulaclouds/nebula/nebulaadmin/plugins"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
	"github.com/nebulaclouds/nebula/nebulastdlib/otelutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/profutils"
)

var pluginRegistryStore = plugins.NewAtomicRegistry(plugins.NewRegistry())

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the Nebula admin server",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		// Serve profiling endpoints.
		cfg := runtimeConfig.NewConfigurationProvider()
		go func() {
			err := profutils.StartProfilingServerWithDefaultHandlers(
				ctx, cfg.ApplicationConfiguration().GetTopLevelConfig().GetProfilerPort(), nil)
			if err != nil {
				logger.Panicf(ctx, "Failed to Start profiling and Metrics server. Error, %v", err)
			}
		}()
		server.SetMetricKeys(cfg.ApplicationConfiguration().GetTopLevelConfig())

		// register otel tracer providers
		for _, serviceName := range []string{otelutils.AdminGormTracer, otelutils.AdminServerTracer} {
			if err := otelutils.RegisterTracerProvider(serviceName, otelutils.GetConfig()) ; err != nil {
				logger.Errorf(ctx, "Failed to create otel tracer provider. %v", err)
				return err
			}
		}

		return server.Serve(ctx, pluginRegistryStore.Load(), nil)
	},
}

func init() {
	// Command information
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(secretsCmd)
}
