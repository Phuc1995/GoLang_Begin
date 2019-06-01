package main

import (
	"fmt"
	"sync"
	"time"
)

var start time.Time

func init()  {
	start = time.Now()
}

func service(wg *sync.WaitGroup, instance int)  {
		time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
		fmt.Println("Time: ", time.Since(start))
	wg.Done()
}

func main()  {
	fmt.Println("main() started")
	var wg sync.WaitGroup // create waitgroup(empty struct)

	for i := 1; i<=3; i++{
		wg.Add(1)
		go service(&wg, i)
	}

	wg.Wait()//blocks here
	fmt.Println("main() stopped")
}
