package transformers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/datacatalog/pkg/common"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
)

func TestPaginationDefaults(t *testing.T) {
	listModelsInput := &models.ListModelsInput{}
	err := ApplyPagination(nil, listModelsInput)
	assert.NoError(t, err)
	assert.Equal(t, common.DefaultPageOffset, listModelsInput.Offset)
	assert.Equal(t, common.MaxPageLimit, listModelsInput.Limit)
	assert.Equal(t, "artifacts.created_at desc", listModelsInput.SortParameter.GetDBOrderExpression("artifacts"))
}

func TestPaginationInvalidToken(t *testing.T) {
	listModelsInput := &models.ListModelsInput{}
	err := ApplyPagination(&datacatalog.PaginationOptions{Token: "pg. 1"}, listModelsInput)
	assert.Error(t, err)
}

func TestCorrectPagination(t *testing.T) {
	listModelsInput := &models.ListModelsInput{}
	err := ApplyPagination(&datacatalog.PaginationOptions{
		Token:     "100",
		Limit:     50,
		SortKey:   datacatalog.PaginationOptions_CREATION_TIME,
		SortOrder: datacatalog.PaginationOptions_DESCENDING,
	}, listModelsInput)
	assert.NoError(t, err)
	assert.Equal(t, 50, listModelsInput.Limit)
	assert.Equal(t, 100, listModelsInput.Offset)
	assert.Equal(t, "artifacts.created_at desc", listModelsInput.SortParameter.GetDBOrderExpression("artifacts"))
}
