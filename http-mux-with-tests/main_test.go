package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	return router
}

func TestCreateBook(t *testing.T) {
	router := setupRouter()

	book := Book{ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Year: "2015"}
	payload, _ := json.Marshal(book)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdBook Book
	json.Unmarshal(resp.Body.Bytes(), &createdBook)
	assert.Equal(t, book.ID, createdBook.ID)
	assert.Equal(t, book.Title, createdBook.Title)
}

func TestGetBooks(t *testing.T) {
	router := setupRouter()

	// Seed with initial data
	books = []Book{
		{ID: "1", Title: "Book One", Author: "Author One", Year: "2001"},
		{ID: "2", Title: "Book Two", Author: "Author Two", Year: "2002"},
	}

	req, _ := http.NewRequest("GET", "/books", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var retrievedBooks []Book
	json.Unmarshal(resp.Body.Bytes(), &retrievedBooks)
	assert.Equal(t, 2, len(retrievedBooks))
}

func TestGetBook(t *testing.T) {
	router := setupRouter()

	// Seed with initial data
	books = []Book{
		{ID: "1", Title: "Book One", Author: "Author One", Year: "2001"},
	}

	req, _ := http.NewRequest("GET", "/books/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var retrievedBook Book
	json.Unmarshal(resp.Body.Bytes(), &retrievedBook)
	assert.Equal(t, "1", retrievedBook.ID)
}

func TestUpdateBook(t *testing.T) {
	router := setupRouter()

	// Seed with initial data
	books = []Book{
		{ID: "1", Title: "Book One", Author: "Author One", Year: "2001"},
	}

	updatedBook := Book{Title: "Updated Book", Author: "Updated Author", Year: "2021"}
	payload, _ := json.Marshal(updatedBook)

	req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var updated Book
	json.Unmarshal(resp.Body.Bytes(), &updated)
	assert.Equal(t, "Updated Book", updated.Title)
}

func TestDeleteBook(t *testing.T) {
	router := setupRouter()

	// Seed with initial data
	books = []Book{
		{ID: "1", Title: "Book One", Author: "Author One", Year: "2001"},
	}

	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// Check if the book was deleted
	req, _ = http.NewRequest("GET", "/books/1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
