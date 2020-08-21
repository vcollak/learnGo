package main

import "fmt"

/*
In Go, only constants are immutable. However because arguments are passed by value,
a function receiving an value argument and mutating it, won’t mutate the original value.

As you can see the total amount of songs on the me variable’s value wasn’t changed.
To mutate the passed value, we need to pass it by reference, using a pointer.

*/

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}

func newReleaseA(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {

	fmt.Println("### Mutability ###")
	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
	fmt.Println()

	/*
		As you can see the total amount of songs on the me variable’s value wasn’t changed.
		To mutate the passed value, we need to pass it by reference, using a pointer.
		The only change between the two versions is that newRelease takes a pointer to
		an Artist value and when I intialize our me variable, I used the & symbol to get a pointer to the value.


	*/
	meA := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", meA.Name, newReleaseA(meA))
	fmt.Printf("%s has a total of %d songs", meA.Name, meA.Songs)
	fmt.Println()

}
