package main

import "fmt"

type Location struct {
	x int
	y int
}

func (l *Location) ToString() string  {
	return fmt.Sprintf("x: %d, y: %d",l.x, l.y)
}
