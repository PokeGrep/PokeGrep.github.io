package builder

func buildLanguage(p_lang string) bool {
	return buildGeneration(p_lang, GENERATIONS.GEN1)
}
