package main

import (
	"encoding/json"
	"fmt"

	web_pull "github.com/Dev-Jace/pokedex/internal/web_pull"
)

type Pokemon_available struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(config *Config, supCommand string) error {

	if supCommand == "" {
		return fmt.Errorf("no area was given")
	}
	//pull data from PokeAPI site
	pokeAPI_URL := "https://pokeapi.co/api/v2/location-area/" + supCommand

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

	var poke_avail Pokemon_available
	err := json.Unmarshal(body, &poke_avail)
	if err != nil {
		fmt.Printf("\nerrr~decode error: %v~\n", err)
		return nil
	}

	for i := range poke_avail.PokemonEncounters {
		fmt.Println(" - ", poke_avail.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}
