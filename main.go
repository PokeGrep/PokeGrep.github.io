package main

import (
	"os"
	"pokegrep/internal/builder"
)

func main() {
	builder.Build()
	os.Exit(0)
}
