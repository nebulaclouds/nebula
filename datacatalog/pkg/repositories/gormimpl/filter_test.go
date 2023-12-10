package gormimpl

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/datacatalog/pkg/common"
)

func TestGormValueFilter(t *testing.T) {
	filter := NewGormValueFilter(common.Equal, "key", "region")
	expression, err := filter.GetDBQueryExpression("partitions")
	assert.NoError(t, err)
	assert.Equal(t, expression.Query, "partitions.key = ?")
	assert.Equal(t, expression.Args, "region")
}

func TestGormValueFilterInvalidOperator(t *testing.T) {
	filter := NewGormValueFilter(123, "key", "region")
	_, err := filter.GetDBQueryExpression("partitions")
	assert.Error(t, err)
}
