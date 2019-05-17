package main

import "fmt"

func main()  {
	//Unbuered channel
	ch := make(chan int)

	go func() {
		ch <- 100
		fmt.Println("sent")
	}()
	fmt.Println(<-ch)
	fmt.Println("done")
}