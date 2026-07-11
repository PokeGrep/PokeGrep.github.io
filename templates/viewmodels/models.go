package viewmodels

import "pokegrep/internal/localization"

type Page struct {
	Lang   string
	Title  string
	Layout localization.LayoutLabels
}

type GenerationPokemonType struct {
	Name      string
	Shortname string
}

type GenerationPokemonTypes []struct {
	Slot int
	Type GenerationPokemonType
}

type GenerationPokemon struct {
	Lang          string
	GenerationURL string
	PokedexId     int
	Name          string
	Shortname     string
	SpriteURL     string
	PrimaryType   string
	Types         GenerationPokemonTypes
}

type GenerationPage struct {
	Page
	GenerationName     string
	Labels             localization.GenerationLabels
	GenerationPokemons []GenerationPokemon
}
