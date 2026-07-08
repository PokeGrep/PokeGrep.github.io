package builder

import (
	"os"
	"pokegrep/internal/localization"

	"github.com/otiai10/copy"
)

func Build() bool {
	if err := os.MkdirAll("dist", 0755); err != nil {
		panic(err)
	}
	err := copy.Copy("static", "dist")
	if err != nil {
		panic(err)
	}
	tmpl := []string{
		"templates/index.html",
	}
	return buildTemplate(tmpl, nil, "dist") &&
		buildLanguage(localization.LOCALES.ENGLISH) &&
		buildLanguage(localization.LOCALES.FRENCH)
}
