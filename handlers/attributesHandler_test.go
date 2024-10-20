package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type attributeData struct {
	Items []string `json:"items"`
}

type attributeResult struct {
	Exists   bool `json:"exists"`
	Mistakes int  `json:"mistakes"`
	Finished bool `json:"finished"`
}

type PostAttribute struct {
	ClickedAttribute string   `json:"clickedAttribute"`
	Attributes       []string `json:"attributes"`
}

func TestMain(m *testing.M) {
	err := os.Chdir("..")
	if err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}

// Helper function to create a temporary file with content
func createTempFile(t *testing.T, content string) string {
	file, err := ioutil.TempFile("", "testfile")
	assert.NoError(t, err)
	_, err = file.WriteString(content)
	assert.NoError(t, err)
	err = file.Close()
	assert.NoError(t, err)
	return file.Name()
}

func TestReadLines(t *testing.T) {
	fileContent := "line1\nline2\nline3\n"
	filePath := createTempFile(t, fileContent)
	defer os.Remove(filePath)

	lines, err := ReadLines(filePath)
	assert.NoError(t, err)
	assert.Equal(t, []string{"line1", "line2", "line3"}, lines)
}

func TestShuffleLines(t *testing.T) {
	lines := []string{"line1", "line2", "line3", "line4", "line5"}
	originalLines := make([]string, len(lines))
	copy(originalLines, lines)

	rand.Seed(1)
	ShuffleLines(lines)

	assert.NotEqual(t, originalLines, lines, "Lines should be shuffled")
}

func TestGetItems(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/person/attributes", GetItems)

	server := httptest.NewServer(router)
	defer server.Close()

	url := fmt.Sprintf("%s/api/person/attributes", server.URL)

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
	var attributes attributeData

	err = json.Unmarshal([]byte(responseString), &attributes)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	var items = attributes.Items

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, items, "Male")
	assert.Contains(t, items, "Female")
}

func TestCheckAttribute(t *testing.T) {
	itemData := PostAttribute{
		ClickedAttribute: "Male",
		Attributes:       []string{"Full beard", "Obese"},
	}
	jsonData, err := json.Marshal(itemData)

	router := mux.NewRouter()
	router.HandleFunc("/api/person/{id}/check-attribute", CheckAttribute).Methods("POST")

	// Use the router in the test server
	server := httptest.NewServer(router)
	defer server.Close()

	url := fmt.Sprintf("%s/api/person/1/check-attribute", server.URL)
	t.Log(url)

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	responseString := string(body)

	var item attributeResult

	err = json.Unmarshal([]byte(responseString), &item)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, item.Exists, "Expected result to be true")
	assert.Equal(t, item.Mistakes, 2)
	assert.Equal(t, item.Finished, false)
}
