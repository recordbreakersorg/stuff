package stuff

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/recordbreakersorg/ins-stuff/stuff/db"
)

func handleFile(w http.ResponseWriter, r *http.Request) {
	var file db.File
	filename := mux.Vars(r)["fileid"]
	fmt.Println(mux.Vars(r))
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("File ID is required"))
		return
	}
	fileID, err := strconv.ParseInt(filename, 10, 64)
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}
	file, err = dbQ.GetFileById(dbCtx, fileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", file.Mime)
	http.ServeFile(w, r, "files/"+filename)
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	var file db.File

	// Parse multipart form with a reasonable max memory (e.g., 10MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
		return
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File not provided", http.StatusBadRequest)
		return
	}
	defer uploadedFile.Close()

	fileBytes := make([]byte, handler.Size)
	_, err = uploadedFile.Read(fileBytes)
	if err != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		return
	}
	fileParams := db.CreateFileParams{
		Mime:     handler.Header.Get("Content-Type"),
		FileSize: handler.Size,
	}
	file, err = dbQ.CreateFile(dbCtx, fileParams)
	if err != nil {
		http.Error(w, "Could not save file metadata", http.StatusInternalServerError)
		return
	}

	filePath := fmt.Sprintf("files/%d", file.ID)
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	_, err = out.Write(fileBytes)
	if err != nil {
		http.Error(w, "Could not write file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"id":%d}`, file.ID)
}
