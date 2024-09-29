package main

import (
	"log"
	"net/http"
	"suspectRecall/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	})

	r.HandleFunc("/api/person/attributes", handlers.GetItems)

	r.HandleFunc("/api/person/{id}/check-attribute", handlers.CheckAttribute).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/build/index.html")
	})

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/build/static"))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/build/"))))

	log.Println("Listening on :8080...")

	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}
