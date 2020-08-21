package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	fmt.Print("Printing form:")
	fmt.Println(r.Form)

	fmt.Print("Printing path:")
	fmt.Println("path", r.URL.Path)

	fmt.Print("Printing scheme:")
	fmt.Println("scheme", r.URL.Scheme)

	fmt.Print("Printing url_long form:")
	fmt.Println(r.Form["url_long"])

	fmt.Println("Printing form key value pairs:")
	for k, v := range r.Form {
		fmt.Println("     key:", k)
		fmt.Println("     val", strings.Join(v, ""))

	}

	fmt.Fprintf(w, "Hello Bobby")

}

func main() {

	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)

	}

}
