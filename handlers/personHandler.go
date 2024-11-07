package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	dir := "frontend/public/persons"
	files, err := os.ReadDir(dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var pngFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".png" {
			pngFiles = append(pngFiles, file.Name())
		}
	}

	count := len(pngFiles)
	if count == 0 {
		http.Error(w, "no .png files found", http.StatusInternalServerError)
		return
	}

	rand.Seed(time.Now().UnixNano())

	randomNum := rand.Intn(count) + 1

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"id": randomNum}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
