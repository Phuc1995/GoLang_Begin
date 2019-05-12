package main

import "fmt"

func main() {
	name := "Phuc"
	fmt.Println("Begin string: ", name)
	fmt.Println("Reversed string: ")
	for _, v := range [] rune(name){
		defer fmt.Printf("%c", v)
	}
}