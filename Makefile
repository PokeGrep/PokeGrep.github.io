.PHONY: all dev dev-run generate fmt run serve

# Default target : builds templates, formats code, then launch SSG
all: generate fmt run

# Generate Go files from *.templ
generate:
	go run github.com/a-h/templ/cmd/templ generate

# Formats Go code
fmt:
	go fmt ./...

# Build SSG
run:
	go run .

# Build SSG then serve the files located at ./dist
serve:
	go run . --serve

# Called by templ at each change
dev-run: fmt serve

# Dev Mode : Watch go files, formats code,
# re-serve.
dev:
	go run github.com/a-h/templ/cmd/templ generate --watch --cmd="make dev-run"
