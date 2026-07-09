package localization

import "fmt"

type GenerationLabels struct {
	Subtitle            string
	SearchPlaceholder   string
	FilterAll           string
	ActiveFilters       string
	NoResultsTitle      string
	NoResultsSubtitle   string
	ResetFilters        string
	NoActiveFiltersText string
}

var generationLabelsFR = GenerationLabels{
	SearchPlaceholder:   "Rechercher par nom ou n°...",
	FilterAll:           "Tous",
	ActiveFilters:       "Filtres actifs :",
	NoResultsTitle:      "Aucun Pokémon ne correspond à votre recherche",
	NoResultsSubtitle:   "Essayez d'ajuster les filtres ou la recherche par nom.",
	ResetFilters:        "Réinitialiser",
	NoActiveFiltersText: "Aucun filtre actif (Tous les Pokémon)",
}

var generationLabelsEN = GenerationLabels{
	SearchPlaceholder:   "Search by name or number...",
	FilterAll:           "All",
	ActiveFilters:       "Active filters:",
	NoResultsTitle:      "No Pokémon matches your criteria",
	NoResultsSubtitle:   "Try adjusting the search query or filters.",
	ResetFilters:        "Reset Filters",
	NoActiveFiltersText: "No active filters (All Pokémon)",
}

func GetGenerationLabels(p_lang string, genName string) GenerationLabels {
	var labels GenerationLabels
	switch p_lang {
	case LOCALES.FRENCH:
		labels = generationLabelsFR
		labels.Subtitle = fmt.Sprintf("Explorez les Pokémon de la %s.", genName)
	case LOCALES.ENGLISH:
		labels = generationLabelsEN
		labels.Subtitle = fmt.Sprintf("Explore the Pokémon of the %s.", genName)
	default:
		labels = generationLabelsEN
		labels.Subtitle = fmt.Sprintf("Explore the Pokémon of the %s.", genName)
	}
	return labels
}
