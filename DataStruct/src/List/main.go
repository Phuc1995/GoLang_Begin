package main

import (
	"container/list"
	"fmt"
)

func main()  {
	var intList list.List

	intList.PushBack(11)
	intList.PushBack(21)
	intList.PushBack(31)
	fmt.Println("intList.Front: ",intList.Front().Value)
	fmt.Println("intList.Front: ",intList.Front().Next().Value)

	for element := intList.Front(); element != nil; element=element.Next() {
		fmt.Println(element.Value)
	}
}
