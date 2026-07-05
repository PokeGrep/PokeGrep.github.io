package builder

func buildAbility(p_lang string, p_gen string) bool {
	switch p_gen {
	case GENERATIONS.GEN1,
		GENERATIONS.GEN2,
		GENERATIONS.GEN8.LA,
		GENERATIONS.GEN9.LZA:
		return true
	default:
		break
	}
	return true
}
