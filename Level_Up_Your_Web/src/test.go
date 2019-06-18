package main

import (
	"errors"
	"fmt"

)

type MyError struct {
	msg string
}

func (error *MyError) Error() string {
	return error.msg
}

func ThisFunctionReturnError() error {
	return errors.New("")
}

func main()  {

	//fmt.Print(a)
}