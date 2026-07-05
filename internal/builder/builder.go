package builder

import (
	"pokegrep/internal/localization"
)

func Build() bool {
	return buildLanguage(localization.LOCALES.ENGLISH) &&
		buildLanguage(localization.LOCALES.FRENCH)
}
