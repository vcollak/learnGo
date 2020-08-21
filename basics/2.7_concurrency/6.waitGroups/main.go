package main

/*
Provides a way to use go routines, but allows them to sync
so we don't exit the script prematurely or too late
*/

import (
	"fmt"
	"sync"
	"time"
)

//define the wait group
var wg sync.WaitGroup

func add(num1 int, num2 int) {

	//this specifies when the function completes (normally or panic)
	//defer makes it so the wg.Done command is only called at the end
	defer wg.Done()

	num := num1 + num2
	fmt.Println("Result:", num)
	time.Sleep(time.Second * 2)
}

func main() {

	//have to add the go routine to waitgroup group
	//Add adds delta, which may be negative, to the WaitGroup counter.
	//If the counter becomes zero, all goroutines blocked on Wait are released.
	//If the counter goes negative, Add panics.
	//https://golang.org/pkg/sync/#WaitGroup.Add
	wg.Add(1)
	go add(11, 1)

	wg.Add(1)
	go add(11, 2)

	wg.Add(1)
	go add(11, 3)

	wg.Add(1)
	go add(11, 4)

	//wait until the Wait group counter is zero
	wg.Wait()

}
