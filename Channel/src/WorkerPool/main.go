package main

import (
	"fmt"
	"runtime"
	"time"
)

func sqrWorker(tasks <-chan int, results chan<- int, instance int )  {

	for num := range tasks{
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] Sending result by worker %v %v\n", instance, instance, num)
		results <- num * num
	}
}

func main()  {
	fmt.Println("[main] main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	//launching 3 worker goroutines
	for i := 1; i<=3; i++{
		go sqrWorker(tasks, results, i)
	}
	fmt.Println(runtime.NumGoroutine())
	//passing 5 tasks
	for i:= 1; i<=5; i++{
		tasks <- i
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	for i := 1; i<=5; i++{
		result := <- results
		fmt.Println("[main] Result", i, ":", result)
		fmt.Println(runtime.NumGoroutine())
	}

	fmt.Println("[main] main() stopped")
}
