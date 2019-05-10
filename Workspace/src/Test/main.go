package main

import "fmt"
import "validator"

/*option 2: import  StringConvert "Test/Helper" */
import  "Test/Helper"

func main() {

	//setup $GOPATH
	//go get github
	fmt.Println("Hello")
	/*option 2: StringConvert.ConvertStringToInt*/
	Helper.ConvertStringToInt()

	validator.CheckValiEmail()
	validator.CheckValiPhone()

}
