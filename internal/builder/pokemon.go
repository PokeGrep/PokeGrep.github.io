package builder

func buildPokemon(p_lang string, p_gen string) bool {
	switch p_gen {
	case GENERATIONS.GEN1,
		GENERATIONS.GEN2,
		GENERATIONS.GEN8.LA,
		GENERATIONS.GEN9.LZA:
		if p_gen == GENERATIONS.GEN1 {
			// Do special stat treatment
		}
		// Don't fetch ability
		break
	default:
		break
	}

	return true
}
