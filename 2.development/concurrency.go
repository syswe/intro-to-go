package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second) // Simulating some work
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start the workers
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect the results
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}
