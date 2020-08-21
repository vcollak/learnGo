package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Bobby")

}

func main() {

	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)

	}

}
