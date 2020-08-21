package main

import (
	"fmt"
	"sync"
	"time"
)

//enables go routines to so we know when they are all done
var wg sync.WaitGroup

//simulated reading file
func printFile(fileName string, c chan string) {

	//after completion call done
	defer wg.Done()

	content := "simulated file content from " + fileName
	time.Sleep(time.Millisecond * 500)

	//send data to the channel
	c <- content
}

func main() {

	//make a channel
	//make sure you add the channel buffer. will have a panic without it
	c := make(chan string, 3)

	//list of fake files to read
	fileNames := []string{"name1", "name2", "name3"}

	//iterate over the list and call printFile using a goroutine
	for i := range fileNames {
		wg.Add(1)
		go printFile(fileNames[i], c)
	}

	//wait until all channels are done
	wg.Wait()

	//close the channel
	close(c)

	for elem := range c {
		fmt.Println(elem)
	}

}
