package builder

import (
	"os"
	"pokegrep/internal/localization"
	"pokegrep/templates"

	"github.com/otiai10/copy"
)

// IsDev indicates if the builder is running in development/preview mode
var IsDev bool

func Build(isDev bool) bool {
	IsDev = isDev

	if err := os.MkdirAll("dist", 0755); err != nil {
		panic(err)
	}
	err := copy.Copy("static", "dist")
	if err != nil {
		panic(err)
	}

	// Render the redirection index page at the root of dist using templ
	indexComponent := templates.Index()
	return renderToFile(indexComponent, "dist/index.html") &&
		buildLanguage(localization.LOCALES.ENGLISH) &&
		buildLanguage(localization.LOCALES.FRENCH)
}
