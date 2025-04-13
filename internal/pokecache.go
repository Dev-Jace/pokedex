package main

import (
	"fmt"
	"os"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func pokecache() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	// init Cache

	return nil
}
