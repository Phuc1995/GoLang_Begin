package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// worker than make squares
func sqrWorker1(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		fmt.Println("[sqrWorker1]Number goroutine: ", runtime.NumGoroutine())
		results <- num * num
	}

	// done with worker
	wg.Done()
}

func main() {
	fmt.Println("[main] main() started")

	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker1(&wg, tasks, results, i)

	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)
	//fmt.Println("Number goroutine: ", runtime.NumGoroutine());
	// wait until all workers done their job
	wg.Wait()
	//fmt.Println("Number goroutine: ", runtime.NumGoroutine())
	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // non-blocking because buffer is non-empty
		fmt.Println("[main] Result", i, ":", result)
	}
	fmt.Println("Number goroutine: ", runtime.NumGoroutine())
	fmt.Println("[main] main() stopped")
}
