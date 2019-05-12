package main

import (
	"fmt"
	"runtime"
)

func main()  {
	runtime.GOMAXPROCS(10)
	numberP := runtime.NumCPU()
	fmt.Println(numberP)

	numberP1 := runtime.GOMAXPROCS(0)
	fmt.Println(numberP1)
}
