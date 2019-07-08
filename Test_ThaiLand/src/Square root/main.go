package main

import (
	"fmt"
	"math"
)

func Solution(A, B int) int {
	var array []float64
	var   count int
	for i:= A; i <= B; i++{
		if square_check(i){
			n := int(math.Sqrt(float64(i)))
			count ++

		}
	}
	fmt.Println(array)
	return count
}
func square_check(a int) bool {
	var int_root int = int(math.Sqrt(float64(a)))
	return (int_root * int_root) == a
}

func main()  {
	var A, B int
	A = 10
	B = 20

	fmt.Println(Solution(A, B))
}
