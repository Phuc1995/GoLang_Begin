package main

import (
	"fmt"
	_ "github.com/russross/blackfriday"
)

type Human struct{
	Name string
	Age int
}

func main()  {
	me := Human{"ma", 29}
	fmt.Println(me)
}
