package helpers

import (
	"fmt"
	"pokegrep/internal/localization"
	"reflect"
)

type TranslatableEntry string

const (
	NamesField TranslatableEntry = "Names"
)

// GetTranslatedName extracts the translated name from a structure (PokemonSpecies, Ability, Move, etc.)
// containing a slice field named "Names"
func GetTranslated(p_item any, p_translatableField TranslatableEntry, p_lang string) string {
	val := reflect.ValueOf(p_item)

	// Handle pointer dereferencing
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return "names NOT_A_STRUCT"
	}

	// PokeAPI uses the "Names" field
	field := val.FieldByName(string(p_translatableField))

	if !field.IsValid() || field.Kind() != reflect.Slice {
		return fmt.Sprintf("NO_%s_FIELD", p_translatableField)
	}

	length := field.Len()

	// 1. Search for the target language (e.g. "fr")
	for i := 0; i < length; i++ {
		entry := field.Index(i)
		if entry.Kind() != reflect.Struct {
			continue
		}
		if getLangName(entry) == p_lang {
			return getTextField(entry, "Name")
		}
	}

	// 2. Fallback to English (localization.LOCALES.ENGLISH)
	for i := 0; i < length; i++ {
		entry := field.Index(i)
		if entry.Kind() != reflect.Struct {
			continue
		}
		if getLangName(entry) == localization.LOCALES.ENGLISH {
			return getTextField(entry, "Name")
		}
	}

	// 3. Fallback to the first available element
	if length > 0 {
		return getTextField(field.Index(0), "Name")
	}

	return "UNDEFINED"
}

// Safely extracts the value of Language.Name (e.g. "fr") from a slice entry
func getLangName(entry reflect.Value) string {
	langField := entry.FieldByName("Language")
	if langField.IsValid() && langField.Kind() == reflect.Struct {
		nameField := langField.FieldByName("Name")
		if nameField.IsValid() && nameField.Kind() == reflect.String {
			return nameField.String()
		}
	}
	return ""
}

// Safely extracts a string field (e.g. "Name") from a slice entry
func getTextField(entry reflect.Value, fieldName string) string {
	field := entry.FieldByName(fieldName)
	if field.IsValid() && field.Kind() == reflect.String {
		return field.String()
	}
	return "UNDEFINED"
}
