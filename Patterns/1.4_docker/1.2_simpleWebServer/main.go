package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello home")

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	fmt.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
