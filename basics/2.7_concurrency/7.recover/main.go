package main

import (
	"fmt"
	"sync"
	"time"
)

//define the wait group
var wg sync.WaitGroup

func cleanup() {

	//clean up is the last thing the go-routine will do
	//so we need to get done with the waitgroup
	defer wg.Done()

	//if there was a panic, recovery will not be nil and we can do something		
	if r := recover(); r != nil {
		fmt.Println("Recovery after panic:", r)
	}

}

func add(num int) {

	//run clean up at the end to see if we need to clean up somehow
	defer cleanup()

	for i := 0; i < num; i++ {

		time.Sleep(time.Microsecond * 2)
		num = i + num
		fmt.Println("Result:", num)

		//simulating panic. this will cause the recover() function to execute
		//otherwise the go routines will terminate because of the panic we called
		if i == 3 {
			panic("Panicked for some reason")
		}
	}

}

func main() {

	wg.Add(1)
	go add(7)

	wg.Add(1)
	go add(12)

	wg.Add(1)
	go add(5)

	wg.Add(1)
	go add(15)

	//wait until the Wait group counter is zero
	wg.Wait()

}
