package transformers

import (
	"github.com/nebulaclouds/nebula/nebulastdlib/contextutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils/labeled"
)

func init() {
	labeled.SetMetricKeys(contextutils.ProjectKey, contextutils.DomainKey, contextutils.WorkflowIDKey, contextutils.TaskIDKey)
}
