package builder

import (
	"context"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
)

// renderToFile renders a templ.Component directly to the specified file path.
func renderToFile(component templ.Component, outputPath string) bool {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = component.Render(context.Background(), file)
	if err != nil {
		panic(err)
	}

	return true
}

// shouldRebuild returns true if the target file needs to be regenerated.
// If IsDev is false, it always returns true (force full build in CI/production).
func shouldRebuild(outputPath string, sourceFiles ...string) bool {
	if !IsDev {
		return true
	}

	destInfo, err := os.Stat(outputPath)
	if err != nil {
		return true // File does not exist, build it
	}

	destModTime := destInfo.ModTime()

	for _, src := range sourceFiles {
		srcInfo, err := os.Stat(src)
		if err != nil {
			continue
		}
		if srcInfo.ModTime().After(destModTime) {
			return true // Source is newer, rebuild
		}
	}

	return false
}
