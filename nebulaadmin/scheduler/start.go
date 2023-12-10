package scheduler

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime"
	"github.com/nebulaclouds/nebula/nebulaidl/clients/go/admin"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

// StartScheduler creates and starts a new scheduler instance. This is a blocking call and will block the calling go-routine
func StartScheduler(ctx context.Context) error {
	configuration := runtime.NewConfigurationProvider()
	applicationConfiguration := configuration.ApplicationConfiguration().GetTopLevelConfig()

	// Define the schedulerScope for prometheus metrics
	schedulerScope := promutils.NewScope(applicationConfiguration.MetricsScope).NewSubScope("nebulascheduler")
	schedulerPanics := schedulerScope.MustNewCounter("initialization_panic",
		"panics encountered initializing the nebula native scheduler")

	defer func() {
		if err := recover(); err != nil {
			schedulerPanics.Inc()
			logger.Fatalf(ctx, fmt.Sprintf("caught panic: %v [%+v]", err, string(debug.Stack())))
		}
	}()

	databaseConfig := configuration.ApplicationConfiguration().GetDbConfig()
	logConfig := logger.GetConfig()

	db, err := repositories.GetDB(ctx, databaseConfig, logConfig)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	dbScope := schedulerScope.NewSubScope("database")
	repo := repositories.NewGormRepo(
		db, errors.NewPostgresErrorTransformer(schedulerScope.NewSubScope("errors")), dbScope)

	clientSet, err := admin.ClientSetBuilder().WithConfig(admin.GetConfig(ctx)).Build(ctx)
	if err != nil {
		logger.Fatalf(ctx, "Nebula native scheduler failed to start due to %v", err)
		return err
	}
	adminServiceClient := clientSet.AdminClient()

	scheduleExecutor := NewScheduledExecutor(repo,
		configuration.ApplicationConfiguration().GetSchedulerConfig().GetWorkflowExecutorConfig(), schedulerScope, adminServiceClient)

	logger.Info(ctx, "Successfully initialized a native nebula scheduler")

	err = scheduleExecutor.Run(ctx)
	if err != nil {
		logger.Fatalf(ctx, "Nebula native scheduler failed to start due to %v", err)
		return err
	}
	return nil
}
