/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here

package main
import (
"fmt"

)


func main()  {
	var numberTest int
	fmt.Scan(&numberTest)
	var count int

	for {
		if count>numberTest{

			fmt.Scan(&numberTest)
		}
		count++
	}
}