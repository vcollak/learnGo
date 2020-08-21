/* using mux to get back json - overwriting how we marshal the tags */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

var todos Todos

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex).Methods("GET")
	router.HandleFunc("/todos", TodoShow).Methods("POST")

	date := time.Date(2020, 7, 27, 20, 34, 58, 651387237, time.UTC)

	todos = Todos{
		Todo{Name: "Task #1", Completed: true, Due: date},
		Todo{Name: "Task #1", Completed: true, Due: date},
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Received:", todos)
}
