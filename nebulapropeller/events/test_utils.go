package events

import (
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/config"
)

var inlineEventConfig = &config.EventConfig{
	RawOutputPolicy: config.RawOutputPolicyInline,
}

var inlineEventConfigFallback = &config.EventConfig{
	RawOutputPolicy:           config.RawOutputPolicyInline,
	FallbackToOutputReference: true,
}

var referenceEventConfig = &config.EventConfig{
	RawOutputPolicy: config.RawOutputPolicyReference,
}

var referenceURI = "s3://foo/bar/outputs.pb"
var deckURI = "s3://foo/bar/deck.html"

var outputData = &core.LiteralMap{
	Literals: map[string]*core.Literal{
		"foo": {
			Value: &core.Literal_Scalar{
				Scalar: &core.Scalar{
					Value: &core.Scalar_Primitive{
						Primitive: &core.Primitive{
							Value: &core.Primitive_Integer{
								Integer: 4,
							},
						},
					},
				},
			},
		},
	},
}

var workflowExecID = &core.WorkflowExecutionIdentifier{
	Project: "p",
	Domain:  "d",
	Name:    "n",
}

var nodeExecID = &core.NodeExecutionIdentifier{
	ExecutionId: workflowExecID,
	NodeId:      "node_id",
}
