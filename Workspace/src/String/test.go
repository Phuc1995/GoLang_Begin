package main

import "fmt"

func sum(vals... int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}


func main()  {
	i := sum(1,6,10,200,6,7)
	fmt.Println("Sum: ",i)
	fmt.Printf("%T ",i )
}
