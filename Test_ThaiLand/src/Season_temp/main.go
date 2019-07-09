package main

import (
	"fmt"
	"math"
)

var indexToSeason = map[int]string{
	0: "WINTER",
	1: "SPRING",
	2: "SUMMER",
	3: "AUTUMN",
}

func solution(T []int) string {

	len := len(T)
	n := len/4
	array := []int{}

	winter := T[0:n]
	array = append(array,Temperature(winter))

	spring := T[n:(n*2)]
	array = append(array,Temperature(spring))

	summer := T[(n*2):(n*3)]
	array = append(array,Temperature(summer))

	autumn := T[(n*3):(n*4)]
	array = append(array,Temperature(autumn))

	max := array[0]
	var index int

	for i, value := range array{
		if value >max {
			max = value
			index = i
		}
	}
	season := indexToSeason[index]
	return season
	}

func Temperature(array []int) int {
	max := array[0]
	min := array[0]
	for _, value := range array{
		if value > max {
			max = value
		}
	}
	for _, value := range array{
		if value < min {
			min = value
		}
	}
	n := min - max
	return int(math.Abs(float64(n)))
}

func main() {

	T := []int{2, -3, 3, 1, 10, 8, 2, 5, 13, -5, 3, -18}
	fmt.Println(solution(T))

	T1 := []int{-3, -14, -5, 7, 8, 42, 8, 3}
	fmt.Println(solution(T1))
	/*	T := []int{-9,1,2,12,36,95,12}
		fmt.Println(Temperature(T))*/
}