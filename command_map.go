package main

import (
	"encoding/json"
	"fmt"

	web_pull "github.com/Dev-Jace/pokedex/internal/web_pull"
)

type Map_Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *Config, supCommand string) error {

	//pull data from PokeAPI site
	pokeAPI_URL := "https://pokeapi.co/api/v2/location-area/"

	if config.next_URL != "none" {
		pokeAPI_URL = config.next_URL // use the next location call provided if map has already been called
	}

	//check cache
	body, exists := config.cachePntr.Get(pokeAPI_URL)
	if !exists {
		//if cache doesn't have add to cache new pull
		var errS string
		body, errS = web_pull.Web_pull(pokeAPI_URL)
		if errS != "" {
			fmt.Println(errS)
		} else {
			config.cachePntr.Add(pokeAPI_URL, body)
		}
	}

	var map_locations Map_Location
	err := json.Unmarshal(body, &map_locations)
	if err != nil {
		fmt.Printf("\nerrr~decode error: %v~\n", err)
		return nil
	}

	//update urls for prev/next
	if map_locations.Next != "" {
		config.next_URL = map_locations.Next
	} else if map_locations.Next == "" {
		fmt.Println("~end of map data~")
	}
	if map_locations.Previous != "" {
		config.prev_URL = map_locations.Previous
	} else if map_locations.Previous == "" {
		config.prev_URL = "none"
	}

	for i := range map_locations.Results {
		fmt.Println(map_locations.Results[i].Name)
	}
	return nil
}

func commandMapb(config *Config, supCommand string) error {
	pokeAPI_URL := "https://pokeapi.co/api/v2/location-area/"

	if config.prev_URL == "none" {
		fmt.Println("you're on the first page")
		return nil
	} else if config.next_URL != "none" {
		pokeAPI_URL = config.prev_URL // use the next location call provided if map has already been called
	}

	//check cache
	body, exists := config.cachePntr.Get(pokeAPI_URL)
	fmt.Println("~~had value in cache:", exists)
	if !exists {
		//if cache doesn't have add to cache new pull
		var errS string
		body, errS = web_pull.Web_pull(pokeAPI_URL)
		if errS != "" {
			fmt.Println(errS)
		}
		config.cachePntr.Add(pokeAPI_URL, body)
	}

	var map_locations Map_Location
	err := json.Unmarshal(body, &map_locations)
	if err != nil {
		fmt.Printf("\nerrr~decode error: %v~\n", err)
		return nil
	}

	//update urls for prev/next
	if map_locations.Next != "" {
		config.next_URL = map_locations.Next
	} else if map_locations.Next == "" {
		config.next_URL = "none"
	}
	if map_locations.Previous != "" {
		config.prev_URL = map_locations.Previous
	} else if map_locations.Previous == "" {
		config.prev_URL = "none"
	}

	for i := range map_locations.Results {
		fmt.Println(map_locations.Results[i].Name)
	}

	return nil
}
