package main

import (
	"fmt"
	"time"
)

//simulated reading file
func printFile(fileName string, c chan string) {
	content := "simulated file content from " + fileName
	time.Sleep(time.Millisecond * 500)

	//send data to the channel
	c <- content
}

func main() {

	//make a channel
	c := make(chan string)

	//list of fake files to read
	fileNames := []string{"name1", "name2", "name3"}

	//iterate over the list and call printFIle using a goroutine
	for i := range fileNames {
		go printFile(fileNames[i], c)
	}

	//get data from the channel
	v1 := <-c
	v2 := <-c
	v3 := <-c

	//print the channel values
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

}
