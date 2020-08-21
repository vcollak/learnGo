package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	First string
	Last  string
	Email string
}

func main() {
	session, err := mgo.Dial("localhost:27100")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//get a session
	c := session.DB("gmailContacts").C("contacts")

	//find the contact
	result := Person{}
	err = c.Find(bson.M{"email": "Bobby@Peters.net"}).One(&result)
	if err != nil {

		//insert the contact
		err = c.Insert(&Person{"Bobbyimir", "Peters", "Bobby@Peters.net"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Email inserted")

	} else {
		fmt.Println("Unable to insert. Contact already exists")
		fmt.Println("Email:", result.Email)
	}

}
