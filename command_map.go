package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func commandMap(config *Config) error {
	//fmt.Println(`This will be a map utility`) //test functionality of call
	//implement pokeAPI
	// https://www.boot.dev/lessons/813eafe1-2e1d-42a0-b358-53e0f4d4fdc8

	//pull data from PokeAPI site
	pokeAPI_URL := "https://pokeapi.co/api/v2/location-area/"

	if config.next_URL != "none" {
		pokeAPI_URL = config.next_URL // use the next location call provided if map has already been called
	}

	res, err := http.Get(pokeAPI_URL)
	if err != nil {
		fmt.Printf("\n~website error: %v~\n", err)
		return nil
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("\n~Response failed with status code: %d and\nbody: %s~\n", res.StatusCode, body)
		return nil
	}
	if err != nil {
		fmt.Printf("\n~~read error: %v~\n", err)
		return nil
	}

	var map_locations Map_Location
	errr := json.Unmarshal(body, &map_locations)
	if errr != nil {
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

func commandMapb(config *Config) error {
	pokeAPI_URL := "https://pokeapi.co/api/v2/location-area/"

	if config.prev_URL == "none" {
		fmt.Println("you're on the first page")
		return nil
	} else if config.next_URL != "none" {
		pokeAPI_URL = config.prev_URL // use the next location call provided if map has already been called
	}

	res, err := http.Get(pokeAPI_URL)
	if err != nil {
		fmt.Printf("\n~website error: %v~\n", err)
		return nil
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("\n~Response failed with status code: %d and\nbody: %s~\n", res.StatusCode, body)
		return nil
	}
	if err != nil {
		fmt.Printf("\n~~read error: %v~\n", err)
		return nil
	}

	var map_locations Map_Location
	errr := json.Unmarshal(body, &map_locations)
	if errr != nil {
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
