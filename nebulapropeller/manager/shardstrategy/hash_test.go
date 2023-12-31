package shardstrategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeKeyRange(t *testing.T) {
	keyspaceSize := 32
	for podCount := 1; podCount < keyspaceSize; podCount++ {
		keysCovered := 0
		minKeyRangeSize := keyspaceSize / podCount
		for podIndex := 0; podIndex < podCount; podIndex++ {
			startIndex, endIndex := ComputeKeyRange(keyspaceSize, podCount, podIndex)

			rangeSize := endIndex - startIndex
			keysCovered += rangeSize
			assert.True(t, rangeSize-minKeyRangeSize >= 0)
			assert.True(t, rangeSize-minKeyRangeSize <= 1)
		}

		assert.Equal(t, keyspaceSize, keysCovered)
	}
}

func TestHashHashCode(t *testing.T) {
	expectedStrategy := &HashShardStrategy{
		ShardCount: 3,
	}
	actualStrategy := &HashShardStrategy{
		ShardCount: 3,
	}
	expectedHashcode, err := expectedStrategy.HashCode()
	assert.NoError(t, err)
	actualHashcode, err := actualStrategy.HashCode()
	assert.NoError(t, err)
	assert.Equal(t, expectedHashcode, actualHashcode)
}
