package helpers

import "pokegrep/internal/localization"

type Names []struct {
	Language struct {
		Name string
		URL  string
	}
	Name string
}

func (p_names Names) Get(p_lang string) string {
	for _, item := range p_names {
		if item.Language.Name == p_lang {
			return item.Name
		}
	}

	// English fallback
	for _, item := range p_names {
		if item.Language.Name == localization.LOCALES.ENGLISH {
			return item.Name
		}
	}
	return "item.Language.Name UNDEFINED"
}
