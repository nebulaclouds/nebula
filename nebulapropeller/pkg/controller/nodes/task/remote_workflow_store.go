package task

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

type RemoteFileWorkflowStore struct {
	store *storage.DataStore
}

func (r RemoteFileWorkflowStore) Exists(ctx context.Context, path storage.DataReference) (bool, error) {
	metadata, err := r.store.Head(ctx, path)
	// If no futures file produced, then declare success and return.
	if err != nil {
		logger.Warnf(ctx, "Failed to read futures file. Error: %v", err)
		return false, errors.Wrap(err, "Failed to do HEAD on futures file.")
	}
	return metadata.Exists(), nil
}

func (r RemoteFileWorkflowStore) PutNebulaWorkflowCRD(ctx context.Context, wf *v1alpha1.NebulaWorkflow, target storage.DataReference) error {
	raw, err := json.Marshal(wf)
	if err != nil {
		return err
	}

	return r.store.WriteRaw(ctx, target, int64(len(raw)), storage.Options{}, bytes.NewReader(raw))
}

func (r RemoteFileWorkflowStore) PutCompiledNebulaWorkflow(ctx context.Context, workflow *core.CompiledWorkflowClosure, target storage.DataReference) error {
	return r.store.WriteProtobuf(ctx, target, storage.Options{}, workflow)
}

func (r RemoteFileWorkflowStore) getRawBytes(ctx context.Context, source storage.DataReference) ([]byte, error) {

	rawReader, err := r.store.ReadRaw(ctx, source)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(rawReader)
	if err != nil {
		return nil, err
	}

	err = rawReader.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (r RemoteFileWorkflowStore) GetWorkflowCRD(ctx context.Context, source storage.DataReference) (*v1alpha1.NebulaWorkflow, error) {
	wfBytes, err := r.getRawBytes(ctx, source)
	if err != nil {
		return nil, err
	}

	wf := &v1alpha1.NebulaWorkflow{}
	return wf, json.Unmarshal(wfBytes, wf)
}

func (r RemoteFileWorkflowStore) GetCompiledWorkflow(ctx context.Context, source storage.DataReference) (*core.CompiledWorkflowClosure, error) {
	var closure core.CompiledWorkflowClosure
	err := r.store.ReadProtobuf(ctx, source, &closure)
	return &closure, err
}

func NewRemoteWorkflowStore(store *storage.DataStore) RemoteFileWorkflowStore {
	return RemoteFileWorkflowStore{store: store}
}
