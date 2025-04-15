package main

import (
	"fmt"
	"os"
)

func commandExit(config *Config, supCommand string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	config.cachePntr.Stop()
	os.Exit(0)
	return nil
}
