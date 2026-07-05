package builder

type Gen8Versions struct {
	SS   string
	LA   string
	BDSP string
}

type Gen9Versions struct {
	SV        string
	LZA       string
	CHAMPIONS string
}

type Generations struct {
	GEN1 string
	GEN2 string
	GEN3 string
	GEN4 string
	GEN5 string
	GEN6 string
	GEN7 string
	GEN8 Gen8Versions
	GEN9 Gen9Versions
}

var GENERATIONS = Generations{
	GEN1: "generation-i",
	GEN2: "generation-ii",
	GEN3: "generation-iii",
	GEN4: "generation-iv",
	GEN5: "generation-v",
	GEN6: "generation-vi",
	GEN7: "generation-vii",
	GEN8: Gen8Versions{
		SS:   "sword-shield",
		LA:   "legends-arceus",
		BDSP: "brilliant-diamond-shining-pearl",
	},
	GEN9: Gen9Versions{
		SV:        "scarlet-violet",
		LZA:       "legends-za",
		CHAMPIONS: "champions",
	},
}

func isVersionGroup(p_gen string) bool {
	switch p_gen {
	case GENERATIONS.GEN8.SS,
		GENERATIONS.GEN8.LA,
		GENERATIONS.GEN8.BDSP,
		GENERATIONS.GEN9.SV,
		GENERATIONS.GEN9.LZA,
		GENERATIONS.GEN9.CHAMPIONS:
		return true
	default:
		return false
	}
}

func buildGeneration(p_lang string, p_gen string) bool {
	return buildAbility(p_lang, p_gen) &&
		buildItem(p_lang, p_gen) &&
		buildMove(p_lang, p_gen) &&
		buildPokemon(p_lang, p_gen)
}
