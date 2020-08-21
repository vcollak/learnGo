package main

import (
	"fmt"

	"github.com/google/uuid"
	"log"
)

func main() {

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Unable to get UUID")
	}

	fmt.Println(uuid)

}
