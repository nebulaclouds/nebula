package schedule

import (
	"context"
	"time"

	gizmoConfig "github.com/NYTimes/gizmo/pubsub/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async"
	awsSchedule "github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/schedule/aws"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/schedule/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/schedule/noop"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/common"
	managerInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/interfaces"
	repoInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	runtimeInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	nebulascheduler "github.com/nebulaclouds/nebula/nebulaadmin/scheduler/dbapi"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type WorkflowSchedulerConfig struct {
	Retries         int
	SchedulerConfig runtimeInterfaces.SchedulerConfig
	Scope           promutils.Scope
}

type WorkflowScheduler interface {
	GetEventScheduler() interfaces.EventScheduler
	GetWorkflowExecutor(executionManager managerInterfaces.ExecutionInterface,
		launchPlanManager managerInterfaces.LaunchPlanInterface) interfaces.WorkflowExecutor
}

type workflowScheduler struct {
	cfg              WorkflowSchedulerConfig
	eventScheduler   interfaces.EventScheduler
	workflowExecutor interfaces.WorkflowExecutor
}

func (w *workflowScheduler) GetEventScheduler() interfaces.EventScheduler {
	return w.eventScheduler
}

func (w *workflowScheduler) GetWorkflowExecutor(
	executionManager managerInterfaces.ExecutionInterface,
	launchPlanManager managerInterfaces.LaunchPlanInterface) interfaces.WorkflowExecutor {
	if w.workflowExecutor == nil {
		sqsConfig := gizmoConfig.SQSConfig{
			QueueName:           w.cfg.SchedulerConfig.WorkflowExecutorConfig.ScheduleQueueName,
			QueueOwnerAccountID: w.cfg.SchedulerConfig.WorkflowExecutorConfig.AccountID,
		}
		sqsConfig.Region = w.cfg.SchedulerConfig.WorkflowExecutorConfig.Region
		w.workflowExecutor = awsSchedule.NewWorkflowExecutor(
			sqsConfig, w.cfg.SchedulerConfig, executionManager, launchPlanManager, w.cfg.Scope.NewSubScope("workflow_executor"))
	}
	return w.workflowExecutor
}

func NewWorkflowScheduler(db repoInterfaces.Repository, cfg WorkflowSchedulerConfig) WorkflowScheduler {
	var eventScheduler interfaces.EventScheduler
	var workflowExecutor interfaces.WorkflowExecutor

	switch cfg.SchedulerConfig.EventSchedulerConfig.Scheme {
	case common.AWS:
		awsConfig := aws.NewConfig().WithRegion(cfg.SchedulerConfig.WorkflowExecutorConfig.Region).WithMaxRetries(cfg.Retries)
		var sess *session.Session
		var err error
		err = async.Retry(cfg.SchedulerConfig.ReconnectAttempts,
			time.Duration(cfg.SchedulerConfig.ReconnectDelaySeconds)*time.Second, func() error {
				sess, err = session.NewSession(awsConfig)
				if err != nil {
					logger.Warnf(context.TODO(), "Failed to initialize new event scheduler with aws config: [%+v] and err: %v", awsConfig, err)
				}
				return err
			})

		if err != nil {
			panic(err)
		}
		eventScheduler = awsSchedule.NewCloudWatchScheduler(
			cfg.SchedulerConfig.EventSchedulerConfig.ScheduleRole, cfg.SchedulerConfig.EventSchedulerConfig.TargetName, sess, awsConfig,
			cfg.Scope.NewSubScope("cloudwatch_scheduler"))
	case common.Local:
		logger.Infof(context.Background(),
			"Using default nebula scheduler implementation")
		eventScheduler = nebulascheduler.New(db)
	default:
		logger.Infof(context.Background(),
			"Using default noop event scheduler implementation for cloud provider type [%s]",
			cfg.SchedulerConfig.EventSchedulerConfig.Scheme)
		eventScheduler = noop.NewNoopEventScheduler()
	}

	switch cfg.SchedulerConfig.WorkflowExecutorConfig.Scheme {
	case common.AWS:
		// Do nothing, this special case depends on the execution manager and launch plan manager having been
		// initialized and is handled in GetWorkflowExecutor.
		break
	case common.Local:
		fallthrough
	default:
		logger.Infof(context.Background(),
			"Using default noop workflow executor implementation for cloud provider type [%s]",
			cfg.SchedulerConfig.EventSchedulerConfig.Scheme)
		workflowExecutor = noop.NewWorkflowExecutor()
	}
	return &workflowScheduler{
		cfg:              cfg,
		eventScheduler:   eventScheduler,
		workflowExecutor: workflowExecutor,
	}
}
