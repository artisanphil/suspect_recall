package main

import (
	"log"
	"net/http"
	"suspectRecall/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	r.HandleFunc("/api/person/attributes", handlers.GetItems)

	r.HandleFunc("/api/person/{id}/check-attribute", handlers.CheckAttribute).Methods("POST")

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", r)
}
