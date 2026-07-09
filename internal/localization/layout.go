package localization

type LayoutLabels struct {
	NavHome           string
	NavGenerations    string
	NavGen1           string
	SearchPlaceholder string
	CurrentLang       string
	LangEnUrl         string
	LangFrUrl         string
	LangEnLabel       string
	LangFrLabel       string
	FooterCopyright   string
	FooterPoweredBy   string
	FooterPokeApiUrl  string
	FooterPokeApiText string
	FooterGithubUrl   string
	FooterGithubText  string
}

var layoutLabelsFR = LayoutLabels{
	NavHome:           "Accueil",
	NavGenerations:    "Générations",
	NavGen1:           "1re Génération",
	SearchPlaceholder: "Rechercher...",
	CurrentLang:       "Français",
	LangEnLabel:       "English",
	LangFrLabel:       "Français",
	FooterCopyright:   "© 2026 PokéGrep. Tous droits réservés.",
	FooterPoweredBy:   "Propulsé par",
	FooterPokeApiUrl:  "https://pokeapi.co",
	FooterPokeApiText: "PokéAPI",
	FooterGithubUrl:   "https://github.com/PokeGrep/PokeGrep.github.io",
	FooterGithubText:  "GitHub",
}

var layoutLabelsEN = LayoutLabels{
	NavHome:           "Home",
	NavGenerations:    "Generations",
	NavGen1:           "1st Generation",
	SearchPlaceholder: "Search...",
	CurrentLang:       "English",
	LangEnLabel:       "English",
	LangFrLabel:       "Français",
	FooterCopyright:   "© 2026 PokéGrep. All rights reserved.",
	FooterPoweredBy:   "Powered by",
	FooterPokeApiUrl:  "https://pokeapi.co",
	FooterPokeApiText: "PokéAPI",
	FooterGithubUrl:   "https://github.com/PokeGrep/PokeGrep.github.io",
	FooterGithubText:  "GitHub",
}

func GetLayoutLabels(p_lang string, genName string) LayoutLabels {
	var labels LayoutLabels
	switch p_lang {
	case LOCALES.FRENCH:
		labels = layoutLabelsFR
	case LOCALES.ENGLISH:
		labels = layoutLabelsEN
	default:
		labels = layoutLabelsEN
	}
	// Dynamically compute the switcher URLs
	labels.LangEnUrl = "/" + LOCALES.ENGLISH + "/" + genName + "/"
	labels.LangFrUrl = "/" + LOCALES.FRENCH + "/" + genName + "/"
	return labels
}
