package ref

import (
	"fmt"
	"reflect"
)

type GenerationInfo struct {
	Order int
	Name  string
}

type Generation7 struct {
	GEN7 GenerationInfo
	LGPE GenerationInfo
}

type Generation8 struct {
	SS   GenerationInfo
	LA   GenerationInfo
	BDSP GenerationInfo
}

type Generation9 struct {
	SV        GenerationInfo
	LZA       GenerationInfo
	CHAMPIONS GenerationInfo
}

type Generations struct {
	GEN1 GenerationInfo
	GEN2 GenerationInfo
	GEN3 GenerationInfo
	GEN4 GenerationInfo
	GEN5 GenerationInfo
	GEN6 GenerationInfo
	GEN7 Generation7
	GEN8 Generation8
	GEN9 Generation9
}

var GENERATIONS = Generations{
	GEN1: GenerationInfo{
		Order: 1,
		Name:  "generation-i",
	},
	GEN2: GenerationInfo{
		Order: 2,
		Name:  "generation-ii",
	},
	GEN3: GenerationInfo{
		Order: 3,
		Name:  "generation-iii",
	},
	GEN4: GenerationInfo{
		Order: 4,
		Name:  "generation-iv",
	},
	GEN5: GenerationInfo{
		Order: 5,
		Name:  "generation-v",
	},
	GEN6: GenerationInfo{
		Order: 6,
		Name:  "generation-vi",
	},
	GEN7: Generation7{
		GEN7: GenerationInfo{
			Order: 7,
			Name:  "generation-vii",
		},
		LGPE: GenerationInfo{
			Order: 7,
			Name:  "lets-go-pikachu-lets-go-eevee",
		},
	},
	GEN8: Generation8{
		SS: GenerationInfo{
			Order: 8,
			Name:  "sword-shield",
		},
		LA: GenerationInfo{
			Order: 8,
			Name:  "legends-arceus",
		},
		BDSP: GenerationInfo{
			Order: 8,
			Name:  "brilliant-diamond-shining-pearl",
		},
	},
	GEN9: Generation9{
		SV: GenerationInfo{
			Order: 9,
			Name:  "scarlet-violet",
		},
		LZA: GenerationInfo{
			Order: 9,
			Name:  "legends-za",
		},
		CHAMPIONS: GenerationInfo{
			Order: 9,
			Name:  "champions",
		},
	},
}

func (p_generationInfo GenerationInfo) IsVersionGroup() bool {
	return (p_generationInfo.Order > 6 &&
		p_generationInfo.Name != GENERATIONS.GEN7.GEN7.Name)
}

func (p_generationInfo GenerationInfo) HasAbilities() bool {
	return (p_generationInfo.Order > GENERATIONS.GEN2.Order &&
		p_generationInfo.Name != GENERATIONS.GEN7.LGPE.Name &&
		p_generationInfo.Name != GENERATIONS.GEN8.LA.Name &&
		p_generationInfo.Name != GENERATIONS.GEN9.LZA.Name)
}

func (p_generations Generations) getFromName(p_gen string) (GenerationInfo, error) {
	generationsReflect := reflect.ValueOf(p_generations)
	// Parse the key, values of p_generations
	for i := 0; i < generationsReflect.NumField(); i++ {
		field := generationsReflect.Field(i)
		switch field.Type() {
		// If the field is an instance of GenerationInfo...
		case reflect.TypeFor[GenerationInfo]():
			name := field.FieldByName("Name").String()
			if name == p_gen {
				return GenerationInfo{
					Order: int(field.FieldByName("Order").Int()),
					Name:  name,
				}, nil
			}
		default:
			// If we're dealing with Gen 7 and later
			for j := 0; j < field.NumField(); j++ {
				subField := field.Field(j)
				name := subField.FieldByName("Name").String()
				if name == p_gen {
					return GenerationInfo{
						Order: int(subField.FieldByName("Order").Int()),
						Name:  name,
					}, nil
				}

			}
		}
	}
	err := fmt.Errorf("%s not found in GENERATIONS", p_gen)
	return GenerationInfo{}, err
}

// Intended to use for Pokemon past attributes checking
func (p_generationInfo GenerationInfo) IsLessOrEqual(p_gen string) bool {
	genInfo, err := GENERATIONS.getFromName(p_gen)
	if err != nil {
		panic(err)
	}
	return p_generationInfo.Order <= genInfo.Order
}

func (p_generationInfo GenerationInfo) CanReceive(p_gen string) bool {
	genInfo, err := GENERATIONS.getFromName(p_gen)
	if err != nil {
		panic(err)
	}
	switch p_generationInfo {
	case GENERATIONS.GEN1, GENERATIONS.GEN2:
		return genInfo.Order < GENERATIONS.GEN3.Order
	case GENERATIONS.GEN3, GENERATIONS.GEN4, GENERATIONS.GEN5, GENERATIONS.GEN6:
		return genInfo.Order > GENERATIONS.GEN2.Order && genInfo.Order <= p_generationInfo.Order
	case GENERATIONS.GEN7.LGPE:
		return p_generationInfo.Name == genInfo.Name
	default:
		return genInfo.Order <= p_generationInfo.Order
	}
}
