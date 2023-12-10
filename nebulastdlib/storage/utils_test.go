package storage

import (
	"os"
	"syscall"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	nebulaerrors "github.com/nebulaclouds/nebula/nebulastdlib/errors"
	"github.com/nebulaclouds/stow"
)

func TestIsNotFound(t *testing.T) {
	sysError := &os.PathError{Err: syscall.ENOENT}
	assert.True(t, IsNotFound(sysError))
	nebulaError := errors.Wrap(sysError, "Wrapping \"system not found\" error")
	assert.True(t, IsNotFound(nebulaError))
	secondLevelError := errors.Wrap(nebulaError, "Higher level error")
	assert.True(t, IsNotFound(secondLevelError))

	// more for stow errors
	stowNotFoundError := stow.ErrNotFound
	assert.True(t, IsNotFound(stowNotFoundError))
	nebulaError = errors.Wrap(stowNotFoundError, "Wrapping stow.ErrNotFound")
	assert.True(t, IsNotFound(nebulaError))
	secondLevelError = errors.Wrap(nebulaError, "Higher level error wrapper of the stow.ErrNotFound error")
	assert.True(t, IsNotFound(secondLevelError))
}

func TestIsExceedsLimit(t *testing.T) {
	sysError := &os.PathError{Err: syscall.ENOENT}
	exceedsLimitError := nebulaerrors.Wrapf(ErrExceedsLimit, sysError, "An error wrapped in ErrExceedsLimits")
	failedToWriteCacheError := nebulaerrors.Wrapf(ErrFailedToWriteCache, sysError, "An error wrapped in ErrFailedToWriteCache")

	assert.True(t, IsExceedsLimit(exceedsLimitError))
	assert.False(t, IsExceedsLimit(failedToWriteCacheError))
	assert.False(t, IsExceedsLimit(sysError))
}

func TestIsFailedWriteToCache(t *testing.T) {
	sysError := &os.PathError{Err: syscall.ENOENT}
	exceedsLimitError := nebulaerrors.Wrapf(ErrExceedsLimit, sysError, "An error wrapped in ErrExceedsLimits")
	failedToWriteCacheError := nebulaerrors.Wrapf(ErrFailedToWriteCache, sysError, "An error wrapped in ErrFailedToWriteCache")

	assert.False(t, IsFailedWriteToCache(exceedsLimitError))
	assert.True(t, IsFailedWriteToCache(failedToWriteCacheError))
	assert.False(t, IsFailedWriteToCache(sysError))
}

func TestMapStrings(t *testing.T) {
	t.Run("nothing", func(t *testing.T) {
		assert.Equal(t, []string{}, MapStrings(func(s string) string {
			return s
		}))
	})

	t.Run("one item", func(t *testing.T) {
		assert.Equal(t, []string{"item"}, MapStrings(func(s string) string {
			return s
		}, "item"))
	})

	t.Run("const", func(t *testing.T) {
		assert.Equal(t, []string{"something"}, MapStrings(func(s string) string {
			return "something"
		}, "item"))
	})

	t.Run("half string", func(t *testing.T) {
		assert.Equal(t, []string{"thing", "some"}, MapStrings(func(s string) string {
			return s[len(s)/2:]
		}, "something", "somesome"))
	})
}
