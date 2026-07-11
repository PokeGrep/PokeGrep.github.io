package builder

import (
	"pokegrep/internal/db"
	"pokegrep/internal/models"
	"pokegrep/internal/ref"
)

func buildPokemonSpecies(p_lang string, p_gen ref.GenerationInfo) bool {
	if !p_gen.IsVersionGroup() {
		generation, err := db.GetByName[models.Generation](db.DirGeneration, p_gen.Name)
		if err != nil {
			panic(err)
		}
		for _, entry := range generation.PokemonSpecies {
			specie, err := db.GetByName[models.PokemonSpecies](db.DirPokemonSpecies, entry.Name)
			if err != nil {
				panic(err)
			}

			if !buildPokemonSpecie(p_lang, p_gen, specie) {
				return false
			}
		}
	}

	return true
}

func buildPokemonSpecie(p_lang string, p_gen ref.GenerationInfo, p_specie *models.PokemonSpecies) bool {

	return true
}
