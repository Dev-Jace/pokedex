package main

import (
	"fmt"
	//"os"
)

func commandInspect(config *Config, supCommand string) error {
	fmt.Println(`Welcome`)
	//check for pokemon in Dex
	response := config.cachePntr.GetPKMN(supCommand)
	fmt.Println(response)
	return nil
}
