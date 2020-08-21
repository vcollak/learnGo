/*

Channels
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.


*/

package main

import (
	"fmt"
	"time"
)

//loop and sleep in between loops
func say(s string, goRoutineCounter int) {
	for i := 0; i < 1000; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Routine: %d Loop: %d Msg: %s\n", goRoutineCounter, i, s)
	}
}

func main() {

	fmt.Println("### Concurrency ###")
	for i := 0; i < 10; i++ {
		go say("world", i)
	}

	//the main program will wait for 10 seconds and exists after that
	//not guaranteed that routines will finish before then
	time.Sleep(10 * time.Second)
	fmt.Println("Done")

}
