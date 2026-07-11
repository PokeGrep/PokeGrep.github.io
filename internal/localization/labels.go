package localization

type LayoutLabels struct {
	NavHome           string
	NavGenerations    string
	NavGen1           string
	NavGen2           string
	NavGen3           string
	NavGen4           string
	NavGen5           string
	NavGen6           string
	NavGen7           string
	NavComingSoon     string
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

type CharacteristicsBlock struct {
	Title                 string
	Size                  string
	Weight                string
	EggGroups             string
	WeaknessesResistances string
	Stats                 string
}

type EvolutionChainBlock struct {
	Title string
}

type MovesetBlock struct {
	Title       string
	Capacity    string
	Type        string
	Category    string
	Power       string
	Accuracy    string
	Method      string
	GameDetails string
}

type LocalizationsBlock struct {
	Title       string
	Method      string
	Area        string
	GameVersion string
	Chance      string
	Levels      string
}

type PokemonTemplateLabels struct {
	Characteristics CharacteristicsBlock
	EvolutionChain  EvolutionChainBlock
	Moveset         MovesetBlock
	Localizations   LocalizationsBlock
}

type Translation struct {
	Layout     LayoutLabels
	Generation GenerationLabels
	Pokemon    PokemonTemplateLabels
}
