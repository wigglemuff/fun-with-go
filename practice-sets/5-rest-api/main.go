package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var books []Book

// curl http://localhost:8000/api/v1/books
func getBooks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// curl http://localhost:8000/api/v1/books -X POST -H 'Content-Type:application/json' -d '{"name": "Eat That Frog", "author": "Brian Tracy"}'
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	book.ID = len(books) + 1
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// curl http://localhost:8000/api/v1/books/1
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idstr := strings.TrimPrefix(r.URL.Path, "/api/v1/books/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	// another way to extract the id:
	// re := regexp.MustCompile(`^/api/v1/books/(\d).*`)
	// id := re.FindStringSubmatch("r.URL.Path")[1]

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "404 book not found", http.StatusNotFound)
}

// curl http://localhost:8000/api/v1/books/1 -X PUT -H 'Content-Type:application/json' -d '{"name":"Bad Habits", "author":"Bad Author"}'
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idstr := strings.TrimPrefix(r.URL.Path, "/api/v1/books/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
	}

	for idx, book := range books {
		// find the book
		if book.ID == id {

			// decode the data
			var newbook Book
			json.NewDecoder(r.Body).Decode(&newbook)

			// write to the book, while keeping the same ID
			books[idx] = newbook
			books[idx].ID = book.ID

			// return success
			json.NewEncoder(w).Encode(book)

			// exit
			return
		}
	}

	http.Error(w, "404 book not found", http.StatusNotFound)
}

// curl http://loaclhost:8000/api/v1/books/1 -X DELETE
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idstr := strings.TrimPrefix(r.URL.Path, "/api/v1/books/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
	}

	for idx, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			books = append(books[:idx], books[idx+1:]...)
			return
		}
	}

	http.Error(w, "404 book not found", http.StatusNotFound)
}

func main() {

	// Sample data
	books = append(books, Book{ID: 1, Name: "Atomic Habits", Author: "James Clear"})
	books = append(books, Book{ID: 2, Name: "Hyperfocus", Author: "Chris Bailey"})

	// Handlers
	http.HandleFunc("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBooks(w)
		case http.MethodPost:
			createBook(w, r)
		}
	})
	http.HandleFunc("/api/v1/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBook(w, r)
		case http.MethodPut:
			updateBook(w, r)
		case http.MethodDelete:
			deleteBook(w, r)
		}
	})

	// Serve traffic
	http.ListenAndServe(":8000", nil)
}
