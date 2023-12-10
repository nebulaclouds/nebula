package task

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/utils"
	"github.com/nebulaclouds/nebula/nebulastdlib/errors"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

// TODO this file exists only until we need to support dynamic nodes instead of closure.
// Once closure migration is done, this file should be deleted.
const implicitFutureFileName = "futures.pb"
const implicitCompileWorkflowsName = "futures_compiled.pb"
const implicitCompiledWorkflowClosureName = "dynamic_compiled.pb"

type FutureFileReader struct {
	RemoteFileWorkflowStore
	loc                    storage.DataReference
	nebulaWfCRDCacheLoc     storage.DataReference
	nebulaWfClosureCacheLoc storage.DataReference
	store                  *storage.DataStore
}

func (f FutureFileReader) GetLoc() storage.DataReference {
	return f.loc
}

func (f FutureFileReader) Exists(ctx context.Context) (bool, error) {
	metadata, err := f.store.Head(ctx, f.loc)
	// If no futures file produced, then declare success and return.
	if err != nil {
		logger.Warnf(ctx, "Failed to read futures file. Error: %v", err)
		return false, errors.Wrapf(utils.ErrorCodeUser, err, "Failed to do HEAD on futures file.")
	}
	return metadata.Exists(), nil
}

func (f FutureFileReader) Read(ctx context.Context) (*core.DynamicJobSpec, error) {
	djSpec := &core.DynamicJobSpec{}
	if err := f.store.ReadProtobuf(ctx, f.loc, djSpec); err != nil {
		logger.Warnf(ctx, "Failed to read futures file. Error: %v", err)
		return nil, errors.Wrapf(utils.ErrorCodeSystem, err, "Failed to read futures protobuf file.")
	}

	return djSpec, nil
}

func (f FutureFileReader) CacheExists(ctx context.Context) (bool, error) {
	exists, err := f.RemoteFileWorkflowStore.Exists(ctx, f.nebulaWfCRDCacheLoc)
	if err != nil || !exists {
		return exists, err
	}
	return f.RemoteFileWorkflowStore.Exists(ctx, f.nebulaWfClosureCacheLoc)
}

func (f FutureFileReader) Cache(ctx context.Context, wf *v1alpha1.NebulaWorkflow, workflowClosure *core.CompiledWorkflowClosure) error {
	err := f.RemoteFileWorkflowStore.PutNebulaWorkflowCRD(ctx, wf, f.nebulaWfCRDCacheLoc)
	if err != nil {
		return err
	}
	return f.RemoteFileWorkflowStore.PutCompiledNebulaWorkflow(ctx, workflowClosure, f.nebulaWfClosureCacheLoc)
}

type CacheContents struct {
	WorkflowCRD      *v1alpha1.NebulaWorkflow
	CompiledWorkflow *core.CompiledWorkflowClosure
}

func (f FutureFileReader) RetrieveCache(ctx context.Context) (CacheContents, error) {
	workflowCRD, err := f.RemoteFileWorkflowStore.GetWorkflowCRD(ctx, f.nebulaWfCRDCacheLoc)
	if err != nil {
		return CacheContents{}, err
	}
	compiledWorkflow, err := f.RemoteFileWorkflowStore.GetCompiledWorkflow(ctx, f.nebulaWfClosureCacheLoc)
	if err != nil {
		return CacheContents{}, err
	}
	return CacheContents{
		WorkflowCRD:      workflowCRD,
		CompiledWorkflow: compiledWorkflow,
	}, nil
}

func NewRemoteFutureFileReader(ctx context.Context, dataDir storage.DataReference, store *storage.DataStore) (FutureFileReader, error) {
	loc, err := store.ConstructReference(ctx, dataDir, implicitFutureFileName)
	if err != nil {
		logger.Warnf(ctx, "Failed to construct data path for futures file. Error: %v", err)
		return FutureFileReader{}, errors.Wrapf(utils.ErrorCodeSystem, err, "failed to construct data path")
	}

	nebulaWfCRDCacheLoc, err := store.ConstructReference(ctx, dataDir, implicitCompileWorkflowsName)
	if err != nil {
		logger.Warnf(ctx, "Failed to construct data path for compile workflows file, error: %s", err)
		return FutureFileReader{}, errors.Wrapf(utils.ErrorCodeSystem, err, "failed to construct reference for workflow CRD cache location")
	}
	nebulaWfClosureCacheLoc, err := store.ConstructReference(ctx, dataDir, implicitCompiledWorkflowClosureName)
	if err != nil {
		logger.Warnf(ctx, "Failed to construct data path for compile workflows file, error: %s", err)
		return FutureFileReader{}, errors.Wrapf(utils.ErrorCodeSystem, err, "failed to construct reference for compiled workflow closure cache location")
	}

	return FutureFileReader{
		loc:                     loc,
		nebulaWfCRDCacheLoc:      nebulaWfCRDCacheLoc,
		nebulaWfClosureCacheLoc:  nebulaWfClosureCacheLoc,
		store:                   store,
		RemoteFileWorkflowStore: NewRemoteWorkflowStore(store),
	}, nil
}
