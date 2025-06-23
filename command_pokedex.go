package main

import (
	"fmt"
	//"os"
)

func commandPokedex(config *Config, supCommand string) error {
	response := config.cachePntr.GetPKMNList()
	fmt.Println(response)
	return nil
}
