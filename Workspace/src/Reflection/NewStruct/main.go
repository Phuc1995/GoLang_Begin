package main

import (
	"fmt"
	"reflect"
)

func MakeStruct(vals ...interface{}) interface{} {
	var sfs []reflect.StructField
	for k, v := range vals{
		t := reflect.TypeOf(v)
		fmt.Println("TypeOf v: ", t)
		sf := reflect.StructField{
			Name: fmt.Sprintf("F%d",(k+1)),
			Type:t,
		}
		sfs = append(sfs,sf)

	}
	st := reflect.StructOf(sfs)
	fmt.Println("st: ",st)
	so := reflect.New(st)
	fmt.Println("so: ",so)
	fmt.Println("so interface: ",so.Interface())
	return so.Interface()
}

func main()  {
	s := MakeStruct(0,"",[]int{})
	// this returned a pointer to a struct with 3 fields:
	// an int, a string, and a slice of ints
	// but you can’t actually use any of these fields
	// directly in the code; you have to reflect them
	sr := reflect.ValueOf(s)
	fmt.Println("sr ValueOf: ",sr)
	fmt.Println(sr.Elem().Field(0).Interface())
	sr.Elem().Field(0).SetInt(20)
	fmt.Println(sr.Elem().Field(0).Interface())
}