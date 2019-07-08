package main

import "fmt"

func IsRowEqual(A [][]int, rowIndex int) bool {
	sumAbove := 0
	sumBelow := 0

	for i := 0; i < rowIndex; i++ {
		for j := 0; j < len(A[0]); j++ {
			sumAbove += A[i][j]
		}
	}
	for i := rowIndex + 1; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			sumBelow += A[i][j]
		}
	}

	return sumAbove == sumBelow
}
func IsColumnEqual(A [][]int, colIndex int) bool {
	sumLeft := 0
	sumRight := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < colIndex; j++ {
			sumLeft += A[i][j]
		}
		for j := colIndex + 1; j < len(A[0]); j++ {
			sumRight += A[i][j]
		}
	}

	return sumLeft == sumRight
}

func Solution(A [][]int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if IsRowEqual(A, i) && IsColumnEqual(A, j) {
				fmt.Println("solution: ", A[i][j],"-", i, "-",j)

				count++
			}
		}
	}
	return count
}



func main(){
	A := [][]int{
		{2, 7, 5},
		{3, 1, 1},
		{2, 1, -7},
		{0, 2, 1},
		{1, 6, 8},
	}
	fmt.Println(Solution(A))
}
