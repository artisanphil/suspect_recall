package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

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

	session, _ := store.Get(r, "session")

	if inquired, ok := session.Values["inquired"].([]int); ok {
		inquired = append(inquired, randomNum)
		session.Values["inquired"] = inquired
	} else {
		session.Values["inquired"] = []int{randomNum}
	}

	sessionError := session.Save(r, w)
	if sessionError != nil {
		fmt.Printf("Failed to save session: %v\n", sessionError)
	}

	fmt.Println(session.Values["inquired"])
	noMore := false
	if val, ok := session.Values["inquired"].([]int); ok && len(val) == count {
		noMore = true
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"id": randomNum, "noMore": noMore}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
