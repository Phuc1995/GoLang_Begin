package main

import (
	"fmt"
	"reflect"
)

func main()  {
	//declaring these vars, so I can make a reflect.Type
	intSlice := make([]int, 0)
	mapStringInt := make(map[string]int)

	//here are the reflect.Types
	sliceType := reflect.TypeOf(intSlice)
	fmt.Println(sliceType)
	mapType := reflect.TypeOf(mapStringInt)
	fmt.Println(mapType)

	//and here are there new values that we are making
	intSliceReflect := reflect.MakeSlice(sliceType,0,0)
	mapReflect := reflect.MakeMap(mapType)

	//and here we are using them
	v := 10
	rv := reflect.ValueOf(v)
	fmt.Println("v: ",rv)
	intSliceReflect = reflect.Append(intSliceReflect,rv)
	intSlice2 := intSliceReflect.Interface().([]int)
	fmt.Println(intSlice2)

	k := "hello"
	rk := reflect.ValueOf(k)
	fmt.Println("k: ",rk)
	mapReflect.SetMapIndex(rk,rv)
	mapStringInt2 := mapReflect.Interface()
	fmt.Println(mapStringInt2)
}
