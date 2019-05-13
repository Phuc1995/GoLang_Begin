package main

import "fmt"

/*for { // Như vòng lặp do while
<Khối lệnh>
}
for <Biểu thức luận lý> { // Như vòng lặp while
<Khối lệnh>
}
for <Mô tả trước>; <Biểu thức luận lý>; <Mô tả sau> {
<Khối lệnh>
}
for <Chỉ số>, <Giá trị> := range <Biến kiểu dữ liệu có phần tử> {
<Khối lệnh>
}*/

func fact1(n uint) uint  {
	var factn uint =1
	var i uint = 1

	for{
		factn = factn * i
		i++
		if i>n{
			break
		}

	}
	return factn
}

func fact2(n uint) uint  {
	var i, factn  uint= 1,1
	for i<=n{
		factn = factn * i
		i++
	}
	return factn
}

func fact3(n uint) uint  {
	var factn  uint= 1

	for i:= uint(1); i<=n; i++{
		factn = factn * i

	}
	return factn
}

func fact4(n uint) uint  {
	var factn  uint= 1
	a := make([]int, n)

	for i, _ := range a {
		//a *= uint(i + 1)

	}
	return factn
}


func main()  {
	var v uint
	v = fact1(6)
	fmt.Println("fact1:  ",v)

	var v2 uint
	v2 = fact2(6)
	fmt.Println("fact2:  ",v2)

	var v3 uint
	v3 = fact3(6)
	fmt.Println("fact3:  ",v3)

}