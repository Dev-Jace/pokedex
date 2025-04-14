package web_pull

import (
	"fmt"
	"io"
	"net/http"
)

func Web_pull(pokeAPI_URL string) ([]byte, string) {
	var empty []byte
	var err_type string
	res, err := http.Get(pokeAPI_URL)
	if err != nil {
		err_type = fmt.Sprintf("\n~website error: %v~\n", err)
		return empty, err_type
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err_type = fmt.Sprintf("\n~Response failed with status code: %d and\nbody: %s~\n", res.StatusCode, body)
		return empty, err_type
	}
	if err != nil {
		err_type = fmt.Sprintf("\n~~read error: %v~\n", err)
		return empty, err_type
	}

	return body, ""
}
