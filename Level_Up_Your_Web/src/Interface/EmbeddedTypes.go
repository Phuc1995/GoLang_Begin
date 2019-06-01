package main

import "fmt"

type User struct {
	Name string
}

func (u User) IsAdmin() bool  {
	return false
}
func (u User) DisplayName() string  {
	return u.Name
}

type Admin struct {
	User
}
func (a Admin) IsAdmin() bool  {
	return true
}
func (a Admin) DisplayName() string  {
	return "[Amin]" + a.User.Name
}

func main()  {
	u := User{"Normal user"}
	fmt.Println(u.Name)
	fmt.Println(u.DisplayName())
	fmt.Println(u.IsAdmin())

	a := Admin{User{"Admin"}}
	fmt.Println(a.Name)
	fmt.Println(a.IsAdmin())
	fmt.Println(a.DisplayName())

}

