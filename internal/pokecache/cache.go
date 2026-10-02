package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           *sync.Mutex
	interval     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry, found := c.cacheEntries[key]
	if !found {
		return nil, false
	}
	return cacheEntry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for t := range ticker.C {
		cutoff := t.Add(-c.interval)
		c.mu.Lock()
		for key, val := range c.cacheEntries {
			if val.createdAt.Before(cutoff) {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval:     interval,
		cacheEntries: map[string]cacheEntry{},
		mu:           &sync.Mutex{},
	}
	go cache.reapLoop()
	return cache
}
