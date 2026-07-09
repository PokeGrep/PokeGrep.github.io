package builder

import (
	"fmt"
	"pokegrep/internal/db"
	"pokegrep/internal/localization"
	"pokegrep/internal/models"
	"pokegrep/internal/models/helpers"
	"pokegrep/internal/ref"
	"sort"
)

func buildGeneration(p_lang string, p_gen ref.GenerationInfo) bool {
	tmpl := []string{
		"templates/layouts/base.html",
		"templates/partials/navbar.html",
		"templates/partials/footer.html",
		"templates/components/pokemon_card.html",
		"templates/generation.html",
	}

	// langModel, err := db.GetByName[models.Language](
	// 	db.DirLanguage,
	// 	p_lang,
	// )
	// if err != nil {
	// 	panic("langModel not found")
	// }
	// langTrans := helpers.GetTranslated(
	// 	langModel,
	// 	helpers.NamesField,
	// 	p_lang,
	// )

	var genTrans string
	var genPokemons []GenerationPokemon

	// If p_gen refers as a Generation (I-VII) instead of a Version Group
	// (sword-shield, scarlet-violet..)
	if !p_gen.IsVersionGroup() {
		genModel, err := db.GetByName[models.Generation](
			db.DirGeneration,
			p_gen.Name,
		)
		if err != nil {
			panic("genModel not found")
		}
		genTrans = helpers.GetTranslated(
			genModel,
			helpers.NamesField,
			p_lang,
		)

		pokemonSpecies := genModel.PokemonSpecies
		for _, pokemonSpecie := range pokemonSpecies {
			id := helpers.GetResourceId(pokemonSpecie.URL)
			entry, err := db.GetById[models.PokemonSpecies](
				db.DirPokemonSpecies,
				id,
			)
			if err != nil {
				panic(fmt.Sprintf("pokemon-species with %d not found", id))
			}

			var primaryType string
			var pokemonTypes GenerationPokemonTypes

			for i := range entry.GetTypes(p_gen) {
				pokemonType, err := db.GetByName[models.Type](db.DirType, entry.GetTypes(p_gen)[i].Type.Name)
				if err != nil {
					panic(err)
				}

				if entry.GetTypes(p_gen)[i].Slot == 1 || primaryType == "" {
					primaryType = pokemonType.Name
				}

				pokemonTypes = append(pokemonTypes, struct {
					Slot int
					Type GenerationPokemonType
				}{
					Slot: entry.GetTypes(p_gen)[i].Slot,
					Type: GenerationPokemonType{
						Name:      helpers.GetTranslated(pokemonType, helpers.NamesField, p_lang),
						Shortname: pokemonType.Name,
					},
				})
			}

			generationPokemon := GenerationPokemon{
				Lang:          p_lang,
				GenerationURL: p_gen.Name,
				PokedexId:     entry.GetPokedexEntryNumber(),
				Name:          helpers.GetTranslated(entry, helpers.NamesField, p_lang),
				Shortname:     entry.Name,
				SpriteURL:     entry.GetSpriteURL(),
				PrimaryType:   primaryType,
				Types:         pokemonTypes,
			}

			genPokemons = append(genPokemons, generationPokemon)

			// Cleanup the types
			pokemonTypes = GenerationPokemonTypes{}

			// Sort genPokemons by their pokedex id
			sort.Slice(genPokemons, func(i1, i2 int) bool {
				return genPokemons[i1].PokedexId < genPokemons[i2].PokedexId
			})
		}

	} else {
		genModel, err := db.GetByName[models.VersionGroup](
			db.DirVersionGroup,
			p_gen.Name,
		)
		if err != nil {
			panic("genModel not found")
		}
		genTrans = helpers.GetTranslated(
			genModel,
			helpers.NamesField,
			p_lang,
		)
	}

	data := GenerationPage{
		Page: Page{
			Lang:   p_lang,
			Title:  genTrans,
			Layout: localization.GetLayoutLabels(p_lang, p_gen.Name),
		},
		GenerationName:     genTrans,
		Labels:             localization.GetGenerationLabels(p_lang, genTrans),
		GenerationPokemons: genPokemons,
	}
	return buildTemplate(tmpl, data, "dist/"+p_lang+"/"+p_gen.Name) &&
		buildAbility(p_lang, p_gen) &&
		buildItem(p_lang, p_gen) &&
		buildMove(p_lang, p_gen) &&
		buildPokemonSpecies(p_lang, p_gen)
}
