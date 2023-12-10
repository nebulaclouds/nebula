package catalog

import (
	"context"
	"time"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/catalog"
	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/io"
)

var (
	_ catalog.Client = &NOOPCatalog{}
)

var disabledStatus = catalog.NewStatus(core.CatalogCacheStatus_CACHE_DISABLED, nil)

type NOOPCatalog struct {
}

func (n NOOPCatalog) Get(_ context.Context, _ catalog.Key) (catalog.Entry, error) {
	return catalog.NewCatalogEntry(nil, disabledStatus), nil
}

func (n NOOPCatalog) Put(_ context.Context, _ catalog.Key, _ io.OutputReader, _ catalog.Metadata) (catalog.Status, error) {
	return disabledStatus, nil
}

func (n NOOPCatalog) Update(_ context.Context, _ catalog.Key, _ io.OutputReader, _ catalog.Metadata) (catalog.Status, error) {
	return disabledStatus, nil
}

func (n NOOPCatalog) GetOrExtendReservation(_ context.Context, _ catalog.Key, _ string, _ time.Duration) (*datacatalog.Reservation, error) {
	return nil, nil
}

func (n NOOPCatalog) ReleaseReservation(_ context.Context, _ catalog.Key, _ string) error {
	return nil
}
