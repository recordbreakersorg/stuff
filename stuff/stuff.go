// Package stuff manages files and the stuff.
package stuff

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() {
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
