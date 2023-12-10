package task

import (
	"k8s.io/apimachinery/pkg/types"

	pluginCore "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/interfaces"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type setupContext struct {
	interfaces.SetupContext
	kubeClient    pluginCore.KubeClient
	secretManager pluginCore.SecretManager
}

func (s setupContext) SecretManager() pluginCore.SecretManager {
	return s.secretManager
}

func (s setupContext) KubeClient() pluginCore.KubeClient {
	return s.kubeClient
}

func (s setupContext) EnqueueOwner() pluginCore.EnqueueOwner {
	return func(ownerId types.NamespacedName) error {
		s.SetupContext.EnqueueOwner()(ownerId.String())
		return nil
	}
}

func (t *Handler) newSetupContext(sCtx interfaces.SetupContext) *setupContext {

	return &setupContext{
		SetupContext:  sCtx,
		kubeClient:    t.kubeClient,
		secretManager: t.secretManager,
	}
}

type nameSpacedSetupCtx struct {
	*setupContext
	rn       pluginCore.ResourceRegistrar
	pluginID string
}

func (n nameSpacedSetupCtx) ResourceRegistrar() pluginCore.ResourceRegistrar {
	return n.rn
}

func (n nameSpacedSetupCtx) MetricsScope() promutils.Scope {
	return n.SetupContext.MetricsScope().NewSubScope(n.pluginID)
}

func newNameSpacedSetupCtx(sCtx *setupContext, rn pluginCore.ResourceRegistrar, pluginID string) nameSpacedSetupCtx {
	return nameSpacedSetupCtx{
		setupContext: sCtx,
		rn:           rn,
		pluginID:     pluginID,
	}
}
