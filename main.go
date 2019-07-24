package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"fistname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get Single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete book
func deleteBooks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init Mux Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{
		ID:    "1",
		Isbn:  "448743",
		Title: "Book One",
		Author: &Author{ // ampersand because Author has its own Struct to which we reference to
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	books = append(books, Book{
		ID:    "2",
		Isbn:  "448745",
		Title: "Book Two",
		Author: &Author{
			Firstname: "Jane",
			Lastname:  "Doe",
		},
	})

	// Route Handlers / Endponts
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
