package api

import "fmt"

const API_URL = "https://pokeapi.co/api"
const API_VERSION = "v2"

func GetApiUrl() string {
	return fmt.Sprintf("%v/%v", API_URL, API_VERSION)
}
