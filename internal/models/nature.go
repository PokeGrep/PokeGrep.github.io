package models

type Nature struct {
	DecreasedStat              any `json:"decreased_stat"`
	HatesFlavor                any `json:"hates_flavor"`
	ID                         int `json:"id"`
	IncreasedStat              any `json:"increased_stat"`
	LikesFlavor                any `json:"likes_flavor"`
	MoveBattleStylePreferences []struct {
		HighHpPreference int `json:"high_hp_preference"`
		LowHpPreference  int `json:"low_hp_preference"`
		MoveBattleStyle  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move_battle_style"`
	} `json:"move_battle_style_preferences"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokeathlonStatChanges []struct {
		MaxChange      int `json:"max_change"`
		PokeathlonStat struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokeathlon_stat"`
	} `json:"pokeathlon_stat_changes"`
}
