package main

import (
	"fmt"
	"sync"
)

var (
	counter int64
	mutex sync.Mutex
)
func main()  {
	fmt.Println("a")
	for i:=0; i<100; i++ {
		go func() {
			for i:=0;i<10000 ;i++  {
			counter++
			mutex.Unlock()
		}
		}()
	}

	var c string
	fmt.Scanln(&c)
	fmt.Println(counter)
}