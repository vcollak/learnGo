package main

import (
	"fmt"
)

func main() {

	fmt.Println("### Range 1 ###")

	//iterate using index and value
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	powA := make([]int, 10)
	//no value. just index
	for i := range powA {
		powA[i] = 1 << uint(i)
	}

	//no index. just value
	for _, value := range powA {
		fmt.Printf("%d\n", value)
	}

	//break. exit the loop
	powB := make([]int, 10)
	for i := range powB {
		powB[i] = 1 << uint(i)
		if powB[i] > 16 {
			break
		}
	}
	fmt.Println(powB)
	// [1 2 4 8 16 0 0 0 0 0]

	//continue. skip and go back to the top of the loop
	powC := make([]int, 10)
	for i := range powC {
		if i%2 == 0 {
			continue
		}
		powC[i] = 1 << uint(i)
	}
	fmt.Println(powC)
	// [0 2 0 8 0 32 0 128 0 512]

	//range on maps
	cities := map[string]int{

		"New York":    33242342,
		"Los Angeles": 2443534,
	}

	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)

	}

}
