package main

import "fmt"

func Solution(A [5]int) int {
	var fix_sum, cur_sum, res_temp, val_return int
	for i:=0;i<len(A);i++{
		fix_sum += A[i]
		if A[i] < 0 {
			val_return += ^A[i]+1
		} else {
			val_return += A[i]
		}
	}
	for i:=0;i<len(A)-1;i++{
		cur_sum += A[i]
		res_temp = cur_sum - (fix_sum - cur_sum)
		if res_temp < 0 {
			res_temp = ^res_temp+1
		}
		if res_temp < val_return {
			val_return = res_temp
		}
	}

	return val_return
}

func main() {
	var  A [5]int
	A[0] = 3
	A[1] = 1
	A[2] = 2
	A[3] = 4
	A[4] = 3
	Solution(A)
	fmt.Println(Solution(A))
}