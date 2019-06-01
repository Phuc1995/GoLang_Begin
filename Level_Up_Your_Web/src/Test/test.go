package main

import "fmt"

func main()  {
	mySlice := []string{"Cat", "Dong", "Trung"}
	for animail := range mySlice{
		fmt.Println(animail, "-", )
	}

	myMap := map[int]string{
		1 : "Phuc",
		2 : "Hao",
		3 : "Trung",
	}
	result, exist := myMap[4]
	fmt.Println(result)
	fmt.Println(exist)

}
