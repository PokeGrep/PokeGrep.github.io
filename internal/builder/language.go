package builder

import "pokegrep/internal/ref"

func buildLanguage(p_lang string) bool {
	return buildGeneration(p_lang, ref.GENERATIONS.GEN1) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN2) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN3) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN4) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN5) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN6) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN7.GEN7) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN7.LGPE) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN8.SS) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN8.LA) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN8.BDSP) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN9.SV) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN9.LZA) &&
		buildGeneration(p_lang, ref.GENERATIONS.GEN9.CHAMPIONS)
}
