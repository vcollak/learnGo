/*
Implemented http using the gorilla mux library
*/

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SayHelloWorld will response back with Hello World
func SayHelloWorld(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, World!"))
}

func main() {

	r := mux.NewRouter()

	//when user accesses "/" the mus routes to SayHelloWorld function
	r.HandleFunc("/", SayHelloWorld)

	//use the router and listen on port 9000
	http.ListenAndServe(":9000", r)

}
