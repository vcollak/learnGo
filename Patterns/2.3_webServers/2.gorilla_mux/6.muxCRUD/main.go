package main

import (
	"encoding/json"
	"strconv"

	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux" //http router
)

//Book struct - holds information about a single book

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct - info about the author. embedded in the Book Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//init books variable as a slice of book structs
var books []Book

//getBooks - get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//encodes and writes to response writer
	json.NewEncoder(w).Encode(books)

}

//getBook - get one book
func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//get the POST params
	params := mux.Vars(r)

	//look for the right book in our books struct
	//once found retun that book
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	//return empty
	json.NewEncoder(w).Encode(&Book{})

}

//createBook - create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book Book

	//decode the body and write to &book via a reference
	//@todo - handle an error
	_ = json.NewDecoder(r.Body).Decode(&book)

	//create a random ID @todo - change to some counter or let DB handle in the future
	book.ID = strconv.Itoa(rand.Intn(10000000)) //mock ID - not safe

	//append to books
	books = append(books, book)

	//return the book we added - could return success / error instead?
	json.NewEncoder(w).Encode(book)

}

//updateBook - updates a book
func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//get the POST params
	params := mux.Vars(r)

	//look for the right book in books
	for index, item := range books {
		if item.ID == params["id"] {

			//found the books. take all books before and after and append them
			//into books. This will actually create new arrays
			//... basically splits into individual books and appends
			books = append(books[:index], books[index+1:]...)

			var book Book

			//decode the book that's being updated
			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = strconv.Itoa(rand.Intn(10000000)) //mock ID - not safe

			//append to books
			books = append(books, book)

			//return the book we updated - could return success / fail
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	//did not find a book - return all books we have.
	//@todo - return success / fail
	json.NewEncoder(w).Encode(books)

}

//deleteBook - delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//get the POST params
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {

			//append before this book and after this book
			//it actually creates a new array
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)

}

func main() {

	//init the http router
	r := mux.NewRouter()

	//mock data @todo - implement database
	books = append(books, Book{
		ID:    "1",
		Isbn:  "33423412",
		Title: "Atlas Shrugged",
		Author: &Author{
			FirstName: "Ayn",
			LastName:  "Rand",
		},
	})

	books = append(books, Book{
		ID:    "2",
		Isbn:  "w42543535356463564",
		Title: "Foundation",
		Author: &Author{
			FirstName: "Isac",
			LastName:  "Asimov",
		},
	})

	//api endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
