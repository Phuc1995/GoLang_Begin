package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("%q\n", p)
	return len(p), nil
}
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

func main()  {
	s := strings.NewReader("onetwothree")
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 2)
	bw.ReadFrom(s)
	err := bw.Flush()
	if err != nil {
		panic(err)
	}

	/*err := bw.Flush()
	fmt.Println(err)*/
}
