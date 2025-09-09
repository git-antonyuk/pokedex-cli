# Pokédex CLI

A command-line Pokédex application that lets you explore Pokémon location areas using the PokéAPI. Navigate through different areas, discover Pokémon, and build your knowledge of the Pokémon world!

## Quick Start
```sh
go run .
```

## Commands
- `help` - Show available commands
- `map` - Get next 20 location areas
- `mapb` - Go back to previous location areas
- `explore <area-name>` - Explore a specific location area
- `exit` - Exit the application

## Testing logs and save history of commands
```sh
go run . | tee repl.log
```
