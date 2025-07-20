// Package stuff manages files and the stuff.
package stuff

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var StufPath string

func Setup() {
	StufPath = os.Getenv("STUFF_PATH")
	if StufPath == "" {
		log.Fatal("STUFF_PATH environment variable is not set")
		os.Exit(1)
	}
	SetupDB()
}

func Unset() {
	UnsetDB()
}

func CreateRoutes() http.Handler {
	server := mux.NewRouter()
	server.HandleFunc("/file/{fileid}", handleFile)
	server.HandleFunc("/upload", handleFileUpload).Methods("POST")
	return server
}
