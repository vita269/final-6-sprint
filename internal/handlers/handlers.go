package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		log.Println("ParseMultipartForm error:", err)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
		log.Println("FormFile error:", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		log.Println("ReadAll error:", err)
		return
	}

	converted, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "Data conversion error", http.StatusInternalServerError)
		log.Println("Convert error:", err)
		return
	}

	filename := time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)
	f, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		log.Println("Create file error", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(converted)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		log.Println("WriteString error:", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(converted))
	if err != nil {
		log.Println("Write error:", err)
	}
}
