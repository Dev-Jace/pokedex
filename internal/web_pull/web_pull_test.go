package web_pull

import (
	"testing"
)

func TestWeb_pull(t *testing.T) {
	pokeAPI_URL := "https://pokeapi.co/api/v2/location-area/"

	_, errS := Web_pull(pokeAPI_URL)

	if errS != "" {
		t.Errorf("webpull error")
		return
	}

}
