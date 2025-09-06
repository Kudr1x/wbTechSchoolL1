package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Printf("Worker %d get and finish task: %d\n", id, j)
		time.Sleep(time.Second) //Типо работает
	}
}

func main() {
	numWorkers := flag.Int("workers", 3, "Number of concurrent workers")
	flag.Parse()

	jobs := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= *numWorkers; i++ {
		workerID := i
		wg.Go(func() {
			worker(workerID, jobs)
		})
	}

	go func() {
		jobID := 1
		for {
			jobs <- jobID
			jobID++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}
