package factory

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/service"
	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/catalog"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/config"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/executors"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/array"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/branch"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/dynamic"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/end"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/gate"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/interfaces"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/recovery"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/start"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/subworkflow"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/subworkflow/launchplan"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/task"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type handlerFactory struct {
	handlers map[v1alpha1.NodeKind]interfaces.NodeHandler

	workflowLauncher launchplan.Executor
	launchPlanReader launchplan.Reader
	kubeClient       executors.Client
	kubeClientset    kubernetes.Interface
	catalogClient    catalog.Client
	recoveryClient   recovery.Client
	eventConfig      *config.EventConfig
	clusterID        string
	signalClient     service.SignalServiceClient
	scope            promutils.Scope
}

func (f *handlerFactory) GetHandler(kind v1alpha1.NodeKind) (interfaces.NodeHandler, error) {
	h, ok := f.handlers[kind]
	if !ok {
		return nil, errors.Errorf("Handler not registered for NodeKind [%v]", kind)
	}
	return h, nil
}

func (f *handlerFactory) Setup(ctx context.Context, executor interfaces.Node, setup interfaces.SetupContext) error {
	t, err := task.New(ctx, f.kubeClient, f.kubeClientset, f.catalogClient, f.eventConfig, f.clusterID, f.scope)
	if err != nil {
		return err
	}

	arrayHandler, err := array.New(executor, f.eventConfig, f.scope)
	if err != nil {
		return err
	}

	f.handlers = map[v1alpha1.NodeKind]interfaces.NodeHandler{
		v1alpha1.NodeKindBranch:   branch.New(executor, f.eventConfig, f.scope),
		v1alpha1.NodeKindTask:     dynamic.New(t, executor, f.launchPlanReader, f.eventConfig, f.scope),
		v1alpha1.NodeKindWorkflow: subworkflow.New(executor, f.workflowLauncher, f.recoveryClient, f.eventConfig, f.scope),
		v1alpha1.NodeKindGate:     gate.New(f.eventConfig, f.signalClient, f.scope),
		v1alpha1.NodeKindArray:    arrayHandler,
		v1alpha1.NodeKindStart:    start.New(),
		v1alpha1.NodeKindEnd:      end.New(),
	}

	for _, v := range f.handlers {
		if err := v.Setup(ctx, setup); err != nil {
			return err
		}
	}
	return nil
}

func NewHandlerFactory(ctx context.Context, workflowLauncher launchplan.Executor, launchPlanReader launchplan.Reader,
	kubeClient executors.Client, kubeClientset kubernetes.Interface, catalogClient catalog.Client, recoveryClient recovery.Client, eventConfig *config.EventConfig,
	clusterID string, signalClient service.SignalServiceClient, scope promutils.Scope) (interfaces.HandlerFactory, error) {

	return &handlerFactory{
		workflowLauncher: workflowLauncher,
		launchPlanReader: launchPlanReader,
		kubeClient:       kubeClient,
		kubeClientset:    kubeClientset,
		catalogClient:    catalogClient,
		recoveryClient:   recoveryClient,
		eventConfig:      eventConfig,
		clusterID:        clusterID,
		signalClient:     signalClient,
		scope:            scope,
	}, nil
}
