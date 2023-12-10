package gormimpl

import (
	"context"

	"gorm.io/gorm"

	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/models"
	idl_datacatalog "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type tagRepo struct {
	db               *gorm.DB
	errorTransformer errors.ErrorTransformer
	repoMetrics      gormMetrics
}

func NewTagRepo(db *gorm.DB, errorTransformer errors.ErrorTransformer, scope promutils.Scope) interfaces.TagRepo {
	return &tagRepo{
		db:               db,
		errorTransformer: errorTransformer,
		repoMetrics:      newGormMetrics(scope),
	}
}

func (h *tagRepo) Create(ctx context.Context, tag models.Tag) error {
	timer := h.repoMetrics.CreateDuration.Start(ctx)
	defer timer.Stop()

	db := h.db.WithContext(ctx).Create(&tag)

	if db.Error != nil {
		return h.errorTransformer.ToDataCatalogError(db.Error)
	}
	return nil
}

func (h *tagRepo) Get(ctx context.Context, in models.TagKey) (models.Tag, error) {
	timer := h.repoMetrics.GetDuration.Start(ctx)
	defer timer.Stop()

	var tag models.Tag
	result := h.db.WithContext(ctx).Preload("Artifact").
		Preload("Artifact.ArtifactData").
		Preload("Artifact.Partitions", func(db *gorm.DB) *gorm.DB {
			return db.WithContext(ctx).Order("partitions.created_at ASC") // preserve the order in which the partitions were created
		}).
		Preload("Artifact.Tags").
		Order("tags.created_at DESC").
		First(&tag, &models.Tag{
			TagKey: in,
		})

	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return models.Tag{}, errors.GetMissingEntityError("Tag", &idl_datacatalog.Tag{
				Name: tag.TagName,
			})
		}
		return models.Tag{}, h.errorTransformer.ToDataCatalogError(result.Error)
	}

	return tag, nil
}
