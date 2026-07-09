package main

import (
	"flag"
	"fmt"
	"net/http"
	"pokegrep/internal/builder"
)

func main() {
	// Env variables parsing
	serveFlag := flag.Bool("serve", false, "Démarrer un serveur HTTP de prévisualisation local")
	portFlag := flag.String("port", "8080", "Port d'écoute du serveur de prévisualisation")
	flag.Parse()

	fmt.Println("Building...")
	builder.Build()
	fmt.Println("Build complete!")

	if *serveFlag {
		fs := http.FileServer(http.Dir("dist"))
		fmt.Println("Server running at http://localhost:8080")
		err := http.ListenAndServe(":"+*portFlag, fs)
		if err != nil {
			panic(err)
		}
	}
}
