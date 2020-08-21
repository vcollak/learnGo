package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

//player has user struct embedded
type Player struct {
	User
	GameId int
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

func main() {

	fmt.Println("### Composition ###")

	//instantiate the player. Playser has
	//access to User's struct becasue it's embedded
	p := Player{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)

	//option 2 - using a struct literal
	pA := Player{
		User{Id: 42, Name: "Matt", Location: "LA"},
		90404,
	}

	fmt.Printf(
		"Id: %d, Name: %s, Location: %s, Game id: %d\n",
		pA.Id, pA.Name, pA.Location, pA.GameId)
	// Directly set a field define on the Player struct
	pA.Id = 11
	fmt.Printf("%+v", pA)
	fmt.Println(pA.Greetings())

}
