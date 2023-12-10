package impl

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/datacatalog/pkg/errors"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

const artifactDataFile = "data.pb"

// ArtifactDataStore stores and retrieves ArtifactData values in a data.pb
type ArtifactDataStore interface {
	PutData(ctx context.Context, artifact *datacatalog.Artifact, data *datacatalog.ArtifactData) (storage.DataReference, error)
	GetData(ctx context.Context, dataModel models.ArtifactData) (*core.Literal, error)
	DeleteData(ctx context.Context, dataModel models.ArtifactData) error
}

type artifactDataStore struct {
	store         *storage.DataStore
	storagePrefix storage.DataReference
}

func (m *artifactDataStore) getDataLocation(ctx context.Context, artifact *datacatalog.Artifact, data *datacatalog.ArtifactData) (storage.DataReference, error) {
	dataset := artifact.Dataset
	return m.store.ConstructReference(ctx, m.storagePrefix, dataset.Project, dataset.Domain, dataset.Name, dataset.Version, artifact.Id, data.Name, artifactDataFile)
}

// Store marshalled data in data.pb under the storage prefix
func (m *artifactDataStore) PutData(ctx context.Context, artifact *datacatalog.Artifact, data *datacatalog.ArtifactData) (storage.DataReference, error) {
	dataLocation, err := m.getDataLocation(ctx, artifact, data)
	if err != nil {
		return "", errors.NewDataCatalogErrorf(codes.Internal, "Unable to generate data location %s, err %v", dataLocation.String(), err)
	}
	err = m.store.WriteProtobuf(ctx, dataLocation, storage.Options{}, data.Value)
	if err != nil {
		return "", errors.NewDataCatalogErrorf(codes.Internal, "Unable to store artifact data in location %s, err %v", dataLocation.String(), err)
	}

	return dataLocation, nil
}

// Retrieve the literal value of the ArtifactData from its specified location
func (m *artifactDataStore) GetData(ctx context.Context, dataModel models.ArtifactData) (*core.Literal, error) {
	var value core.Literal
	err := m.store.ReadProtobuf(ctx, storage.DataReference(dataModel.Location), &value)
	if err != nil {
		return nil, errors.NewDataCatalogErrorf(codes.Internal, "Unable to read artifact data from location %s, err %v", dataModel.Location, err)
	}

	return &value, nil
}

// DeleteData removes the stored artifact data from the underlying blob storage
func (m *artifactDataStore) DeleteData(ctx context.Context, dataModel models.ArtifactData) error {
	if err := m.store.Delete(ctx, storage.DataReference(dataModel.Location)); err != nil {
		return errors.NewDataCatalogErrorf(codes.Internal, "Unable to delete artifact data in location %s, err %v", dataModel.Location, err)
	}

	return nil
}

func NewArtifactDataStore(store *storage.DataStore, storagePrefix storage.DataReference) ArtifactDataStore {
	return &artifactDataStore{
		store:         store,
		storagePrefix: storagePrefix,
	}
}
