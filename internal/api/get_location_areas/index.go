package api_get_location_areas

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	pokecache "github.com/git-antonyuk/pokedex-cli/internal"
	"github.com/git-antonyuk/pokedex-cli/internal/api"
)

type LocationItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationItem `json:"results"`
}

func GetLocationAreas(cache pokecache.Cache, url string) (Location, error) {
	// Reading data from cache
	cacheRes, _ := cache.Get(url)
	if cacheRes != nil {
		location, err := api.BytesToJson[Location](cacheRes)
		if (err == nil) {
			return location, nil
		}
	}
	baseApiUrl := url
	if baseApiUrl == "" {
		baseApiUrl = fmt.Sprintf("%v/location-area", api.GetApiUrl())
	}
	res, err := http.Get(baseApiUrl)
	if err != nil {
		log.Fatal(err)
		return Location{}, errors.New("Error while getting locations")
	}
	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		var errorMessage = fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		log.Fatal(errorMessage)
		return Location{}, errors.New(errorMessage)
	}
	if err != nil {
		log.Fatal(err)
		return Location{}, errors.New("Error in reading body")
	}
	locationsData := Location{}
	err = json.Unmarshal(data, &locationsData)
	if err != nil {
		log.Fatal(err)
		return Location{}, errors.New("Error with parsing json")
	}
	// Addint data to cache
	localionBytes, err := api.JsonToBytes(locationsData)
	if err == nil {
		cache.Add(url, localionBytes)
	}
	return locationsData, nil
}
