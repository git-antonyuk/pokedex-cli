package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu     *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		cutoff := time.Now().Add(-interval)
		c.mu.Lock()
		for key, entry := range c.entries {
			if entry.createdAt.Before(cutoff) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.entries[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exist := c.entries[key]
	// if (exist) {
	// 	fmt.Printf("✅ Getting data from cache, key: %v", key)
	// } else {
	// 	fmt.Printf("⚠️ No data from cache, need to fetch item for key: %v", key)
	// }
	return entry.val, exist
}
