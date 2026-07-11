package builder

import (
	"fmt"
	"pokegrep/internal/db"
	"pokegrep/internal/localization"
	"pokegrep/internal/models"
	"pokegrep/internal/models/helpers"
	"pokegrep/internal/ref"
	"pokegrep/templates"
	"pokegrep/templates/partials"
	"pokegrep/templates/viewmodels"
	"sort"
)

func buildGeneration(p_lang string, p_gen ref.GenerationInfo) bool {
	outputPath := "dist/" + p_lang + "/" + p_gen.Name + "/index.html"
	dependencies := []string{
		"templates/generation.templ",
		"templates/layouts/base.templ",
		"templates/partials/navbar.templ",
		"templates/partials/footer.templ",
		"templates/components/pokemon_card.templ",
	}

	if !shouldRebuild(outputPath, dependencies...) {
		return buildAbility(p_lang, p_gen) &&
			buildItem(p_lang, p_gen) &&
			buildMove(p_lang, p_gen) &&
			buildPokemonSpecies(p_lang, p_gen)
	}

	var genTrans string
	var genPokemons []viewmodels.GenerationPokemon

	// If p_gen refers as a Generation (I-VII) instead of a Version Group
	// (sword-shield, scarlet-violet..)
	if !p_gen.IsVersionGroup() {
		// Parse all the gens from Gen 1 to p_gen
		for i := p_gen.Order; i > ref.GENERATIONS.GEN1.Order-1; i-- {
			genModel, err := db.GetById[models.Generation](
				db.DirGeneration,
				(p_gen.Order-i)+1,
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
				var pokemonTypes viewmodels.GenerationPokemonTypes

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
						Type viewmodels.GenerationPokemonType
					}{
						Slot: entry.GetTypes(p_gen)[i].Slot,
						Type: viewmodels.GenerationPokemonType{
							Name:      helpers.GetTranslated(pokemonType, helpers.NamesField, p_lang),
							Shortname: pokemonType.Name,
						},
					})
				}

				generationPokemon := viewmodels.GenerationPokemon{
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
				pokemonTypes = viewmodels.GenerationPokemonTypes{}

				// Sort genPokemons by their pokedex id
				sort.Slice(genPokemons, func(i1, i2 int) bool {
					return genPokemons[i1].PokedexId < genPokemons[i2].PokedexId
				})
			}
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

	translation := localization.Get(p_lang, p_gen.Name)
	data := viewmodels.GenerationPage{
		Page: viewmodels.Page{
			Lang:   p_lang,
			Title:  genTrans,
			Layout: translation.Layout,
		},
		GenerationName:     genTrans,
		Labels:             translation.Generation,
		GenerationPokemons: genPokemons,
	}

	navBar := partials.Navbar(p_lang, genTrans, translation.Layout)
	footer := partials.Footer(translation.Layout)
	component := templates.Generation(data, navBar, footer)

	return renderToFile(component, "dist/"+p_lang+"/"+p_gen.Name+"/index.html") &&
		buildAbility(p_lang, p_gen) &&
		buildItem(p_lang, p_gen) &&
		buildMove(p_lang, p_gen) &&
		buildPokemonSpecies(p_lang, p_gen)
}
