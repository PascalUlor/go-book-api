package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/me/bookdbconfig"

	// bookdbconfig "../config/book"

	// "../models"
	"example.com/me/models"

	"github.com/gorilla/mux"
)

// Controller app controller
type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks ... Get all books
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}

		bookStore := bookdbconfig.BookDbConfig{}

		books = bookStore.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

// GetBook ... Get all books
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])

		logFatal(err)

		bookStore := bookdbconfig.BookDbConfig{}

		book = bookStore.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

// AddBook ... Get all books
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		json.NewDecoder(r.Body).Decode(&book)

		bookStore := bookdbconfig.BookDbConfig{}

		bookID = bookStore.AddBook(db, book)

		json.NewEncoder(w).Encode(bookID)
	}
}

// UpdateBook ... Get all books
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		bookStore := bookdbconfig.BookDbConfig{}

		rowsUpdated := bookStore.AddBook(db, book)

		json.NewEncoder(w).Encode(rowsUpdated)

	}
}

// RemoveBook ... Get all books
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		// result, err := db.Exec("delete from books where id=$1", params["id"])
		// logFatal(err)

		// rowsDeleted, err := result.RowsAffected()
		// logFatal(err)
		id, err := strconv.Atoi(params["id"])

		logFatal(err)

		bookStore := bookdbconfig.BookDbConfig{}

		rowsDeleted := bookStore.RemoveBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}

}
