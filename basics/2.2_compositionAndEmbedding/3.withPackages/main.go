package main

import (
	"fmt"
	"log"
	"os"
)

type Job struct {
	Command string
	Logger  *log.Logger
}

type JobA struct {
	Command string
	*log.Logger
}

func main() {

	fmt.Println("### Composition 1 ###")

	//struct Job defined a logger type
	//we initaized it and now Logger is avaialbe under job.Logger
	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Logger.Print("test")

	//in Job we ommited the Logger in the struct and now thanks to
	//composition all log.Logger methods are avaialble directly
	jobA := &JobA{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	jobA.Print("test")

}
