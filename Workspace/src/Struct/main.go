package main

import "fmt"

type StudentInfo struct {
	class string
	email string
	age int
}

type Student struct {
	id int
	name string
	info StudentInfo
}

func main()  {
	//named
	st1 := Student{
		id: 1,
		name: "phuc",
	}

	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)
	fmt.Println("================================================")
	//anonymous
	var anonymous = struct {
		id int
		name string
	}{
		id: 2,
		name: "phuc2",
	}
	fmt.Println(anonymous)

	//pointer with struct
/*	fmt.Println("================================================")
	pointer := &Student{
		999,
		"Phuc999",
	}
	fmt.Println(pointer)
	fmt.Println((*pointer).name)
	fmt.Println(pointer.id)
*/



	//anonymous fields
	type NoName struct {
		string
		int
	}

	fmt.Println("================================================")
	n := NoName{"nguyen hoang phuc", 777}
	fmt.Println(n)
	fmt.Println("================================================")

	student1 := Student{
		id: 123,
		name: "ABC",
		info:StudentInfo{
			class:"Speakup",
			email:"Hoangphuc735@gmail.com",
			age: 27,
		},}
	fmt.Println(student1)

	//Compare 2 struct
	type struct1 struct {
		id int
		name string
	}

	s1 := struct1{1,"AB"}
	s2 := struct1{1,"AB"}
	if (s1 ==s2) {
		fmt.Println("s1 == s2")
	}else {
		fmt.Println("S1 != S2")
	}
}
