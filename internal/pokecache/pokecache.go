package pokecache

import (
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(URL string, val []byte) {
	c.entry[URL] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
func (c *Cache) Get(URL string) ([]byte, bool) {
	val, exists := c.entry[URL]
	if !exists {
		return nil, false
	} else {
		return val.val, true
	}
}

func (c *Cache) reapLoop() {
	for key, _ := range c.entry {
		if time.Since(c.entry[key].createdAt).Seconds() > 5 {
			delete(c.entry, key)
		}
	}
}

func NewCache(interval time.Duration) {
	//make(map[string]int)

}
