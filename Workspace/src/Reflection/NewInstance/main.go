package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main()  {
	greeting := "hello"
	f := Foo{10, "Salutations"}

	gVal := reflect.ValueOf(greeting)
	// not a pointer so all we can do is read it
	fmt.Println(gVal.Interface())

	gpVal := reflect.ValueOf(&greeting)
	// it’s a pointer, so we can change it, and it changes the underlying variable
	gpVal.Elem().SetString("Bye")
	fmt.Println(greeting)

	fType := reflect.TypeOf(f)
	fVal := reflect.New(fType)
	fVal.Elem().Field(0).SetInt(20)
	fVal.Elem().Field(1).SetString("Greetings")
	f2 := fVal.Elem().Interface().(Foo)
	fmt.Printf("%+v, %d, %s ",f2, f2.A, f2.B)
}
