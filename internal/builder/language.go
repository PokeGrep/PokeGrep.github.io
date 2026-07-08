package builder

import "pokegrep/internal/ref"


func buildLanguage(p_lang string) bool {
	return buildGeneration(p_lang, ref.GENERATIONS.GEN1)
}
