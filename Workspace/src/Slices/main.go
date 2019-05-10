package main

import "fmt"

func main()  {
	//declare slice
	var slice []int
	fmt.Println(slice)

	//declare and crete slice
	var slice1 = []int {1,2,3,4}
	fmt.Println(slice1)

	//create 1 slice to 1 array
	var array  = [4]int {1,2,3,4}
	slice2 := array[1:3] //array[1] -> array[3-1=2] <=> array[2]
	fmt.Println(slice2)

	slice3 := array[:]
	fmt.Println(slice3)

	slice4 := array[2:]
	fmt.Println(slice4)

	slice5 := array[:3]
	fmt.Println(slice5)

	//create 1 slice from 1 other slice
	var slice6 = []int{1,2,3,4,5,6,7,8,9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:]
	fmt.Println(slice8)

	//slice => reference type
	var array1 = [5]int {1,2,3,4,5}
	slice9 := array1[:]
	slice9[0] = 777
	fmt.Println(slice9)
	fmt.Println(array1)

	//length and capacity of slice
	countries := [...]string{"VN", "USA", "CANADA", "CHINA"}
	slice10 := countries[2:3] //CANADA
	fmt.Println(slice10)
	fmt.Println(len(countries))
	//len: number of elements slice
	//cap: number of elements underlying array begin point when slice create

	//make
	fmt.Println("------------------------")
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	//append
	fmt.Println("------------------------")
	var slice13 []int
	slice13 = append(slice13, 100)
	fmt.Println(slice13)

	//coppy
	src := []string {"A", "B", "C", "D"}
	dest := make([]string, 2)
	copy(dest, src)
	fmt.Println(dest)

	//delete item with index=1
	fmt.Println("------------------------")
	src = append(src[:1], src[2:]...)
	fmt.Println(src)
}
