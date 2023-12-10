package validators

import (
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
)

const (
	tagName   = "tagName"
	tagEntity = "tag"
)

func ValidateTag(tag *datacatalog.Tag) error {
	if tag == nil {
		return NewMissingArgumentError(tagEntity)
	}
	if err := ValidateDatasetID(tag.Dataset); err != nil {
		return err
	}

	if err := ValidateEmptyStringField(tag.Name, tagName); err != nil {
		return err
	}

	if err := ValidateEmptyStringField(tag.ArtifactId, artifactID); err != nil {
		return err
	}
	return nil
}
