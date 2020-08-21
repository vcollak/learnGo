package main

import (
	"fmt"
	"sync"
	"time"
)

//function that can be turned into a go routine
func Counter() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Counter ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	//call the function as a go routine. inside we call the
	//function and can also handle when the routine is completed.
	//this approach allows us not to have special go routine code
	//like the waitgroups in the function
	go func() {
		Counter()
		wg.Done()
	}()

	wg.Wait()

}
