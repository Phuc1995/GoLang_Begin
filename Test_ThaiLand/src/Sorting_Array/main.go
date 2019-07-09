package main

import "fmt"

func Solution(elements []int) bool {

	var  n, index, indexMax int
	var result bool
	for i:=0; i <len(elements)-1; i++{
		if elements[i] > elements[i+1] {
			n = i
			break
		}
	}
	if n == 0 {
		result = true
		return result
	}else {
		 min := elements[n]
		for n:=n; n < len(elements)-1; n++ {
			if elements[n+1] <= min{
				min = elements[n+1]
				index = n +1
			}
		}
		for j:=0; j <= n; j++ {
			if elements[j] == elements[n]{
				indexMax = j
				break
			}
		}
		swap(elements,indexMax,index)
		var check int
		for i:=0; i <len(elements)-1; i++{
			if elements[i] > elements[i+1] {
				check = i
			}
		}
		if check == 0{
			result = true
			return result
		}else {
			result = true
			return false
		}
	}

	return result
}

func swap(elements []int, i int, j int) {
	var temp int
	temp = elements[j]
	elements[j] = elements[i]
	elements[i] = temp
}
func main()  {
	elements := [][]int{{1,3,3,2,7}, {1,5,3,3,7},{1,3,5,3,4},{0,1,6,3,3,2,8,9,}}
	for _,value := range elements{
		fmt.Println(Solution(value))
		fmt.Println(value)
		fmt.Println("*********************************")
	}
}
