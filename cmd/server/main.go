package main

import (
	"bookexchange/api"
	"bookexchange/store"
	"bookexchange/workflow"
	"log"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("BOOK_DB")
	if path == "" {
		path = "books.db"
	}
	s, e := store.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	w := workflow.New(s)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", api.New(w).Handler()))
}
