package ref

type PokedexInfo struct {
	Name string
}

type Pokedexes struct {
	NATIONAL PokedexInfo
}

var POKEDEXES = Pokedexes{
	NATIONAL: PokedexInfo{
		Name: "national",
	},
}
