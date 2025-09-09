package string_pokemon_info

import (
	"fmt"
	"strings"

	api_get_pokemon "github.com/git-antonyuk/pokedex-cli/internal/api/get_pokemon"
)

func PrintPokemonInspectInfo(info api_get_pokemon.PokemonInfo) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Name: %s\n", info.Name)
	fmt.Fprintf(&sb, "Height: %d\n", info.Height)
	fmt.Fprintf(&sb, "Weight: %d\n", info.Weight)
	fmt.Fprint(&sb, "Stats:\n")
	for _, stat := range info.Stats {
		fmt.Fprintf(&sb, " -%v %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Fprint(&sb, "Types:\n")
	for _, infoType := range info.Types  {
		fmt.Fprintf(&sb, " -%v\n", infoType.Type.Name)
	}
	fmt.Print(sb.String())
}
