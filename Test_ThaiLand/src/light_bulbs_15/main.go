package main

import "fmt"

func bulbs(array [5]int) int {
	count := 0
	var arrayExit [5]int
	for i, value := range array{
		switch value {
		case 1:
			arrayExit[i] = value
			if isExist(value, array){
				count ++
			}
		case 2:
			 //j := i-1
			arrayExit[i] = value
			if isExist(value,array) && isExist(1,arrayExit){
				count ++
			}
		case 3:
			arrayExit[i] = value
			if isExist(value,array) && isExist(1,arrayExit) && isExist(2,arrayExit){
				count ++
			}
		case 4:
			arrayExit[i] = value
			if isExist(value,array) && isExist(1,arrayExit) && isExist(2,arrayExit) && isExist(3,arrayExit){
				count ++
			}
		case 5:
			arrayExit[i] = value
			if isExist(value,array) && isExist(1,arrayExit) && isExist(2,arrayExit) && isExist(3,arrayExit)&& isExist(4,arrayExit){
				count ++
			}
		}
	}
	return count
}

func isExist(n int, array [5]int) bool {
	var check bool
	check = false
	for _,v := range array{
		if(n ==v){
			check = true
			return check
		}
	}
	return check
}

func main()  {
	array := [5]int{2,1,3,5,4}
	fmt.Println(bulbs(array))

	array2 := [5]int{2,3,4,1,5}
	fmt.Println(bulbs(array2))

	array3 := [5]int{1,3,4,2,5}
	fmt.Println(bulbs(array3))

}
