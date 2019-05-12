package main

import "fmt"

func main() {
	a4 := []int{1:1, 3:2}
	fmt.Println(a4)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
