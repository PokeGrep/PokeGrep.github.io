package models

type Chain struct {
	EvolutionDetails []any   `json:"evolution_details"`
	EvolvesTo        []Chain `json:"evolves_to"`
	IsBaby           bool    `json:"is_baby"`
	Species          struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

type EvolutionChain struct {
	BabyTriggerItem any   `json:"baby_trigger_item"`
	Chain           Chain `json:"chain"`
	ID              int   `json:"id"`
}
