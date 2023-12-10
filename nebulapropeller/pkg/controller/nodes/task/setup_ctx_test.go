package task

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core/mocks"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/interfaces"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type dummySetupCtx struct {
	interfaces.SetupContext
	testScopeName string
}

func (d dummySetupCtx) MetricsScope() promutils.Scope {
	return promutils.NewScope(d.testScopeName)
}

func Test_nameSpacedSetupCtx_MetricsScope(t *testing.T) {
	r := &mocks.ResourceRegistrar{}
	ns := newNameSpacedSetupCtx(&setupContext{SetupContext: &dummySetupCtx{testScopeName: "test-scope-1"}}, r, "p1")
	scope := ns.MetricsScope()
	assert.NotNil(t, scope)
	assert.Equal(t, "test_scope_1:p1:", scope.CurrentScope())
}
