package builder

// PokeApi don't provide Gen 2's data for items
func buildItem(p_lang string, p_gen string) bool {
	switch p_gen {
	case GENERATIONS.GEN1,
		GENERATIONS.GEN2:
		return true
	default:
		break
	}
	return true
}
