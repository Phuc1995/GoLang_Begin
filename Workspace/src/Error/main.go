package main

import "fmt"

func input()(int, error)  {
	var n int
	fmt.Print("Nhap vao so khong am: ")
	_, err := fmt.Scanf("%d", &n)
	if err == nil && n<0{
		return n, fmt.Errorf("Loi! %d la so am",n)
	}
	return n,err
}

func main() {
	input()

}