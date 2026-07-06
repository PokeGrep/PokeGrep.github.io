package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	dataDir = "data"
	cache   = make(map[string]any)
)

// Fetches resource by id as long as p_resourceType, id are valid
func GetById[T any](p_resourceType string, id int) (*T, error) {
	idStr := strconv.Itoa(id)
	cacheKey := fmt.Sprintf("%s/%s", p_resourceType, idStr)

	// Check if we already loaded this item
	if val, exists := cache[cacheKey]; exists {
		return val.(*T), nil
	}

	// If not, load the JSON file
	filePath := filepath.Join(dataDir, p_resourceType, idStr, idStr+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("resource %s (ID: %d) not found : %w", p_resourceType, id, err)
	}

	var item T
	if err := json.Unmarshal(data, &item); err != nil {
		return nil, fmt.Errorf("error while reading the JSON : %w", err)
	}

	cache[cacheKey] = &item
	return &item, nil
}

// Fetches the resource list as long as p_resourceType is valid
func GetAll[T any](p_resourceType string) ([]*T, error) {
	cacheKey := fmt.Sprintf("%s/all", p_resourceType)

	// Check if the list is already cached
	if val, exists := cache[cacheKey]; exists {
		return val.([]*T), nil
	}

	pattern := filepath.Join(dataDir, p_resourceType, "*", "*.json")
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
