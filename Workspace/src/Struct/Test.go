package main

import "fmt"

type Student struct {
	ID int
	Name string
	Age int
}

func main()  {
	student := Student{}
	fmt.Println(student)

	student.Age = 10
	student.ID = 1
	student.Name = "Phuc"
	fmt.Println(student)

	student1 := Student{12, "Phuc1",65}
	fmt.Println(student1)

}