package main

import "fmt"

func Chao1()  {
	fmt.Printf("hello")
}
func Chao(name string)  {
	fmt.Printf("hello ", name )
}

func greeting(name string) string {
	result := fmt.Sprintln("hello ", name)
	return result
}

//Multiple return values
 func recInfo(w, h int)(int, int){
	return w, h
}

//Named return values
func recInfo1(w, h int)(width int, height int, isSquare bool){
	u := 100
	isSquare = u == h
	return w, h, isSquare
}

func main()  {
	result := greeting("phuc")
	fmt.Println(result)

	w,h, isSquare := recInfo1(100,100)
	fmt.Println("width", w)
	fmt.Println("hight", h)
	fmt.Println("hight", isSquare)

}