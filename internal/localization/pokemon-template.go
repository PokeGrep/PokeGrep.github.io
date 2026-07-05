package localization

type characteristicsBlock struct {
	Title                 string
	Size                  string
	Weight                string
	EggGroups             string
	WeaknessesResistances string
	Stats                 string
}

type evolutionChainBlock struct {
	Title string
}

type movesetBlock struct {
	Title       string
	Capacity    string
	Type        string
	Category    string
	Power       string
	Accuracy    string
	Method      string
	GameDetails string
}

type localizationsBlock struct {
	Title       string
	Method      string
	Area        string
	GameVersion string
	Chance      string
	Levels      string
}

type PokemonTemplateLabels struct {
	Characteristics characteristicsBlock
	EvolutionChain  evolutionChainBlock
	Moveset         movesetBlock
	Localizations   localizationsBlock
}

var pokemonTemplateLabelsFR = PokemonTemplateLabels{
	Characteristics: characteristicsBlock{
		Title:                 "Caractéristiques",
		Size:                  "Taille",
		Weight:                "Poids",
		EggGroups:             "Groupe d'œufs",
		WeaknessesResistances: "Résistances et Faiblesses",
		Stats:                 "Statistiques",
	},
	EvolutionChain: evolutionChainBlock{
		Title: "Chaîne d'Évolution",
	},
	Moveset: movesetBlock{
		Title:       "Liste des capacités",
		Capacity:    "Capacité",
		Type:        "Type",
		Category:    "Catégorie",
		Power:       "Puissance",
		Accuracy:    "Précision",
		Method:      "Méthode",
		GameDetails: "Détails (Jeu)",
	},
	Localizations: localizationsBlock{
		Title:       "Localisations",
		Method:      "Méthode",
		Area:        "Zone",
		GameVersion: "Versions",
		Levels:      "Niveaux",
	},
}
var pokemonTemplateLabelsEN = PokemonTemplateLabels{
	Characteristics: characteristicsBlock{
		Title:                 "Caracteristics",
		Size:                  "Size",
		Weight:                "Weight",
		EggGroups:             "Egg Groups",
		WeaknessesResistances: "Resistances and Weaknesses",
		Stats:                 "Statistics",
	},
	EvolutionChain: evolutionChainBlock{
		Title: "Evolution Chain",
	},
	Moveset: movesetBlock{
		Title:       "Capacity list",
		Capacity:    "Capacity",
		Type:        "Type",
		Category:    "Category",
		Power:       "Power",
		Accuracy:    "Accuracy",
		Method:      "Method",
		GameDetails: "Details (Game)",
	},
	Localizations: localizationsBlock{
		Title:       "Localizations",
		Method:      "Method",
		Area:        "Area",
		GameVersion: "Versions",
		Levels:      "Levels",
	},
}

func GetPokemonTemplateLabels(p_lang string) PokemonTemplateLabels {
	switch p_lang {
	case LOCALES.FRENCH:
		return pokemonTemplateLabelsFR
	case LOCALES.ENGLISH:
		return pokemonTemplateLabelsEN
	}
	return pokemonTemplateLabelsEN
}
