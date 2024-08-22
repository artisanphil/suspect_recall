package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type attributeData struct {
	Items []string `json:"items"`
}

type attributeExists struct {
	Exists bool `json:"exists"`
}

type PostAttribute struct {
	ClickedAttribute string   `json:"clickedAttribute"`
	Attributes       []string `json:"attributes"`
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
	resp, err := http.Get("http://localhost:3000/api/person/attributes")
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

	resp, err := http.Post(
		"http://localhost:3000/api/person/1/check-attribute",
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

	var item attributeExists

	err = json.Unmarshal([]byte(responseString), &item)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, item.Exists, "Expected result to be true")
}
