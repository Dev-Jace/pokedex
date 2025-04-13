package main

import (
	"fmt"
	//"os"
)

func commandHelp(config *Config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help : Displays a help message
exit : Exit the Pokedex
map  : displays map information, 
	-if already called, displays next page
mapb : displays prevoius map information, 
	-if "map" was called before`)
	return nil
}
