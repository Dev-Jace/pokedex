package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	interval time.Duration
	quitChan chan struct{}
	entry    map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(URL string, body []byte) {
	fmt.Println("~Add Called~")
	//lock mutex
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[URL] = cacheEntry{
		createdAt: time.Now(),
		val:       body,
	}
	fmt.Println("add-cache len:", len(c.entry))
	return
}
func (c *Cache) Get(URL string) ([]byte, bool) {
	fmt.Println("~Get Called~")
	fmt.Println("get-cache len:", len(c.entry))
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.entry[URL]
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// Do reaping
			c.mu.Lock()
			for key := range c.entry {
				if time.Since(c.entry[key].createdAt).Seconds() > c.interval.Seconds() {
					delete(c.entry, key)
				}
			}
			//unlock mutex
			c.mu.Unlock()
		case <-c.quitChan:
			return // Exit the goroutine when quit channel is closed
		}

	}

}
func (c *Cache) Stop() { // got from Boots
	close(c.quitChan)
	// Could wait here if needed to ensure reapLoop has exited
}

func NewCache(interval time.Duration) *Cache {
	new_cache := &Cache{
		interval: interval,
		quitChan: make(chan struct{}),
		entry:    map[string]cacheEntry{},
	}
	//call reapLoop
	go new_cache.reapLoop()

	return new_cache
}
