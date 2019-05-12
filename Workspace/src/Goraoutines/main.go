package main

import (
	"fmt"
	"sync"
)

func g1()  {
	fmt.Println("g1")
	wg.Done()
}

func g2()  {
	fmt.Println("g2")
	wg.Done()
}

var wg sync.WaitGroup
func main()  {
/*	go g1();
	ng := runtime.NumGoroutine()
	fmt.Println(ng)
	time.Sleep(time.Second)
	fmt.Println("--------------------------------")*/

	//synchronized goroutines
	fmt.Println("Begin")
	wg.Add(2)
	go g1()
	go g2()
	wg.Wait()
	fmt.Println("End")

}