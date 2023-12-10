package entrypoints

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/nebulaclouds/nebula/nebulaadmin/plugins"
	"github.com/nebulaclouds/nebula/nebulastdlib/config"
	"github.com/nebulaclouds/nebula/nebulastdlib/config/viper"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
)

var (
	cfgFile        string
	kubeMasterURL  string
	configAccessor = viper.NewAccessor(config.Options{})
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "nebulaadmin",
	Short: "Fill in later",
	Long: `
To get started run the serve subcommand which will start a server on localhost:8088:

    nebulaadmin serve

Then you can hit it with the client:

    nebulaadmin adminservice foo bar baz

Or over HTTP 1.1 with curl:
    curl -X POST http://localhost:8088/api/v1/projects'
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initConfig(cmd.Flags())
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(pluginRegistry *plugins.Registry) error {
	pluginRegistryStore.Store(pluginRegistry)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func init() {
	// allows `$ nebulaadmin --logtostderr` to work
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Add persistent flags - persistent flags persist through all sub-commands
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./nebulaadmin_config.yaml)")
	RootCmd.PersistentFlags().StringVar(&kubeMasterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")

	RootCmd.AddCommand(viper.GetConfigCommand())

	// Allow viper to read the value of the flags
	configAccessor.InitializePflags(RootCmd.PersistentFlags())

	err := flag.CommandLine.Parse([]string{})
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initConfig(flags *pflag.FlagSet) error {
	configAccessor = viper.NewAccessor(config.Options{
		SearchPaths: []string{cfgFile, ".", "/etc/nebula/config", "$GOPATH/src/github.com/nebulaclouds/nebulaadmin"},
		StrictMode:  false,
	})

	logger.Infof(context.TODO(), "Using config file: %v", configAccessor.ConfigFilesUsed())

	configAccessor.InitializePflags(flags)

	return configAccessor.UpdateConfig(context.TODO())
}
