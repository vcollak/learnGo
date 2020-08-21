package main

import (
	"fmt"
)

//Dog type and method
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("woof")
}

//cat type and method
type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("meow")
}

//goldendoodle embeds the speaker interface
type Animal struct {
	Speaker
}

//type that also includes Speaker interface
type SpeakerPrefixer struct {
	Speaker
}

//when we call the Speak() function here it's called, but
//then we call the Speaker's Speak()
func (sp SpeakerPrefixer) Speak() {
	fmt.Print("Prefix:")
	sp.Speaker.Speak()
}

//types can implement the Speaker interface
type Speaker interface {
	Speak()
}

func main() {

	//animal embeds the dog. because it implements the
	//Speaker interface, we can call the Speak method
	//in this case the Dog.Speak() is called
	//this also shows interface chaining:
	//Animal embeds Speaker (Dog, Cat, SpeakerPrefix)
	//Speaker Prefix also embeds Speaker (Dog, Cat)
	dog := Animal{SpeakerPrefixer{Dog{}}}
	dog.Speak()

}
