package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

//go:generate mockery -name=SignalInterface -output=../mocks -case=underscore

// Interface for managing Nebula Signals
type SignalInterface interface {
	GetOrCreateSignal(ctx context.Context, request admin.SignalGetOrCreateRequest) (*admin.Signal, error)
	ListSignals(ctx context.Context, request admin.SignalListRequest) (*admin.SignalList, error)
	SetSignal(ctx context.Context, request admin.SignalSetRequest) (*admin.SignalSetResponse, error)
}
