package adminservice

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/golang/protobuf/proto"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/cloudevent"
	eventWriter "github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/events/implementations"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/notifications"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/schedule"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/data"
	executionCluster "github.com/nebulaclouds/nebula/nebulaadmin/pkg/executioncluster/impl"
	manager "github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/impl"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/impl/resources"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	runtimeIfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	workflowengineImpl "github.com/nebulaclouds/nebula/nebulaadmin/pkg/workflowengine/impl"
	"github.com/nebulaclouds/nebula/nebulaadmin/plugins"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/service"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

type AdminService struct {
	service.UnimplementedAdminServiceServer
	TaskManager              interfaces.TaskInterface
	WorkflowManager          interfaces.WorkflowInterface
	LaunchPlanManager        interfaces.LaunchPlanInterface
	ExecutionManager         interfaces.ExecutionInterface
	NodeExecutionManager     interfaces.NodeExecutionInterface
	TaskExecutionManager     interfaces.TaskExecutionInterface
	ProjectManager           interfaces.ProjectInterface
	ResourceManager          interfaces.ResourceInterface
	NamedEntityManager       interfaces.NamedEntityInterface
	VersionManager           interfaces.VersionInterface
	DescriptionEntityManager interfaces.DescriptionEntityInterface
	MetricsManager           interfaces.MetricsInterface
	Metrics                  AdminMetrics
}

// Intercepts all admin requests to handle panics during execution.
func (m *AdminService) interceptPanic(ctx context.Context, request proto.Message) {
	err := recover()
	if err == nil {
		return
	}

	m.Metrics.PanicCounter.Inc()
	logger.Fatalf(ctx, "panic-ed for request: [%+v] with err: %v with Stack: %v", request, err, string(debug.Stack()))
}

const defaultRetries = 3

