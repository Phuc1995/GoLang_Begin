package main

import "fmt"

func giveMePear(fruit *string)  {
	*fruit = "banano"
}

func main(){
	var u *string

	fruit := "chuoi"
	giveMePear(u)
	fmt.Println(u)
}