package builder

import "pokegrep/internal/ref"

func buildPokemonSpecies(p_lang string, p_gen ref.GenerationInfo) bool {
	// switch p_gen {
	// case GENERATIONS.GEN1,
	// 	GENERATIONS.GEN2,
	// 	GENERATIONS.GEN8.LA,
	// 	GENERATIONS.GEN9.LZA:
	// 	if p_gen == GENERATIONS.GEN1 {
	// 		// Do special stat treatment
	// 	}
	// 	// Don't fetch ability
	// 	break
	// default:
	// 	break
	// }
	// pokemon, _ := db.GetById[models.PokemonSpecies](db.DirPokemonSpecies, 6)
	// data := struct {
	// 	Name string
	// }{
	// 	Name: helpers.GetTranslated(pokemon, helpers.NamesField, p_lang),
	// }
	// fmt.Println(data.Name)
	// templates := []string{
	// 	"templates/layouts/base.html",
	// 	"templates/partials/navbar.html",
	// 	"templates/partials/footer.html",
	// 	"templates/pokemon.html",
	// }

	// return buildTemplate(templates, nil, "dist/"+p_lang+"/"+p_gen+"/pokemon/"+strconv.Itoa(pokemon.ID))
	return true
}
