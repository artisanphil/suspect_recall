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

	session, _ := store.Get(r, "session")
	var inquired []int
	if sessionInquired, ok := session.Values["inquired"].([]int); ok {
		inquired = sessionInquired
	} else {
		inquired = []int{}
	}

	if len(inquired) >= count {
		inquired = []int{} //instead of deleting session
	}

	randomNum := getUniqueRandomNumber(count, inquired)

	inquired = append(inquired, randomNum)
	session.Values["inquired"] = inquired
	session.Values["correctAttributes"] = []string{}
	session.Values["wrongAttributes"] = []string{}

	sessionError := session.Save(r, w)
	if sessionError != nil {
		fmt.Printf("Failed to save session: %v\n", sessionError)
	}

	noMore := false
	if val, ok := session.Values["inquired"].([]int); ok && len(val) == count {
		noMore = true
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"id": randomNum, "noMore": noMore}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func getUniqueRandomNumber(count int, inquired []int) int {
	for {
		randomNum := rand.Intn(count) + 1
		if !contains(inquired, randomNum) {
			return randomNum
		}
	}
}

// Helper function to check if a slice contains a number
func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}
