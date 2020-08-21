package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main() {

	//news
	//var locations Locations

	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}

	bites, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bites))

	resp.Body.Close()

}
