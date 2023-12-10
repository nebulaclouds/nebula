package shardstrategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentHashCode(t *testing.T) {
	expectedStrategy := &EnvironmentShardStrategy{
		EnvType: Project,
		PerShardIDs: [][]string{
			{"nebulasnacks"},
			{"nebulafoo", "nebulabar"},
			{"*"},
		},
	}
	actualStrategy := &EnvironmentShardStrategy{
		EnvType: Project,
		PerShardIDs: [][]string{
			{"nebulasnacks"},
			{"nebulafoo", "nebulabar"},
			{"*"},
		},
	}
	expectedHashcode, err := expectedStrategy.HashCode()
	assert.NoError(t, err)
	actualHashcode, err := actualStrategy.HashCode()
	assert.NoError(t, err)
	assert.Equal(t, expectedHashcode, actualHashcode)
}
