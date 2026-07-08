package main

import (
	"fmt"
	"net/http"
	"os"
	"pokegrep/internal/builder"
)

func main() {
	fmt.Println("Building...")
	builder.Build()
	fs := http.FileServer(http.Dir("dist"))
	fmt.Println("Build Done! Server running at http://localhost:8081")
	err := http.ListenAndServe(":8081", fs)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
