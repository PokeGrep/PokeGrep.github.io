package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type ResourceDir string

const (
	DirAbility                 ResourceDir = "ability"
	DirBerry                   ResourceDir = "berry"
	DirBerryFirmness           ResourceDir = "berry-firmness"
	DirBerryFlavor             ResourceDir = "berry-flavor"
	DirCharacteristic          ResourceDir = "characteristic"
	DirContestEffect           ResourceDir = "contest-effect"
	DirContestType             ResourceDir = "contest-type"
	DirEggGroup                ResourceDir = "egg-group"
	DirEncounterCondition      ResourceDir = "encounter-condition"
	DirEncounterConditionValue ResourceDir = "encounter-condition-value"
	DirEncounterMethod         ResourceDir = "encounter-method"
	DirEvolutionChain          ResourceDir = "evolution-chain"
	DirEvolutionTrigger        ResourceDir = "evolution-trigger"
	DirGender                  ResourceDir = "gender"
	DirGeneration              ResourceDir = "generation"
	DirGrowthRate              ResourceDir = "growth-rate"
	DirItem                    ResourceDir = "item"
	DirItemAttribute           ResourceDir = "item-attribute"
	DirItemCategory            ResourceDir = "item-category"
	DirItemFlingEffect         ResourceDir = "item-fling-effect"
	DirItemPocket              ResourceDir = "item-pocket"
	DirLanguage                ResourceDir = "language"
	DirLocation                ResourceDir = "location"
	DirLocationArea            ResourceDir = "location-area"
	DirMachine                 ResourceDir = "machine"
	DirMove                    ResourceDir = "move"
	DirMoveAilment             ResourceDir = "move-ailment"
	DirMoveBattleStyle         ResourceDir = "move-battle-style"
	DirMoveCategory            ResourceDir = "move-category"
	DirMoveDamageClass         ResourceDir = "move-damage-class"
	DirMoveLearnMethod         ResourceDir = "move-learn-method"
	DirMoveTarget              ResourceDir = "move-target"
	DirNature                  ResourceDir = "nature"
	DirPalParkArea             ResourceDir = "pal-park-area"
	DirPokeathlonStat          ResourceDir = "pokeathlon-stat"
	DirPokedex                 ResourceDir = "pokedex"
	DirPokemon                 ResourceDir = "pokemon"
	DirPokemonColor            ResourceDir = "pokemon-color"
	DirPokemonForm             ResourceDir = "pokemon-form"
	DirPokemonHabitat          ResourceDir = "pokemon-habitat"
	DirPokemonShape            ResourceDir = "pokemon-shape"
	DirPokemonSpecies          ResourceDir = "pokemon-species"
	DirRegion                  ResourceDir = "region"
	DirStat                    ResourceDir = "stat"
	DirSuperContestEffect      ResourceDir = "super-contest-effect"
	DirType                    ResourceDir = "type"
	DirVersion                 ResourceDir = "version"
	DirVersionGroup            ResourceDir = "version-group"
)

var (
	dataDir = "data"
	cache   = make(map[string]any)
)

// Fetches resource by id as long as p_resourceType, id are valid
//
// Usage showcase : db.GetById[models.Pokemon](db.DirPokemon, 6)
func GetById[T any](p_resourceType ResourceDir, id int) (*T, error) {
	idStr := strconv.Itoa(id)
	cacheKey := fmt.Sprintf("%s/%s", p_resourceType, idStr)

	// Check if we already loaded this item
	if val, exists := cache[cacheKey]; exists {
		return val.(*T), nil
	}

	// If not, load the JSON file
	filePath := filepath.Join(dataDir, string(p_resourceType), idStr, idStr+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("resource %s (ID: %d) not found : %w", p_resourceType, id, err)
	}

	var item T
	if err := json.Unmarshal(data, &item); err != nil {
		return nil, fmt.Errorf("error while reading the JSON : %w", err)
	}

	// Cache by ID
	cache[cacheKey] = &item

	// Also cache by Name if the field exists
	val := reflect.ValueOf(item)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		nameField := val.FieldByName("Name")
		if nameField.IsValid() && nameField.Kind() == reflect.String {
			nameCacheKey := fmt.Sprintf("%s/name/%s", p_resourceType, nameField.String())
			cache[nameCacheKey] = &item
		}
	}

	return &item, nil
}

// Fetches the resource list as long as p_resourceType is valid
//
// Usage showcase : db.GetAll[models.Pokemon](db.DirPokemon)
func GetAll[T any](p_resourceType ResourceDir) ([]*T, error) {
	cacheKey := fmt.Sprintf("%s/all", p_resourceType)

	// Check if the list is already cached
	if val, exists := cache[cacheKey]; exists {
		return val.([]*T), nil
	}

	pattern := filepath.Join(dataDir, string(p_resourceType), "*", "*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	var list []*T
	for _, file := range files {
		// Extract the ID from the file name (eg: "*/6.json" -> id = 6)
		idStr := strings.TrimSuffix(filepath.Base(file), ".json")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}

		// Get the item from cache, or put it in cache
		item, err := GetById[T](p_resourceType, id)
		if err == nil {
			list = append(list, item)
		}
	}

	cache[cacheKey] = list
	return list, nil
}

// Fetches resource by name as long as p_resourceType is valid
//
// Usage showcase : db.GetByName[models.Pokemon](db.DirPokemon, "charizard")
func GetByName[T any](p_resourceType ResourceDir, p_name string) (*T, error) {
	cacheKey := fmt.Sprintf("%s/name/%s", p_resourceType, p_name)

	// Return item if cache[p_resourceType/name/p_name] found (should be
	// the case if the element has already been loaded by id)
	if val, exists := cache[cacheKey]; exists {
		return val.(*T), nil
	}

	// Load all elements (which will populate the name cache for everyone)
	_, err := GetAll[T](p_resourceType)
	if err != nil {
		return nil, err
	}

	// Try lookup again from populated cache
	if val, exists := cache[cacheKey]; exists {
		return val.(*T), nil
	}

	return nil, fmt.Errorf("resource %s (Name: %s) not found", p_resourceType, p_name)
}
