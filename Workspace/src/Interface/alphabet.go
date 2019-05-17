package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Begin Gorountines")

	go func() {
		//Display lowercase alphabet 3 times
		for count := 0; count <3; count++{
			for char:='a';char<'a'+26;char++  {
			fmt.Printf("%c", char)
				time.Sleep(time.Millisecond * 300)
			}
		}
	}()

	go func() {
		//Display uppercase alphabet 3 times
		for count := 0; count <3; count++{
			for char:='A';char<'A'+26;char++  {
				fmt.Printf("%c", char)
				time.Sleep(time.Millisecond * 300)
			}
		}
	}()

	fmt.Println("End gorountine")
	var input string
	fmt.Scanln(&input)
	fmt.Println("\nKết thúc chương trình")
}
