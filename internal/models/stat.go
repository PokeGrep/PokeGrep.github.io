package models

type Stat struct {
	AffectingItems []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"affecting_items"`
	AffectingMoves struct {
		Decrease []any `json:"decrease"`
		Increase []any `json:"increase"`
	} `json:"affecting_moves"`
	AffectingNatures struct {
		Decrease []any `json:"decrease"`
		Increase []any `json:"increase"`
	} `json:"affecting_natures"`
	Characteristics []struct {
		URL string `json:"url"`
	} `json:"characteristics"`
	GameIndex       int    `json:"game_index"`
	ID              int    `json:"id"`
	IsBattleOnly    bool   `json:"is_battle_only"`
	MoveDamageClass any    `json:"move_damage_class"`
	Name            string `json:"name"`
	Names           []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}
