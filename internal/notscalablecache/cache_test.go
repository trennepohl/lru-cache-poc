package notscalablecache_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/trennepohl/lru-cache-poc/internal/notscalablecache"
	"github.com/trennepohl/lru-cache-poc/pkg/cache"
	"testing"
)

func TestCacheKeyLimitIsRespected(t *testing.T){
	c, err := notscalablecache.New(cache.Config{MaxKeyCount: 3})
	assert.NoError(t, err)

	c.Write("key1", "key1")
	c.Write("key2", "key2")
	c.Write("key3", "key3")
	c.Write("key4", "key4")
	c.Write("key5", "key5")

	assert.Equal(t, 3,c.Count())
}

func TestLeastLastUsedIsDeletedWhenCacheIsFull(t *testing.T){
	c, err := notscalablecache.New(cache.Config{MaxKeyCount: 3})
	assert.NoError(t, err)

	c.Write("key1", "key1")
	c.Write("key2", "key2")
	c.Write("key3", "key3")
	c.Write("key4", "key4")

	v, err := c.Read("key1")
	assert.Error(t, err)
	assert.Empty(t, v)
}

func TestLeastLastUsedPolicyIsCorrect(t *testing.T){
	c, err := notscalablecache.New(cache.Config{MaxKeyCount: 3})
	assert.NoError(t, err)

	c.Write("key1", "key1")
	c.Write("key2", "key2")
	c.Write("key3", "key3")

	v, err := c.Read("key1")
	assert.NoError(t, err)
	assert.Equal(t, "key1", v)

	c.Write("key4", "key4")
	v, err = c.Read("key1")
	assert.NoError(t, err)
	assert.Equal(t, v, "key1")
}

func TestCount(t *testing.T){
	c, _:= notscalablecache.New(cache.Config{MaxKeyCount: 3})

	count := c.Count()
	assert.Equal(t, 0, count)

	c.Write("key1", "key1")

	count = c.Count()
	assert.Equal(t, 1, count)

}

func TestDelete(t *testing.T){
	c, _:= notscalablecache.New(cache.Config{MaxKeyCount: 3})
	c.Write("key1", "key1")

	err := c.Delete("key1")
	assert.NoError(t, err)

	count := c.Count()
	assert.Equal(t, 0, count)

}

func BenchmarkNotScalableCacheWrites(b *testing.B){
	c, _ := notscalablecache.New(cache.Config{MaxKeyCount: 100})
	for n := 0; n < b.N; n++{
		kv := fmt.Sprintf("key%d", n)
		c.Write(kv,kv)
	}
}