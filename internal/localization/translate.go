package localization

import "fmt"

// Get returns the Translation for the given language and version group/generation name.
func Get(lang string, genName string) Translation {
	var t Translation
	switch lang {
	case LOCALES.FRENCH:
		t = fr
		t.Generation.Subtitle = fmt.Sprintf("Explorez les Pokémon de la %s.", genName)
	case LOCALES.ENGLISH:
		t = en
		t.Generation.Subtitle = fmt.Sprintf("Explore the Pokémon of the %s.", genName)
	default:
		t = en
		t.Generation.Subtitle = fmt.Sprintf("Explore the Pokémon of the %s.", genName)
	}

	// Dynamically compute the language switcher URLs
	t.Layout.LangEnUrl = "/" + LOCALES.ENGLISH + "/" + genName + "/"
	t.Layout.LangFrUrl = "/" + LOCALES.FRENCH + "/" + genName + "/"

	return t
}

// GetLayoutLabels is a compatibility helper that retrieves LayoutLabels.
func GetLayoutLabels(lang string, genName string) LayoutLabels {
	return Get(lang, genName).Layout
}

// GetGenerationLabels is a compatibility helper that retrieves GenerationLabels.
func GetGenerationLabels(lang string, genName string) GenerationLabels {
	return Get(lang, genName).Generation
}

// GetPokemonTemplateLabels is a compatibility helper that retrieves PokemonTemplateLabels.
func GetPokemonTemplateLabels(lang string) PokemonTemplateLabels {
	return Get(lang, "").Pokemon
}
