package handlers

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/gorilla/mux"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ShuffleLines(lines []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"] //getting id from route /api/person/{id}/attributes

	if !ok {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	has, err := ReadLines("./private/persons/" + id + "-has.txt")
	if err != nil {
		http.Error(w, "Unable to read "+id+"-has.txt", http.StatusInternalServerError)
		return
	}

	hasNot, err := ReadLines("./private/persons/" + id + "-hasnot.txt")
	if err != nil {
		http.Error(w, "Unable to read "+id+"-hasnot.txt", http.StatusInternalServerError)
		return
	}

	items := append(has, hasNot...)
	ShuffleLines(items)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"items": items}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func CheckAttribute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"] //getting id from route /api/person/{id}/check-attribute

	if !ok {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	var req struct {
		ClickedAttribute string   `json:"clickedAttribute"`
		Attributes       []string `json:"attributes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	hasFilePath := "./private/persons/" + id + "-has.txt"
	hasAttributes, err := ReadLines(hasFilePath)
	if err != nil {
		http.Error(w, "Unable to read "+id+"-has.txt", http.StatusInternalServerError)
		return
	}

	correct := 0
	mistakes := 0
	for _, attribute := range req.Attributes {
		if slices.Contains(hasAttributes, attribute) {
			correct++
		} else {
			mistakes++
		}
	}

	exists := slices.Contains(hasAttributes, req.ClickedAttribute)

	if exists {
		correct++
	} else {
		mistakes++
	}

	finished := correct == len(hasAttributes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"exists":   exists,
		"mistakes": mistakes,
		"finished": finished,
	})
}
