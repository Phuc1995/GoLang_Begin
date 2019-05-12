package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var i int = 1234
	s1 := fmt.Sprintf("%d", i)
	s2 := fmt.Sprintf("%x", i)
	s3 := strconv.Itoa(i)
	s4 := strconv.FormatInt(int64(i), 16)
	fmt.Println(s1) // "1234"
	fmt.Printf("%T",i)
	fmt.Println()
	fmt.Println(s2)  // "4d2" là giá trị của 1234 trong hệ 16

	fmt.Println(s3)  // "1234"
	fmt.Printf("Type s3: %T",s3)
	fmt.Println()

	fmt.Println(s4)  // "4d2" là giá trị của 1234 trong hệ 16
	fmt.Printf("Type s4: %T",s4)
	fmt.Println()

	s := "1234"
	x1, _ := strconv.Atoi(s)
	x2, _ := strconv.ParseInt(s, 10, 32)
	s = "4d2"
	x3, _ := strconv.ParseInt(s, 16, 32)
	fmt.Println(x1)  // 1234
	fmt.Printf("Type x1: %T",s3)
	fmt.Println()

	fmt.Println(x2)  // 1234
	fmt.Printf("Type x2: %T",s3)
	fmt.Println()

	fmt.Println(x3)  // 1234  là giá trị trong hệ 10, ứng với 4d2 trong hệ 16
	fmt.Printf("Type x3: %T",s3)
	fmt.Println()
}