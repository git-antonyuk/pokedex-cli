package api_get_location_area_details

func GetPokemonsList(details LocationAreaDetail) []string {
	list := []string{}
	for _, pokemon := range details.PokemonEncounters {
		list = append(list, pokemon.Pokemon.Name)
	}
	return list
}
