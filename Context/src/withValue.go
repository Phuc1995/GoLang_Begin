package main

import (
	"fmt"
	"golang.org/x/net/context"
)

func main()  {
	ctx := context.WithValue(context.Background(),"number",10)
	A(ctx)
}

func A(ctx context.Context)  {
	if value := ctx.Value("number3"); value != nil{
		ctx := context.WithValue(ctx, "number1",5)
		B(ctx)
	}else {
		fmt.Println("key not exist")
	}
}

func B(ctx context.Context)  {
	value1 := ctx.Value("number")
	value2 := ctx.Value("number1")
	fmt.Println("number: ", value1)
	fmt.Println("number: ", value2)

}