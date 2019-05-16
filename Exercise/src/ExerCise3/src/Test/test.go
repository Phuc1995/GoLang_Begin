package main

import "fmt"

type Person struct {
	age int
	weight int
	firtName string
	next *Person
}

func main()  {
	phuc := &Person{10, 50, "Phuc", nil}
	personList := phuc

	Kha := &Person{10, 50, "Kha", nil}
	Long := &Person{10, 50, "Long", nil}
	Hai := &Person{10, 50, "Hai", nil}

	personList = addNode(Kha, personList);
	personList = addNode(Long, personList)
	personList = addNode(Hai, personList)

	fmt.Println(personList)
	personList = reverseRecurrsive(personList)
	//showList(personList)
	fmt.Println(personList)

}

func showList(personlist *Person)  {
	for p:=personlist; p!=nil; p=p.next {
		fmt.Println(p)
	}
}

func addNode(newPerson, personList *Person) *Person {
	if personList == nil{
		return personList
	}
	for p := personList; p != nil; p=p.next{
		if(p.next==nil){
			p.next = newPerson
			return personList
		}
	}
	return personList
}

func reverseRecurrsive(personList *Person) *Person {
	if (personList == nil){
		return personList
	}
	p := personList

	if p.next == nil{
		return p
	}else {
		newHead := reverseRecurrsive(p.next)
		p.next.next = p
		p.next = nil
		return newHead
	}

}