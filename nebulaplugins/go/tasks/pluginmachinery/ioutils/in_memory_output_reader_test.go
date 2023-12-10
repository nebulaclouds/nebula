package ioutils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	nebulaIdlCore "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

func TestInMemoryOutputReader(t *testing.T) {
	deckPath := storage.DataReference("s3://bucket/key")
	lt := map[string]*nebulaIdlCore.Literal{
		"results": {
			Value: &nebulaIdlCore.Literal_Scalar{
				Scalar: &nebulaIdlCore.Scalar{
					Value: &nebulaIdlCore.Scalar_Primitive{
						Primitive: &nebulaIdlCore.Primitive{Value: &nebulaIdlCore.Primitive_Integer{Integer: 3}},
					},
				},
			},
		},
	}
	or := NewInMemoryOutputReader(&nebulaIdlCore.LiteralMap{Literals: lt}, &deckPath, nil)

	assert.Equal(t, &deckPath, or.DeckPath)
	ctx := context.TODO()

	ok, err := or.IsError(ctx)
	assert.False(t, ok)
	assert.NoError(t, err)

	assert.False(t, or.IsFile(ctx))

	ok, err = or.Exists(ctx)
	assert.True(t, ok)
	assert.NoError(t, err)

	literalMap, executionErr, err := or.Read(ctx)
	assert.Equal(t, lt, literalMap.Literals)
	assert.Nil(t, executionErr)
	assert.NoError(t, err)
}
