package main

import "fmt"

type location struct {
	x     int
	y     int
	previous *location
	next     *location
}

type node struct {
	head       *location
	tail       *location
	nowPlaying *location
}

func solution(array []string) int {
	var result int
	if checkImpossible(array){
		result = 0
	}else {
		findLocationO(array)
	}
	return result
}

//check cause "o" at location can't go on when begin game
func checkImpossible(array []string) bool {
	var result bool
	result = false
	for i:=0; i<2; i++{
		for j:=0; j<len(array)-1; j++{
			if string(array[i][j]) == "o"{
				result = true
				break
			}
		}
	}

	return result
}

func findLocationO(array []string) (x,y int)  {
	for i:=2; i<len(array); i++{
		for j:=0; j<len(array); j++{
			if string(array[i][j]) == "o"{
				x = i
				y = j
			}
		}
	}
	fmt.Println("x,y ",x,y)
	return x, y
}

func (p *node) addLocation(x, y int) error {
	s := &location{
		x:   x,
		y: y,
	}
	if p.head == nil{
		p.head = s
	}else {
		currentNode := p.tail
		currentNode.next = s
		s.previous = p.tail
	}
	p.tail = s
	return nil
}

func main()  {
	B := []string{"..x...","......","....x.",".x....","..x.x.","...o.."}


	for i:=0; i< len(B); i++{
		for j:=0; j<len(B); j++{
			fmt.Print(string(B[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("********")
	fmt.Println("solution: ",solution(B))

}
