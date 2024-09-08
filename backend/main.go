package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"suspectRecall/handlers"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})

	mux.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)
		mimeType := mime.TypeByExtension(ext)

		if mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		}

		http.ServeFile(w, r, "."+r.URL.Path)
	})

	mux.HandleFunc("/api/person/attributes", handlers.GetItems)
	//mux.HandleFunc("/api/person/:id/check-attribute", handlers.CheckAttribute)
	mux.HandleFunc("/api/person/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/check-attribute") {
			handlers.CheckAttribute(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
