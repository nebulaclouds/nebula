package mocks

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/data/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

// Mock implementation of a RemoteURLInterface
type MockRemoteURL struct {
	GetCallback func(ctx context.Context, uri string) (admin.UrlBlob, error)
}

func (m *MockRemoteURL) Get(ctx context.Context, uri string) (admin.UrlBlob, error) {
	if m.GetCallback != nil {
		return m.GetCallback(ctx, uri)
	}
	return admin.UrlBlob{}, nil
}

func NewMockRemoteURL() interfaces.RemoteURLInterface {
	return &MockRemoteURL{}
}
