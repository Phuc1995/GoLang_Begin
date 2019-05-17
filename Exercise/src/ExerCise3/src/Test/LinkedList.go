package main

import (
	"fmt"
)

type Note struct {
	value string
	next *Note
}

type LinkedList struct {
	length int
	start *Note
	end *Note
}

func (f *LinkedList) Append(value *Note) {
	if(f.length == 0){
		f.start = value
		f.end = value
	}else {
		lastPost := f.end
		lastPost.next = value
		f.end = value
	}
	f.length++
}



func (l *LinkedList ) Display( ){
	list := l.start
	for list != nil{
		fmt.Printf("%v -", list.value)
		list = list.next
	}
	fmt.Println()
}

func (l *LinkedList) Size() int {
	return l.length
}

func (l *LinkedList) Get(index int) interface{} {
	head := l.start
	var listIndex int = 1
	if(l.start == nil){
		return "list emty"
	}else {
		for listIndex != index{
			if head.next == nil{
				return "Not find value"
			}else {
				head = head.next
				listIndex ++
			}
		}
		return head.value
	}
}

func (l *LinkedList) IndexOf(value interface{}) int {
	head := l.start
	var point int = 1
	if (l == nil) {
		return -1
	} else {
		for head.value != value {
			if (head.next == nil) {
				return -2
			} else {
				point++
				head = head.next
			}
		}
		return point
	}
}
func (f *LinkedList) RemoveAt(index int)  {
	previous := f.start
	after := f.start

	var listIndex int = 0
	if(f.start == nil){
		fmt.Println("list emty")
	}else {
		for listIndex !=index  {
			if after.next == nil{
				fmt.Printf("Out size value")
			}else {
				previous = after
				after = after.next
				listIndex++
			}

		}
		previous.next = after.next
	}
}

func main()  {
	f := LinkedList{}
	note1 := Note{value:"Mua"}
	note2 := Note{value:"Troi"}
	note3 := Note{value:"Ca"}
	note4 := Note{value:"Bau"}
	note5 := Note{value:"Troi"}
	f.Append(&note1)
	f.Append(&note2)
	f.Append(&note3)
	f.Append(&note4)
	f.Append(&note5)

	f.Display()
	println("Size: ",f.Size())

	fmt.Println("Value of Get() at 4: ",f.Get(4))
	fmt.Printf("Inter value need to find IndexOf: ")
	var n string
	fmt.Scan(&n)
	if(f.IndexOf(n) == -1){
		println("Emty list")
	}else if f.IndexOf(n) == -2 {
		println("Not exist value in List")
	}else {
		println("Index of funtion IndexOf: ",f.IndexOf(n))
	}
	fmt.Println("After remove index 4")
	f.RemoveAt(4)

	f.Display()
	}