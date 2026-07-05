package models

type PokemonForm struct {
	FormName     string `json:"form_name"`
	FormNames    []any  `json:"form_names"`
	FormOrder    int    `json:"form_order"`
	ID           int    `json:"id"`
	IsBattleOnly bool   `json:"is_battle_only"`
	IsDefault    bool   `json:"is_default"`
	IsMega       bool   `json:"is_mega"`
	Name         string `json:"name"`
	Names        []any  `json:"names"`
	Order        int    `json:"order"`
	Pokemon      struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
		Versions         struct {
			GenerationIx struct {
				ScarletViolet struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"scarlet-violet"`
			} `json:"generation-ix"`
			GenerationViii struct {
				BrilliantDiamondShiningPearl struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"brilliant-diamond-shining-pearl"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	VersionGroup struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_group"`
}
