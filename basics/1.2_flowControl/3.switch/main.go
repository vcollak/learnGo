package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("### Switch ###")
	myNum := 8
	switch myNum {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 4 - 1:
		fmt.Println("3")
	case 5, 6, 7:
		fmt.Println("5,6,7")
	default:
		fmt.Println("other")
	}

	//fallthrogh falls through the next statement no matter what the case is.
	//it's basically as if there was not "break" that's used in other languages
	fmt.Println("### Fall through ###")
	n := 4
	switch n {
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 0:
		//the number is 4 so above shows.
		//this (0) will also execute since it's set to fall through
		//it shops here since there is no further fallthough statement
		fmt.Println("is <= 0")
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		fallthrough
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
		fallthrough
	default:
		fmt.Println("Try again!")
	}

	fmt.Println("### Fall through - another ###")
	m := 1
	switch m {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("<= 1")
		fallthrough
	case 2:
		fmt.Println("<= 2")
		fallthrough
	case 3:
		fmt.Println("<= 3")

		//if time is divisible by 2 then print and break (exits)
		//otherwise, it continues to fall through
		if time.Now().Unix()%2 == 0 {
			fmt.Println("Lorem ipsum dolor sit amet")
			break
		}
		fallthrough
	case 4:
		fmt.Println("<= 4")
		fallthrough
	case 5:
		fmt.Println("<= 5")
	}
	fmt.Println("------------------")

	//checking strings
	fmt.Println("### Checking strings ###")
	name := "Bobby"
	switch name {
	case "Bobby":
		fmt.Println("Bobby")
	case "Liz":
		fmt.Println("Liz")
	}
}
