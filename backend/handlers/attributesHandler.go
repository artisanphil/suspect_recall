package handlers

import (
	"bufio"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

func GetItems(c *gin.Context) {
	has, err := ReadLines("./private/persons/1-has.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read 1-has.txt"})
		return
	}

	hasNot, err := ReadLines("./private/persons/1-hasnot.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read 1-hasnot.txt"})
		return
	}

	items := append(has, hasNot...)
	ShuffleLines(items)

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func CheckAttribute(c *gin.Context) {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	var req struct {
		ClickedAttribute string   `json:"clickedAttribute" binding:"required"`
		Attributes       []string `json:"attributes" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hasFilePath := "./private/persons/" + id + "-has.txt"
	has, err := ReadLines(hasFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read " + id + "-has.txt"})
		return
	}

	for _, line := range has {
		if line == req.ClickedAttribute {
			c.JSON(http.StatusOK, gin.H{"exists": true})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"exists": false})
}