func NewAdminServer(ctx context.Context, pluginRegistry *plugins.Registry, configuration runtimeIfaces.Configuration,
	kubeConfig, master string, dataStorageClient *storage.DataStore, adminScope promutils.Scope) *AdminService {
	applicationConfiguration := configuration.ApplicationConfiguration().GetTopLevelConfig()

	panicCounter := adminScope.MustNewCounter("initialization_panic",
		"panics encountered initializing the admin service")

	defer func() {
		if err := recover(); err != nil {
			panicCounter.Inc()
			logger.Fatalf(ctx, fmt.Sprintf("caught panic: %v [%+v]", err, string(debug.Stack())))
		}
	}()

	databaseConfig := configuration.ApplicationConfiguration().GetDbConfig()
	logConfig := logger.GetConfig()

	db, err := repositories.GetDB(ctx, databaseConfig, logConfig)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	dbScope := adminScope.NewSubScope("database")
	repo := repositories.NewGormRepo(
		db, errors.NewPostgresErrorTransformer(adminScope.NewSubScope("errors")), dbScope)
	execCluster := executionCluster.GetExecutionCluster(
		adminScope.NewSubScope("executor").NewSubScope("cluster"),
		kubeConfig,
		master,
		configuration,
		repo)
	workflowBuilder := workflowengineImpl.NewNebulaWorkflowBuilder(
		adminScope.NewSubScope("builder").NewSubScope("nebulapropeller"))
	workflowExecutor := workflowengineImpl.NewK8sWorkflowExecutor(configuration, execCluster, workflowBuilder, dataStorageClient)
	logger.Info(ctx, "Successfully created a workflow executor engine")
	pluginRegistry.RegisterDefault(plugins.PluginIDWorkflowExecutor, workflowExecutor)

	publisher := notifications.NewNotificationsPublisher(*configuration.ApplicationConfiguration().GetNotificationsConfig(), adminScope)
	processor := notifications.NewNotificationsProcessor(*configuration.ApplicationConfiguration().GetNotificationsConfig(), adminScope)
	eventPublisher := notifications.NewEventsPublisher(*configuration.ApplicationConfiguration().GetExternalEventsConfig(), adminScope)
	cloudEventPublisher := cloudevent.NewCloudEventsPublisher(ctx, *configuration.ApplicationConfiguration().GetCloudEventsConfig(), adminScope)
	go func() {
		logger.Info(ctx, "Started processing notifications.")
		processor.StartProcessing()
	}()

	// Configure workflow scheduler async processes.
	schedulerConfig := configuration.ApplicationConfiguration().GetSchedulerConfig()
	workflowScheduler := schedule.NewWorkflowScheduler(repo, schedule.WorkflowSchedulerConfig{
		Retries:         defaultRetries,
		SchedulerConfig: *schedulerConfig,
		Scope:           adminScope,
	})

	eventScheduler := workflowScheduler.GetEventScheduler()
	launchPlanManager := manager.NewLaunchPlanManager(
		repo, configuration, eventScheduler, adminScope.NewSubScope("launch_plan_manager"))

	// Configure admin-specific remote data handler (separate from storage)
	remoteDataConfig := configuration.ApplicationConfiguration().GetRemoteDataConfig()
	urlData := data.GetRemoteDataHandler(data.RemoteDataHandlerConfig{
		CloudProvider:            remoteDataConfig.Scheme,
		SignedURLDurationMinutes: remoteDataConfig.SignedURL.DurationMinutes,
		SigningPrincipal:         remoteDataConfig.SignedURL.SigningPrincipal,
		Region:                   remoteDataConfig.Region,
		Retries:                  defaultRetries,
		RemoteDataStoreClient:    dataStorageClient,
	}).GetRemoteURLInterface()

	workflowManager := manager.NewWorkflowManager(
		repo, configuration, workflowengineImpl.NewCompiler(), dataStorageClient, applicationConfiguration.GetMetadataStoragePrefix(),
		adminScope.NewSubScope("workflow_manager"))
	namedEntityManager := manager.NewNamedEntityManager(repo, configuration, adminScope.NewSubScope("named_entity_manager"))
	descriptionEntityManager := manager.NewDescriptionEntityManager(repo, configuration, adminScope.NewSubScope("description_entity_manager"))

	executionEventWriter := eventWriter.NewWorkflowExecutionEventWriter(repo, applicationConfiguration.GetAsyncEventsBufferSize())
	go func() {
		executionEventWriter.Run()
	}()

	executionManager := manager.NewExecutionManager(repo, pluginRegistry, configuration, dataStorageClient,
		adminScope.NewSubScope("execution_manager"), adminScope.NewSubScope("user_execution_metrics"),
		publisher, urlData, workflowManager, namedEntityManager, eventPublisher, cloudEventPublisher, executionEventWriter)
	versionManager := manager.NewVersionManager()

	scheduledWorkflowExecutor := workflowScheduler.GetWorkflowExecutor(executionManager, launchPlanManager)
	logger.Info(ctx, "Successfully initialized a new scheduled workflow executor")
	go func() {
		logger.Info(ctx, "Starting the scheduled workflow executor")
		scheduledWorkflowExecutor.Run()
	}()

	nodeExecutionEventWriter := eventWriter.NewNodeExecutionEventWriter(repo, applicationConfiguration.GetAsyncEventsBufferSize())
	go func() {
		nodeExecutionEventWriter.Run()
	}()

	nodeExecutionManager := manager.NewNodeExecutionManager(repo, configuration, applicationConfiguration.GetMetadataStoragePrefix(), dataStorageClient,
		adminScope.NewSubScope("node_execution_manager"), urlData, eventPublisher, cloudEventPublisher, nodeExecutionEventWriter)
	taskExecutionManager := manager.NewTaskExecutionManager(repo, configuration, dataStorageClient,
		adminScope.NewSubScope("task_execution_manager"), urlData, eventPublisher, cloudEventPublisher)

	logger.Info(ctx, "Initializing a new AdminService")
	return &AdminService{
		TaskManager: manager.NewTaskManager(repo, configuration, workflowengineImpl.NewCompiler(),
			adminScope.NewSubScope("task_manager")),
		WorkflowManager:          workflowManager,
		LaunchPlanManager:        launchPlanManager,
		ExecutionManager:         executionManager,
		NamedEntityManager:       namedEntityManager,
		DescriptionEntityManager: descriptionEntityManager,
		VersionManager:           versionManager,
		NodeExecutionManager:     nodeExecutionManager,
		TaskExecutionManager:     taskExecutionManager,
		ProjectManager:           manager.NewProjectManager(repo, configuration),
		ResourceManager:          resources.NewResourceManager(repo, configuration.ApplicationConfiguration()),
		MetricsManager: manager.NewMetricsManager(workflowManager, executionManager, nodeExecutionManager,
			taskExecutionManager, adminScope.NewSubScope("metrics_manager")),
		Metrics: InitMetrics(adminScope),
	}
}
