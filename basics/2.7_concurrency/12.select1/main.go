package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "Every 500 ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			ch2 <- "Every 2 sec"
			time.Sleep(time.Second * 2)
		}
	}()

	//below code will not produce the desired effects. (getting ameesage every 500ms and also
	//every 2s)
	//that's because the first statement of getting data from the channel will block
	//as a result, we'll only get alternating 500ms and 2s messages
	/*
		for {
			fmt.Println(<-ch1)
			fmt.Println(<-ch2)

		}*/

	counter := 0
	for {
		//select is what we want. we'll loop continuously and
		//if there is something in the channel we'll receive it
		//this way we get 500ms messages more frequently.
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}

		counter++
		if counter > 25 {
			break
		}
	}
}
