package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/recordbreakersorg/stuff/stuff"
)

func main() {
	println("Starting INS Backend...")

	versionCmd := flag.Bool("version", false, "Show version and exit")
	flag.Parse()

	if *versionCmd {
		println("Version 0.1.0")
		return
	}

	stuff.Setup()
	println("INS Stuff setup successfully.")
	defer stuff.Unset()
	routes := stuff.CreateRoutes()
	http.Handle("/", routes)
	println("Starting HTTP server on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
