package main

import (
	"fmt"
	"reflect"
)



func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("=====================================")
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind()==reflect.Float64)
	fmt.Println("value:", v.Float())

	fmt.Println("=====================================")
	var x1 uint8 = 'x'
	v2 := reflect.ValueOf(x1)
	fmt.Println("type:", v2.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v2.Kind() == reflect.Uint8) // true.
	//x = uint8(v2.Uint())                                       // v.Uint returns a uint64.

}