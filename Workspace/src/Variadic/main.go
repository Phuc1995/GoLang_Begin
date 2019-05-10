package main

import "fmt"

func addItem(item int, list ...int)  {
	list = append(list, item)
	fmt.Println(list)
}

func change(list ...int)  {
	list[0] = 999
}

func main()  {
	addItem(1, 100,200,300,400)

	var list  = []int {1,2,3,4}
	addItem(100, list...)

	change(list...)
	fmt.Println(list)
}
