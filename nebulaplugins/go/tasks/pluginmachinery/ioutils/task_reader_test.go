package ioutils

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core/mocks"
	"github.com/nebulaclouds/nebula/nebulastdlib/contextutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils/labeled"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

const dummyPath = storage.DataReference("test")

func TestLazyUploadingTaskReader_Happy(t *testing.T) {
	ttm := &core.TaskTemplate{}

	ctx := context.TODO()
	tr := &mocks.TaskReader{}
	tr.OnRead(ctx).Return(ttm, nil)

	ds, err := storage.NewDataStore(&storage.Config{
		Type: storage.TypeMemory,
	}, promutils.NewTestScope())
	assert.NoError(t, err)

	ltr := NewLazyUploadingTaskReader(tr, dummyPath, ds)

	x, err := ltr.Read(ctx)
	assert.NoError(t, err)
	assert.Equal(t, x, ttm)

	p, err := ltr.Path(ctx)
	assert.NoError(t, err)
	assert.Equal(t, p, dummyPath)

	v, err := ds.Head(ctx, dummyPath)
	assert.NoError(t, err)
	assert.True(t, v.Exists())
}

// test storage.ProtobufStore to test upload failure
type failingProtoStore struct {
	storage.ProtobufStore
}

func (d *failingProtoStore) WriteProtobuf(ctx context.Context, reference storage.DataReference, opts storage.Options, msg proto.Message) error {
	return fmt.Errorf("failed")
}

func TestLazyUploadingTaskReader_TaskWriteFailure(t *testing.T) {
	ttm := &core.TaskTemplate{}

	ctx := context.TODO()
	tr := &mocks.TaskReader{}
	tr.OnRead(ctx).Return(ttm, nil)

	ltr := NewLazyUploadingTaskReader(tr, dummyPath, &failingProtoStore{})

	x, err := ltr.Read(ctx)
	assert.NoError(t, err)
	assert.Equal(t, x, ttm)

	p, err := ltr.Path(ctx)
	assert.Error(t, err)
	assert.Equal(t, p, storage.DataReference(""))
}

func TestLazyUploadingTaskReader_TaskReadFailure(t *testing.T) {

	ctx := context.TODO()
	tr := &mocks.TaskReader{}
	tr.OnRead(ctx).Return(nil, fmt.Errorf("read fail"))

	ds, err := storage.NewDataStore(&storage.Config{
		Type: storage.TypeMemory,
	}, promutils.NewTestScope())
	assert.NoError(t, err)

	ltr := NewLazyUploadingTaskReader(tr, dummyPath, ds)

	x, err := ltr.Read(ctx)
	assert.Error(t, err)
	assert.Nil(t, x)

	p, err := ltr.Path(ctx)
	assert.Error(t, err)
	assert.Equal(t, p, storage.DataReference(""))

	v, err := ds.Head(ctx, dummyPath)
	assert.NoError(t, err)
	assert.False(t, v.Exists())
}

func init() {
	labeled.SetMetricKeys(contextutils.ExecIDKey)
}
