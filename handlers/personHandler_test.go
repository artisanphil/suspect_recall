package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type personResult struct {
	Id int `json:"id"`
}

func TestGetPerson(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/person", GetPerson)

	server := httptest.NewServer(router)
	defer server.Close()

	url := fmt.Sprintf("%s/api/person", server.URL)

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	responseString := string(body)

	var result personResult

	err = json.Unmarshal([]byte(responseString), &result)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	id := result.Id

	if id <= 0 {
		t.Errorf("expected id to be an integer greater than 0, got %v", id)
	}
}
