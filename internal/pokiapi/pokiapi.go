package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
}

type locationReqRes struct {
	Count    int    `json:"count"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempt"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func GetLocations(url string) (locationReqRes, error) {
    res, err := http.Get(url)
    if err != nil {
        return locationReqRes{}, err
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return locationReqRes{}, err
    }
    defer res.Body.Close()

    locData := locationReqRes{}
    err = json.Unmarshal(body, &locData)
    if err != nil {
        return locationReqRes{}, err
    }

    return locData, nil

}
