package main

import "fmt"

func main() {
	xs := string("Hello World and Coders")
	fmt.Println("Begin string: ", xs)
	fmt.Print("Reversed string: ")
	for _, v := range string(xs){
		defer fmt.Printf("%c", v)
	}
}