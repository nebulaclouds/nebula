package single

import (
	"context"
	"net/http"
	"os"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	ctrlWebhook "sigs.k8s.io/controller-runtime/pkg/webhook"

	datacatalogConfig "github.com/nebulaclouds/nebula/datacatalog/pkg/config"
	datacatalogRepo "github.com/nebulaclouds/nebula/datacatalog/pkg/repositories"
	datacatalog "github.com/nebulaclouds/nebula/datacatalog/pkg/rpc/datacatalogservice"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/clusterresource"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/common"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime"
	adminServer "github.com/nebulaclouds/nebula/nebulaadmin/pkg/server"
	"github.com/nebulaclouds/nebula/nebulaadmin/plugins"
	adminScheduler "github.com/nebulaclouds/nebula/nebulaadmin/scheduler"
	propellerEntrypoint "github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller"
	propellerConfig "github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/config"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/signals"
	webhookEntrypoint "github.com/nebulaclouds/nebula/nebulapropeller/pkg/webhook"
	webhookConfig "github.com/nebulaclouds/nebula/nebulapropeller/pkg/webhook/config"
	"github.com/nebulaclouds/nebula/nebulastdlib/contextutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
	"github.com/nebulaclouds/nebula/nebulastdlib/otelutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/profutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils/labeled"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
	_ "github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	_ "gorm.io/driver/postgres" // Required to import database driver.
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

const defaultNamespace = "all"
const propellerDefaultNamespace = "nebula"

func startDataCatalog(ctx context.Context, _ DataCatalog) error {
	if err := datacatalogRepo.Migrate(ctx); err != nil {
		return err
	}
	catalogCfg := datacatalogConfig.GetConfig()
	return datacatalog.ServeInsecure(ctx, catalogCfg)
}

func startClusterResourceController(ctx context.Context) error {
	configuration := runtime.NewConfigurationProvider()
	scope := promutils.NewScope(configuration.ApplicationConfiguration().GetTopLevelConfig().MetricsScope).NewSubScope("clusterresource")
	clusterResourceController, err := clusterresource.NewClusterResourceControllerFromConfig(ctx, scope, configuration)
	if err != nil {
		return err
	}
	clusterResourceController.Run()
	logger.Infof(ctx, "ClusterResourceController started running successfully")
	return nil
}

func startAdmin(ctx context.Context, cfg Admin) error {
	logger.Infof(ctx, "Running Database Migrations...")
	if err := adminServer.Migrate(ctx); err != nil {
		return err
	}

	logger.Infof(ctx, "Seeding default projects...")
	projects := []string{"nebulasnacks"}
	if len(cfg.SeedProjects) != 0 {
		projects = cfg.SeedProjects
	}
	logger.Infof(ctx, "Seeding default projects...", projects)
	if err := adminServer.SeedProjects(ctx, projects); err != nil {
		return err
	}

	g, childCtx := errgroup.WithContext(ctx)

	if !cfg.DisableScheduler {
		logger.Infof(ctx, "Starting Scheduler...")
		g.Go(func() error {
			return adminScheduler.StartScheduler(childCtx)
		})
	}

	if !cfg.DisableClusterResourceManager {
		logger.Infof(ctx, "Starting cluster resource controller...")
		g.Go(func() error {
			return startClusterResourceController(childCtx)
		})
	}

	if !cfg.Disabled {
		g.Go(func() error {
			logger.Infof(ctx, "Starting Admin server...")
			registry := plugins.NewAtomicRegistry(plugins.NewRegistry())
			return adminServer.Serve(childCtx, registry.Load(), GetConsoleHandlers())
		})
	}
	return g.Wait()
}

func startPropeller(ctx context.Context, cfg Propeller) error {
	propellerCfg := propellerConfig.GetConfig()
	propellerScope := promutils.NewScope(propellerConfig.GetConfig().MetricsPrefix).NewSubScope("propeller").NewSubScope(propellerCfg.LimitNamespace)
	limitNamespace := ""
	var namespaceConfigs map[string]cache.Config
	if propellerCfg.LimitNamespace != defaultNamespace {
		limitNamespace = propellerCfg.LimitNamespace
		namespaceConfigs = map[string]cache.Config{
			limitNamespace: {},
		}
	}

	options := manager.Options{
		Cache: cache.Options{
			SyncPeriod:        &propellerCfg.DownstreamEval.Duration,
			DefaultNamespaces: namespaceConfigs,
		},
		NewCache: func(config *rest.Config, options cache.Options) (cache.Cache, error) {
			k8sCache, err := cache.New(config, options)
			if err != nil {
				return k8sCache, err
			}

			return otelutils.WrapK8sCache(k8sCache), nil
		},
		NewClient: func(config *rest.Config, options client.Options) (client.Client, error) {
			k8sClient, err := client.New(config, options)
			if err != nil {
				return k8sClient, err
			}

			return otelutils.WrapK8sClient(k8sClient), nil
		},
		Metrics: metricsserver.Options{
			// Disable metrics serving
			BindAddress: "0",
		},
		WebhookServer: ctrlWebhook.NewServer(ctrlWebhook.Options{
			CertDir: webhookConfig.GetConfig().ExpandCertDir(),
			Port:    webhookConfig.GetConfig().ListenPort,
		}),
	}

	mgr, err := propellerEntrypoint.CreateControllerManager(ctx, propellerCfg, options)
	if err != nil {
		logger.Errorf(ctx, "Failed to create controller manager. %v", err)
		return err
	}
	g, childCtx := errgroup.WithContext(ctx)

	if !cfg.DisableWebhook {
		g.Go(func() error {
			logger.Infof(childCtx, "Starting to initialize certificate...")
			err := webhookEntrypoint.InitCerts(childCtx, propellerCfg, webhookConfig.GetConfig())
			if err != nil {
				logger.Errorf(childCtx, "Failed to initialize certificates for Secrets Webhook. %v", err)
				return err
			}
			logger.Infof(childCtx, "Starting Webhook server...")
			// set default namespace for pod template store
			podNamespace, found := os.LookupEnv(webhookEntrypoint.PodNamespaceEnvVar)
			if !found {
				podNamespace = propellerDefaultNamespace
			}

			return webhookEntrypoint.Run(signals.SetupSignalHandler(childCtx), propellerCfg, webhookConfig.GetConfig(), podNamespace, &propellerScope, mgr)
		})
	}

	if !cfg.Disabled {
		g.Go(func() error {
			logger.Infof(childCtx, "Starting Nebula Propeller...")
			return propellerEntrypoint.StartController(childCtx, propellerCfg, defaultNamespace, mgr, &propellerScope)
		})
	}

	if !cfg.DisableWebhook || !cfg.Disabled {
		handlers := map[string]http.Handler{
			"/k8smetrics": promhttp.HandlerFor(metrics.Registry, promhttp.HandlerOpts{
				ErrorHandling: promhttp.HTTPErrorOnError,
			}),
		}

		g.Go(func() error {
			return profutils.StartProfilingServerWithDefaultHandlers(childCtx, propellerCfg.ProfilerPort.Port, handlers)
		})

		g.Go(func() error {
			err := propellerEntrypoint.StartControllerManager(childCtx, mgr)
			if err != nil {
				logger.Fatalf(childCtx, "Failed to start controller manager. Error: %v", err)
			}
			return err
		})
	}

	return g.Wait()
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This command will start Nebula cluster locally",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		g, childCtx := errgroup.WithContext(ctx)
		cfg := GetConfig()

		for _, serviceName := range []string{otelutils.AdminClientTracer, otelutils.AdminGormTracer, otelutils.AdminServerTracer,
			otelutils.BlobstoreClientTracer, otelutils.DataCatalogClientTracer, otelutils.DataCatalogGormTracer,
			otelutils.DataCatalogServerTracer, otelutils.NebulaPropellerTracer, otelutils.K8sClientTracer} {
			if err := otelutils.RegisterTracerProvider(serviceName, otelutils.GetConfig()); err != nil {
				logger.Errorf(ctx, "Failed to create otel tracer provider. %v", err)
				return err
			}
		}

		if !cfg.Admin.Disabled {
			g.Go(func() error {
				err := startAdmin(childCtx, cfg.Admin)
				if err != nil {
					logger.Panicf(childCtx, "Failed to start Admin, err: %v", err)
					return err
				}
				return nil
			})
		}

		if !cfg.Propeller.Disabled {
			g.Go(func() error {
				err := startPropeller(childCtx, cfg.Propeller)
				if err != nil {
					logger.Panicf(childCtx, "Failed to start Propeller, err: %v", err)
					return err
				}
				return nil
			})
		}

		if !cfg.DataCatalog.Disabled {
			g.Go(func() error {
				err := startDataCatalog(childCtx, cfg.DataCatalog)
				if err != nil {
					logger.Panicf(childCtx, "Failed to start Datacatalog, err: %v", err)
					return err
				}
				return nil
			})
		}

		return g.Wait()
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
	// Set Keys
	labeled.SetMetricKeys(contextutils.AppNameKey, contextutils.ProjectKey, contextutils.DomainKey,
		contextutils.ExecIDKey, contextutils.WorkflowIDKey, contextutils.NodeIDKey, contextutils.TaskIDKey,
		contextutils.TaskTypeKey, common.RuntimeTypeKey, common.RuntimeVersionKey, storage.FailureTypeLabel)
}
