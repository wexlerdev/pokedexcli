package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	cacheMap	map[string]cacheEntry
	mu			sync.Mutex
}

func NewCache(interval time.Duration) *Cache  {

	cache  := &Cache{
		cacheMap: map[string]cacheEntry{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c * Cache) Add(key string, val []byte) {
	 if c == nil {
		return 
	 }

	 c.mu.Lock()
	 defer c.mu.Unlock()

	 c.cacheMap[key] = cacheEntry{
		 createdAt: time.Now(),
		 val: val,
	 }

	 return
}

func (c * Cache) Get(key string) ([]byte, bool) {
	if c == nil {
		return nil, false
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheMap[key]

	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c * Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for currentTime := range ticker.C{
			c.mu.Lock()

			var cacheValueLifetime time.Duration 
			for key, entry := range c.cacheMap {
				cacheValueLifetime = currentTime.Sub(entry.createdAt)

				if cacheValueLifetime > interval {
					delete(c.cacheMap, key)
				}
			}
			c.mu.Unlock()
		}
	}()

	return
}

