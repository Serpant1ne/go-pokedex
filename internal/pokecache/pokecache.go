package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(i time.Duration) Cache {
	return Cache{}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, found := c.cache[key]
	return entry.val, found
}

func (c *Cache) Set(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) setReapLoop(i time.Duration) {
	ticker := time.NewTicker(i)
	for {
		<-ticker.C

		for key, entry := range c.cache {
			if entry.createdAt.Before(time.Now().Add(-i)) {
				delete(c.cache, key)
			}
		}
	}

}
