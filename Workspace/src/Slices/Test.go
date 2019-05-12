package main

import "fmt"

func main()  {
	var a1 []int
	var a2 []int = []int{1,2,3,4}
	a3 := make([]int,3)
	a4 := []int{1:1, 3:2}
	a5 := []int {6,6,6,8,8}
	fmt.Println(a1)  // []
	fmt.Println(a2)  // [1 2 3]
	fmt.Println(a3)  // [0 0 0]
	fmt.Println(a4)  // [0 1 0 2]
	fmt.Println("a5:",a5)  // [0 0 0]
	a1 = append(a1, 1)
	fmt.Println("append a1",a1)
	a3 = append(a3, a2...)
	fmt.Println("append a3",a3)
	fmt.Println(copy(a4, a5), a4)
}
