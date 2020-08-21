package main

import (
	"fmt"
)

func main() {

	//jobs will have a list of numbers to generate fib on
	jobs := make(chan int, 10)

	//results are the fib results
	results := make(chan int, 10)

	//call as many go routine workers are want to
	//they will run as long as there are jobs on teh queue
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	//add jobs to the jobs queue
	for i := 0; i < 10; i++ {
		jobs <- i
	}

	//close jobs. we are done with the channel
	close(jobs)

	//get the results
	for j := 0; j < 10; j++ {
		fmt.Println(<-results)
	}

}

//worker to process jobs. notice that we can only send to jobs
//and only receive from jobs.
//range through the jobs, cal the fib and put results into results channel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

//fibronachi function. if less or equal to 1 just return the numver passed
//once more than one return sum of n - 1 and n - 2
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
