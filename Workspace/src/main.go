package main

import (
	"fmt"
	"unicode"
)

func IsLetter(s string) int {
	count := 0
	for _, r := range s {
		if !unicode.IsLetter(r) {
			count ++
		}
	}
	return count
}

func IsNumber(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			count ++
		}
	}
	return count
}

func main()  {
	myString := "aZ$6"
	isNumber := IsNumber(myString)
	fmt.Println(isNumber)
	isString := IsLetter(myString)
	fmt.Println(isString)
}
