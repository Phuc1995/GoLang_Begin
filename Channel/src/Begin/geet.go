package main

import "fmt"

func greet(c chan string)  {
	fmt.Println("Hello: "+ <-c )
}

func main()  {
	fmt.Println("Begin-----")
	c := make(chan string)
	//f := make(chan string)
	//c <- "phuc"
	go greet(c)

	fmt.Println("end")

}

