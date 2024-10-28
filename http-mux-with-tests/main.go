// clone and run these commands:
// go mod init http-mux-with-tests
// go mod tidy
// go run main.go
// go test -v

// You can test with curl as well:
// Create a book: curl -X POST http://localhost:8000/books -d '{"id":"1", "title":"The Go Programming Language", "author":"Alan A. A. Donovan", "year":"2015"}' -H "Content-Type: application/json"
// Get all books: curl http://localhost:8000/books
// Get a book: curl http://localhost:8000/books/1
// Update a book: curl -X PUT http://localhost:8000/books/1 -d '{"title":"Go in Action", "author":"William Kennedy", "year":"2015"}' -H "Content-Type: application/json"
// Delete a book: curl -X DELETE http://localhost:8000/books/1

package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	books = append(books, Book{ID: "1", Title: "Atomic Habits"})
	books = append(books, Book{ID: "2", Title: "Hyperfocus"})
	router := mux.NewRouter()
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":8000", router)

	// Example without mux:
	// http.HandleFunc("/api/books", getBooks)
	// http.ListenAndServe(":8000", nil)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	book.ID = strconv.Itoa(len(books) + 1)
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book not found"})
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book not found"})
	// or http.NotFound(w, r)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book not found"})
}
