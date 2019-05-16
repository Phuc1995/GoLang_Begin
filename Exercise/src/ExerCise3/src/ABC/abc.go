package main

import (
	"fmt"
)

type node struct {
	value interface{}
	//prev  *node
	next *node
}
type LinkedList struct {
	head *node
}

func (List *LinkedList) Append(value interface{}) {
	new_node := &node{value: value, next: nil}
	headtemp := List.head
	if List.head == nil {
		List.head = new_node
	} else {
		for headtemp.next != nil {
			headtemp = headtemp.next
		}
		headtemp.next = new_node
	}
}
func (List *LinkedList) IndexOf(value interface{}) int {
	headtemp := List.head
	var index int = 0
	if List.head == nil {
		return -1
	} else {
		for headtemp.value != value {
			if headtemp.next == nil {
				return -2
			} else {
				index++
				headtemp = headtemp.next
			}
		}
		return index
	}
}
func (List *LinkedList) Get(index int) interface{} {
	headtemp := List.head
	var listindex int = 0
	if List.head == nil {
		return -1
	} else {
		for listindex != index {
			if headtemp.next == nil {
				return -2
			} else {
				listindex++
				headtemp = headtemp.next
			}
		}
		return headtemp.value
	}
}
func (List *LinkedList) RemoveAt(index int) {
	previous := List.head
	current := List.head
	var indexlist int = 0

	if List.head == nil {
		fmt.Println("Empty List")
	} else {
		indexlist = 1
		for indexlist != index {
			if current.next == nil {
				fmt.Printf("Can not remove at ", index)
			} else {
				indexlist++
				previous = current
				current = current.next
			}
		}
		previous.next = current.next
	}
}
func (List *LinkedList) Size() int {
	headtemp := List.head
	var size int = 0
	if List.head == nil {
		return size
	} else {
		for headtemp.next != nil {
			size++
			headtemp = headtemp.next
		}
		return size + 1
	}
}
func (l *LinkedList) Show() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.value)
		list = list.next
	}
	fmt.Println()
}
func main() {
	a := LinkedList{}

	a.Append("truong")
	a.Append("huu")
	a.Append("thang")
	a.Append("linkedlist")
	a.Append("learning")
	a.Append("golang")
	a.Append("DC")
	a.Append(9)
	a.Show()
	fmt.Println("size:", a.Size())
	fmt.Println("size:", a.IndexOf("huu"))
	fmt.Println("size:", a.IndexOf("oanh"))
	fmt.Println("size:", a.Get(7))
	a.RemoveAt(6)
	a.Show()

}
