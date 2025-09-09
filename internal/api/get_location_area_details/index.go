package api_get_location_area_details

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

func GetLocationAreaDetails(cache pokecache.Cache, name string) (LocationAreaDetail, error) {
	if name == "" {
		return LocationAreaDetail{}, errors.New("Error, you have to add name param")
	}
	// Reading data from cache
	cacheRes, _ := cache.Get(name)
	if cacheRes != nil {
		location, err := api.BytesToJson[LocationAreaDetail](cacheRes)
		if (err == nil) {
			return location, nil
		}
	}
	baseApiUrl := fmt.Sprintf("%v/location-area/%v", api.GetApiUrl(), name)
	res, err := http.Get(baseApiUrl)
	if err != nil {
		log.Fatal(err)
		return LocationAreaDetail{}, errors.New("Error while getting locations")
	}
	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		var errorMessage = fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		log.Fatal(errorMessage)
		return LocationAreaDetail{}, errors.New(errorMessage)
	}
	if err != nil {
		log.Fatal(err)
		return LocationAreaDetail{}, errors.New("Error in reading body")
	}
	locationsData := LocationAreaDetail{}
	err = json.Unmarshal(data, &locationsData)
	if err != nil {
		log.Fatal(err)
		return LocationAreaDetail{}, errors.New("Error with parsing json")
	}
	// Addint data to cache
	localionBytes, err := api.JsonToBytes(locationsData)
	if err == nil {
		cache.Add(name, localionBytes)
	}
	return locationsData, nil
}
