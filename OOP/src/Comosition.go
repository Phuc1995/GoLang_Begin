package main

import (
	"fmt"
	"github.com/amy/tech-talk/composition"
)
type Human struct {
	FirstName   string
	LastName    string
	CanSwim     bool
}
// Amy is embedded with the Human type
// and can thus invoke methods in Human's method set
type Amy struct {
	Human
}
// Alan is also embedded with the Human type
type Alan struct {
	Human
}
func (h *Human) Name() {

	fmt.Printf("Hello! My name is %v %v", h.FirstName, h.LastName)
}
func (h *Human) Swim() {

	if h.CanSwim == true {
		fmt.Println("I can swim!")
	} else {
		fmt.Println("I can not swim.")
	}
}

func main(){
	c := Composi

}