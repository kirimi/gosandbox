package test

import (
	"example/hello/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cache(t *testing.T) {
	t.Parallel()

	testCache := cache.NewMyCache()

	t.Run("correctly stored value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "some value"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)

		assert.Equal(t, value, storedValue)
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()

		parallelFactor := 100_000
		emulateLoad(t, testCache, parallelFactor)

	})
}
