package pokecache

import (
	"fmt" //for testing
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	interval time.Duration
	quitChan chan struct{}
	entry    map[string]cacheEntry
	pokeDex  map[string]Pokemon
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Pokemon struct {
	stats    PokeStats
	quantity int
}
type PokeStats struct {
	Height int `json:"height"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (c *Cache) Add(URL string, body []byte) {
	//lock mutex
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[URL] = cacheEntry{
		createdAt: time.Now(),
		val:       body,
	}
}
func (c *Cache) Get(URL string) ([]byte, bool) {
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
		pokeDex:  map[string]Pokemon{},
	}
	//call reapLoop
	go new_cache.reapLoop()

	return new_cache
}

func (c *Cache) AddPKMN(URL string, name string, poke_stat PokeStats) {
	//lock mutex
	c.mu.Lock()
	defer c.mu.Unlock()

	pokemon, exists := c.pokeDex[name]
	if exists {
		pokemon.quantity++
	} else {
		c.pokeDex[name] = Pokemon{
			stats:    poke_stat,
			quantity: 1,
		}
	}

}
func (c *Cache) GetPKMN(name string) string {
	//lock mutex
	c.mu.Lock()
	defer c.mu.Unlock()
	_, exists := c.pokeDex[name]
	if exists {
		return c.statFormatter(name)
	} else {
		return "You have not yet caught that Pokemon"
	}

}
func (c *Cache) GetPKMNList() string {
	//lock mutex
	c.mu.Lock()
	defer c.mu.Unlock()
	exists := c.pokeDex
	response := ""
	if len(exists) > 0 {
		for key, _ := range exists {
			response += fmt.Sprintf("- %v \n", key)
		}
		return response
	} else {
		return "You have not yet caught any Pokemon"
	}

}
func (c *Cache) statFormatter(name string) string {
	pkmn := c.pokeDex[name]
	statSheet := fmt.Sprintf(`Name: %v
Height: %v
Weight: %v
Stats:
	-hp: %v
	-attack: %v
	-defense: %v
	-special-attack: %v
	-special-defense: %v
	-speed: %v
Types:`, name, pkmn.stats.Height, pkmn.stats.Weight,
		pkmn.stats.Stats[0].BaseStat,
		pkmn.stats.Stats[1].BaseStat,
		pkmn.stats.Stats[2].BaseStat,
		pkmn.stats.Stats[3].BaseStat,
		pkmn.stats.Stats[4].BaseStat,
		pkmn.stats.Stats[5].BaseStat)
	for i := range pkmn.stats.Types {
		statSheet += fmt.Sprintf("\n\t- %v", pkmn.stats.Types[i].Type.Name)
	}
	return statSheet
}
