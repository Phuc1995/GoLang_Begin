package main

import "fmt"

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

/*func (f *LinkedList) Remove(publishDate int64)  {
	if(f.length==0){
		panic(errors.New("Feed is empty"))
	}
	var previousPost *Note
	currentPost := f.start

	for currentPost.publishDate != publishDate  {
		if currentPost.next == nil {
			panic(errors.New("No such Post found."))
		}
		previousPost = currentPost
		currentPost = currentPost.next
		}
		previousPost.next = currentPost.next

	f.length--
}*/
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
	}