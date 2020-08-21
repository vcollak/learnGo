package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	//"time"
)

type IndexPageContent struct {
	Name     string
	DateTime time.Time
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	page := IndexPageContent{Name: "Hello Bob", DateTime: time.Now()}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("Error parsing the HTML", err)
	}

	t.Execute(w, page)

}

func main() {

	portNumber := "8080"
	portString := ":" + string(portNumber)

	fmt.Println("Web App Started on localhost", portString)

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(portString, nil)
}
