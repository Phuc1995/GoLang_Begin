package main

import "fmt"

func Swap(x int, y int)(int, int)  {
	return y, x
}

func main()  {
	x := 10;
	y := 15
	fmt.Println(x,"-",y)
	x,y = Swap(x,y)
	fmt.Println(x,"-",y)
}