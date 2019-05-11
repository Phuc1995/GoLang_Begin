package main

import (
	"fmt"
)

func main()  {
	a := 100

	var  pointer  *int
	pointer = &a
	fmt.Println(a)
	*pointer = 999 // a;=999
	fmt.Println(a)
	fmt.Printf("%T", pointer)
	applyPointer(pointer)
	fmt.Println("After call function applyPointer",a)
	//zero value
	fmt.Println("--------------")
	var pointer2 *int
	pointer2 = new(int)
	fmt.Println(pointer2)

	fmt.Println("=======================================")

	array := [3]int{1,2,3}
	var pointer3 *[3]int
	pointer3 = &array
	fmt.Println(pointer3)



}

func applyPointer(pointer *int)  {
	*pointer = 777
}