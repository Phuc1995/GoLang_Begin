package main

import "fmt"

func main()  {
	//declare array
	var  myArray [4] int
	fmt.Println(myArray)

	//enter value for array
	myArray[0] =1234
	fmt.Println(myArray)

	//declare 1 array have value
	arrays := [4] int {1,2,3}
	fmt.Println(arrays)

	//size array
	fmt.Println(len(arrays))

	//declare needn't to set value
	arraay3 := [...] string{"a","c", "b"}
	fmt.Println((arraay3))
	fmt.Println(len(arraay3))

	//array is value type not is reference type

	//look array
	fmt.Println("-------------------------------------")
	for i:=0; i<len(arraay3); i++{
		fmt.Println(arrays[i])
	}

	for index, value := range arraay3{
		fmt.Printf("i=%d value=%s ", index,value)
		fmt.Println()
	}

	//blank indentifier
	for _, value := range arraay3{
		fmt.Printf(" value=%s ",value)
		fmt.Println()
	}

	// mang 2 chieu
	matrix := [4][2] int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println(matrix)

	for i:=0; i<4; i++{
		for j:=0; j<2; j++{
			fmt.Print(matrix[i][j], "  ")

		}
		fmt.Println()
	}
}
