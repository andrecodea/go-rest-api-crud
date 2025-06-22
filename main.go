// Indicates that this is an executable
package main

// Imports needed packages
import (
	"encoding/json" // Allows structs to be converted into JSON objects
	"log"           // Logging and debugging
	"math/rand"     // Random ID generation
	"net/http"      // Go's standard package for HTTP server implementation
	"strconv"       // Converts numbers into strings

	"github.com/gorilla/mux" // Router for HTTP routes
)

// Book struct (Model)
type Book struct {
	// Mapping id, isbn and title to JSON
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`

	// Book author is referenced by Author struct
	Author *Author `json:"author"` // Pointer to a struct (allows Author to be nil if no author is defined for a book)
}

// Author struct
type Author struct {
	// Mapping first name and last name to JSON
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice (dynamic array) book struct --> Global variable that serves as a mock DB
var books []Book

// GET all books
// w.Heder().Set("Content-Type", "application/json") --> defines the response's header "Content-Type" as JSON
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Defines content as JSON
	w.Header().Set("Content-Type", "application/json")

	// Sends the books slice as JSON in the response
	// Creates a new JSON encoder thath writes directly into the HTTP response,
	// encodes the books slice to JSON format, and sends it as the response's body
	json.NewEncoder(w).Encode(books)

}

// GET a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	// Defines content as JSON
	w.Header().Set("Content-Type", "application/json")

	// Extracts URL params
	params := mux.Vars(r)

	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] { // Compares the ID in the URL with book's IDs
			json.NewEncoder(w).Encode(item) // Returns corresponding book based in it's ID
			return                          // Leaves function to avoid unecessary processing
		}
	}
	json.NewEncoder(w).Encode(&Book{}) // Returns empty Book struct if none are found

}

// CREATE a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	// Defines content as JSON
	w.Header().Set("Content-Type", "application/json")

	var book Book

	// Decodes request's body's data to the Book struct
	_ = json.NewDecoder(r.Body).Decode(&book)

	// Generates random ID and converts it into a string - not safe lol
	book.ID = strconv.Itoa(rand.Intn(10000000))

	// Appends a new book to global var books (mock DB)
	books = append(books, book)

	// Returns created book
	json.NewEncoder(w).Encode(book)

}

// UPDATE a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	// Defines content as JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract ID param from URL
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {

			// Removes the current book through slicing
			books = append(books[:index], books[index+1:]...) // Creates a new slice that contains all elements before and after index, effectivley removing the element at index
			var book Book

			// Decodes new data for a book
			_ = json.NewDecoder(r.Body).Decode(&book)

			// Mantains same ID --> crucial
			book.ID = params["id"]

			// Appends updated book to global var books
			books = append(books, book)

			json.NewEncoder(w).Encode(book)
			return
		}
	}
	// Returns all books in case ID is not found
	json.NewEncoder(w).Encode(books)

}

// DELETE a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Defines content as JSON
	w.Header().Set("Content-Type", "application/json")

	// Extracts ID from URL
	params := mux.Vars(r)

	for index, item := range books {

		// Compares the given ID
		if item.ID == params["id"] {

			// Removes book from global var books (mock DB)
			books = append(books[:index], books[index+1:]...)

			// Breaks loop if book is removed, no need to iterate further
			break
		}
	}
	// Returns the new books list
	json.NewEncoder(w).Encode(books)

}

// Defining main function
func main() {
	// Init router instance
	router := mux.NewRouter() // will be responsible for mapping HTTP requests to the right CRUD functions

	// Mock data - TODO - Implement DB
	books = append(books, Book{ID: "1", Isbn: "332568", Title: "The Fall of the Soviet Union", Author: &Author{Firstname: "Andrei", Lastname: "Volkov"}})
	books = append(books, Book{ID: "2", Isbn: "458156", Title: "Economy for Dummies", Author: &Author{Firstname: "Steve", Lastname: "Kilbow"}})

	// Route Handlers / Endpoints (URIs) / CRUD --> associate endpoints to CRUD HTTP verbs
	// HandleFund --> Defines a specific endpoint
	// Methods() --> Defines a HTTP verb to an endpoint
	router.HandleFunc("/api/books", getBooks).Methods("GET")          // --> GET request: gets the books
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")      // --> GET request: gets a book individualy
	router.HandleFunc("/api/books", createBook).Methods("POST")       // --> POST request: creates a book
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")   // --> PUT request: updates a book permanently
	router.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE") // --> DELETE request: deletes a book

	// Starts web server
	// ListenAndServe --> Tells Go to start "listening" for HTTP requests at port 8000.
	// The "router" is passed as a main handler for all received requests.
	// log.Fatal() --> If an error is encountered, it will log it and exit the program.
	log.Fatal(http.ListenAndServe(":8000", router))
}
