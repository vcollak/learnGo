/*
Implemented http using the gorilla mux library
using REST. Able to distinguish between GET, PUT, POST and DELETE
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index page"))
}

func getUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprintf(w, "Getting user %v\n", vars["id"])

}

func updateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprintf(w, "Updating user %v\n", vars["id"])
}

func addUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprintf(w, "Adding user %v\n", vars["id"])

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprintf(w, "Deleting user %v\n", vars["id"])

}

func main() {

	r := mux.NewRouter()

	//home and only GET
	r.HandleFunc("/", getIndex).Methods("GET")

	//route for /user where we expect some number
	//to be placed in the id variable
	r.HandleFunc("/user/{id:[0-9]+}", getUser).Methods("GET")

	//same as GET above, but updates the user by ID
	r.HandleFunc("/user/{id:[0-9]+}", updateUser).Methods("PUT")

	//creates the user. No need for an ID
	r.HandleFunc("/user", addUser).Methods("POST")

	//deletes the user by ID
	r.HandleFunc("/user/{id:[0-9]+}", deleteUser).Methods("DELETE")

	//listen in port 9000 and server requests
	http.ListenAndServe(":9000", r)

}
