package builder

import (
	"html/template"
	"os"
)

func buildTemplate(p_templateSource string, p_data any, p_templateOutput string) bool {
	tmpl := template.Must(
		template.ParseFiles(p_templateSource),
	)

	// Create parent directories recursively if they don't exist
	if err := os.MkdirAll(p_templateOutput, 0755); err != nil {
		panic(err)
	}

	file, err := os.Create(p_templateOutput + "/index.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, p_data)
	if err != nil {
		panic(err)
	}

	return true
}
