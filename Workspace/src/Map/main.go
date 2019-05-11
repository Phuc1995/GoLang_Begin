package main

import "fmt"

func main()  {
	var myMap = make(map[string]int)
	fmt.Println(myMap)

	var myMap1 map[string]int
	fmt.Println(myMap1)

	//nil is a predeclared identifier representing the zero value for a
	// pointer, channel, func, interface, map, or slice type.
	if myMap1 == nil{
		fmt.Println("myMap1 = nil")
	}

	//declare with initialize value
	myMap2 := map[string]int{
		"key1" :1,
		"key2" :2,
		"key3" :3,
	}
	fmt.Println(myMap2)
	//add 1 element inti map
	myMap2["key4"] = 4
	myMap2["key5"] = 5
	fmt.Println(myMap2)
	//delete 1 element in map
	delete(myMap2, "key1")
	fmt.Println(myMap2)
	//length of map
	fmt.Println(len(myMap2))

	//Map os reference type
	fmt.Println("----------------------")
	myMap3 := myMap2
	fmt.Println(myMap3)
	myMap3["key5"] = 999
	fmt.Println(myMap2)

	//access 1 element of map

	value, found := myMap2["key2"]
	if found {
		fmt.Println(value)
	}else {
		fmt.Println("Not find value with value 1000")
	}




}