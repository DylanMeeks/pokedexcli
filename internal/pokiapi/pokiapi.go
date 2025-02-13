package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
}

func GetLocationArea(id string) ([]string, error) {
    res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + id)
    if err != nil {
        return []string{}, err
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return []string{}, err
    }
    defer res.Body.Close()

    locData := locationArea{}
    err = json.Unmarshal(body, &locData)
    if err != nil {
        return []string{}, err
    }

    return []string{locData.Name}, nil

}
