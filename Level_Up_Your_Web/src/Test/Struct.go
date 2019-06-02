package main

import (
	"fmt"
	_ "github.com/russross/blackfriday"
)

type Counter struct{
	Count int
}

func (c Counter) Increment () {
	c.Count++
}

func (c *Counter) IncrementWithPointer() {
	c.Count++
}

func main()  {
	counter := &Counter{}
	fmt.Println(counter.Count)

	counter.Increment()
	fmt.Println(counter.Count)

	counter.IncrementWithPointer()
	fmt.Println(counter.Count)
}