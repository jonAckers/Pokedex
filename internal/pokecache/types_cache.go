package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mtx 	*sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entries: make(map[string]cacheEntry),
		mtx:     &sync.Mutex{},
	}

	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)

		c.mtx.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mtx.Unlock()
	}
}
