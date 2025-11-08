package cache

import (
	"log/slog"
	"sync"

	"github.com/ryanmab/rdap-go/internal/model"
)

// Cache is a simple in-memory cache, indexed by the query time (i.e. domain, IP,
// ASN) and the identifier (i.e. example.com, 8.8.8.8, etc.)
type Cache struct {
	sync.RWMutex
	cache map[model.RdapQuery]map[string]*any
}

// New creates a new in-memory cache instance which can be used to store RDAP responses
// and prevent redundant network requests.
func New() *Cache {
	return &Cache{
		cache:   make(map[model.RdapQuery]map[string]*any),
		RWMutex: sync.RWMutex{},
	}
}

// Get a cached RDAP response for the given query and identifier, or nil if not found.
func (cache *Cache) Get(query model.RdapQuery, identifier string) *any {
	cache.RLock()
	defer cache.RUnlock()

	if serverCache, ok := cache.cache[query]; ok {
		if result, ok := serverCache[identifier]; ok {
			return result
		}
	}
	return nil
}

// Set a cached RDAP response for the given query and identifier.
func (cache *Cache) Set(q model.RdapQuery, identifier string, response any) {
	cache.Lock()
	defer cache.Unlock()

	slog.Debug("Storing RDAP response in cache", "identifier", identifier, "query", q)

	if _, ok := cache.cache[q]; !ok {
		cache.cache[q] = make(map[string]*any)
	}
	cache.cache[q][identifier] = &response
}

// Clear the entire cache.
func (cache *Cache) Clear() {
	cache.Lock()
	defer cache.Unlock()

	cache.cache = make(map[model.RdapQuery]map[string]*any)
}
