package cache

import (
	"testing"

	"github.com/ryanmab/rdap-go/internal/query"
	"github.com/stretchr/testify/assert"
)

func TestCache_SetGet(t *testing.T) {
	t.Run("Set and Get an entry", func(t *testing.T) {
		cache := New()

		retrievedValue := cache.Get(query.DomainQuery, "abc.com")
		assert.Nil(t, retrievedValue)

		cache.Set(query.DomainQuery, "abc.com", "cached data")

		retrievedValue = cache.Get(query.DomainQuery, "abc.com")
		assert.Equal(t, "cached data", *retrievedValue)
	})

	t.Run("Clearing cache", func(t *testing.T) {
		cache := New()

		cache.Set(query.DomainQuery, "abc.com", "cached data")

		retrievedValue := cache.Get(query.DomainQuery, "abc.com")

		assert.Equal(t, "cached data", *retrievedValue)
		cache.Clear()

		retrievedValue = cache.Get(query.DomainQuery, "abc.com")
		assert.Nil(t, retrievedValue)
	})

	t.Run("Keys in different query namespaces do not conflict", func(t *testing.T) {
		cache := New()

		cache.Set(query.DomainQuery, "some-key", "cached-value-1")
		cache.Set(query.IPv4Query, "some-key", "cached-value-2")

		retrievedvalue := cache.Get(query.DomainQuery, "some-key")
		assert.Equal(t, "cached-value-1", *retrievedvalue)

		retrievedvalue = cache.Get(query.IPv4Query, "some-key")
		assert.Equal(t, "cached-value-2", *retrievedvalue)

		retrievedvalue = cache.Get(query.IPv6Query, "some-key")
		assert.Nil(t, retrievedvalue)
	})
}
