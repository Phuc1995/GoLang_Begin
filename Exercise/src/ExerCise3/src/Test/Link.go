package main
/*
import "fmt"

type Node struct {
	prev *Node
	next *Node
	name interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next : L.head,
		name: key,
	}
	if L.head != nil{
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil  {
		l = l.next
	}
	L.tail = l
}

func (l *List ) Display( ){
	list := l.head
	for list != nil{
		fmt.Printf("%v <-", list.name)
		list = list.next
	}
	fmt.Println()
}
func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.name)
		list = list.next
	}
	fmt.Println()
}
func (l *List)Reverse()  {
	curr := l.head
	var prev *Node
	l.tail = l.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

func main()  {
	link := List{}
	link.Insert(1)
	link.Insert(2)
	link.Insert(3)
	link.Insert(4)
	link.Insert(5)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.name)
	fmt.Printf("Tail: %v\n", link.tail.name)
	link.Display()
	link.Reverse()

}
*/