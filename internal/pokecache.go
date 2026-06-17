package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	data map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}


func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval: interval,
	}
	return cache
}

func (c *Cache) Add(key string, val []byte){
	c.mu.Lock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	retrievedData, ok := c.data[key]
	defer c.mu.Unlock()
	if ok{
		return retrievedData.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(){
	c.mu.Lock()
	for k, v := range c.data{
		dataInterval := time.Since(v.createdAt)
		if dataInterval > c.interval{
			delete(c.data, k)
		}
	}
	c.mu.Unlock()
}


