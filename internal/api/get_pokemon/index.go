package api_get_pokemon

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

func GetPolemonInfo(cache pokecache.Cache, name string) (PokemonInfo, error) {
	if name == "" {
		return PokemonInfo{}, errors.New("Error, you have to add name param")
	}
	// Reading data from cache
	cacheRes, _ := cache.Get(name)
	if cacheRes != nil {
		location, err := api.BytesToJson[PokemonInfo](cacheRes)
		if (err == nil) {
			return location, nil
		}
	}
	baseApiUrl := fmt.Sprintf("%v/pokemon/%v", api.GetApiUrl(), name)
	res, err := http.Get(baseApiUrl)
	if err != nil {
		log.Fatal(err)
		return PokemonInfo{}, errors.New("Error while getting locations")
	}
	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		var errorMessage = fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		log.Fatal(errorMessage)
		return PokemonInfo{}, errors.New(errorMessage)
	}
	if err != nil {
		log.Fatal(err)
		return PokemonInfo{}, errors.New("Error in reading body")
	}
	pokemonInfo := PokemonInfo{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		log.Fatal(err)
		return PokemonInfo{}, errors.New("Error with parsing json")
	}
	// Addint data to cache
	localionBytes, err := api.JsonToBytes(pokemonInfo)
	if err == nil {
		cache.Add(name, localionBytes)
	}
	return pokemonInfo, nil
}
