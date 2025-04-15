package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	PS "github.com/Dev-Jace/pokedex/internal/pokecache"
	web_pull "github.com/Dev-Jace/pokedex/internal/web_pull"
)

type Pokemon_Capture struct {
	CaptureRate int `json:"capture_rate"`
}

func commandCatch(config *Config, supCommand string) error {

	if supCommand == "" {
		return fmt.Errorf("no pokemon was given")
	}

	//pull data from PokeAPI site for catch rate
	pokeAPI_URL := "https://pokeapi.co/api/v2/pokemon-species/" + supCommand
	//check cache
	body, exists := config.cachePntr.Get(pokeAPI_URL)
	if !exists {
		//if cache doesn't have add to cache new pull
		var errS string
		body, errS = web_pull.Web_pull(pokeAPI_URL)
		if errS != "" {
			//fmt.Println(errS)
			fmt.Println("catch attempt failed")
			return nil
		} else {
			config.cachePntr.Add(pokeAPI_URL, body)
		}
	}
	//pull data from PokeAPI site for stats
	pokeAPI_URL2 := "https://pokeapi.co/api/v2/pokemon/" + supCommand
	//check cache
	body2, exists2 := config.cachePntr.Get(pokeAPI_URL2)
	if !exists2 {
		//if cache doesn't have add to cache new pull
		var errS string
		body2, errS = web_pull.Web_pull(pokeAPI_URL2)
		if errS != "" {
			//fmt.Println(errS)
			fmt.Println("catch attempt failed")
			return nil
		} else {
			config.cachePntr.Add(pokeAPI_URL2, body2)
		}
	}

	//unmarshall catch rate
	var pokeCap Pokemon_Capture
	err := json.Unmarshal(body, &pokeCap)
	if err != nil {
		fmt.Printf("\nerrr~decode error: %v~\n", err)
		return nil
	}
	//unmarshall stats
	var pokeStats PS.PokeStats
	err2 := json.Unmarshal(body2, &pokeStats)
	if err2 != nil {
		fmt.Printf("\nerrr2~decode error: %v~\n", err2)
		return nil
	}

	//	//catch mechanics
	fmt.Printf("Throwing a Pokeball at %v...\n", supCommand)

	//config rand
	if rand.Intn(255) < pokeCap.CaptureRate {
		fmt.Println(supCommand, " was caught!")
		config.cachePntr.AddPKMN(pokeAPI_URL, supCommand, pokeStats)
	} else {
		fmt.Println(supCommand, " escaped!")
	}

	//
	return nil
}
