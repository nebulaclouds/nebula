package array

import (
	"bytes"
	"context"
	"fmt"

	idlcore "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/interfaces"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/task"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/task/codex"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/task/k8s"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

func appendLiteral(name string, literal *idlcore.Literal, outputLiterals map[string]*idlcore.Literal, length int) {
	outputLiteral, exists := outputLiterals[name]
	if !exists {
		outputLiteral = &idlcore.Literal{
			Value: &idlcore.Literal_Collection{
				Collection: &idlcore.LiteralCollection{
					Literals: make([]*idlcore.Literal, 0, length),
				},
			},
		}

		outputLiterals[name] = outputLiteral
	}

	collection := outputLiteral.GetCollection()
	collection.Literals = append(collection.Literals, literal)
}

func buildSubNodeID(nCtx interfaces.NodeExecutionContext, index int) string {
	return fmt.Sprintf("%s-n%d", nCtx.NodeID(), index)
}

func bytesFromK8sPluginState(pluginState k8s.PluginState) ([]byte, error) {
	buffer := make([]byte, 0, task.MaxPluginStateSizeBytes)
	bufferWriter := bytes.NewBuffer(buffer)

	codec := codex.GobStateCodec{}
	if err := codec.Encode(pluginState, bufferWriter); err != nil {
		return nil, err
	}

	return bufferWriter.Bytes(), nil
}

func constructOutputReferences(ctx context.Context, nCtx interfaces.NodeExecutionContext, postfix ...string) (storage.DataReference, storage.DataReference, error) {
	subDataDir, err := nCtx.DataStore().ConstructReference(ctx, nCtx.NodeStatus().GetDataDir(), postfix...)
	if err != nil {
		return "", "", err
	}

	subOutputDir, err := nCtx.DataStore().ConstructReference(ctx, nCtx.NodeStatus().GetOutputDir(), postfix...)
	if err != nil {
		return "", "", err
	}

	return subDataDir, subOutputDir, nil
}

func isTerminalNodePhase(nodePhase v1alpha1.NodePhase) bool {
	return nodePhase == v1alpha1.NodePhaseSucceeded || nodePhase == v1alpha1.NodePhaseFailed || nodePhase == v1alpha1.NodePhaseTimedOut ||
		nodePhase == v1alpha1.NodePhaseSkipped || nodePhase == v1alpha1.NodePhaseRecovered
}
