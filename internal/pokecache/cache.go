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
	data map[string]cacheEntry
	mu   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	ticker := time.NewTicker(interval)
	cache := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go func() {
		for t := range ticker.C {
			cache.reapLoop(interval, t)
		}
	}()
	return cache
}

func (c *Cache) Add(k string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[k]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration, now time.Time) {
	for k, v := range c.data {
		if now.Sub(v.createdAt) > interval {
			delete(c.data, k)
		}
	}
}
