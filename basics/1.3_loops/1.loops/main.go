package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("-- classic for loop --")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-- classic for with index starting with 1 --")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-- basically a while loop --")
	i := 1
	for i < 10 {
		i++
		fmt.Println(i)
	}

	fmt.Println("-- basically a while true loop with break --")
	iii := 0
	for {
		iii++
		fmt.Println(iii)

		if iii > 100 {
			break
		}
	}

	fmt.Println("-- basically a while true that loops forever (will force exit after 20) --")

	counter := 0
	for true {
		fmt.Println(time.Now().Format(time.RFC850))
		time.Sleep(100 * time.Millisecond)

		//let's just exit after 100
		counter++ 
		if counter > 20 {
			break
		}
	}

}
