package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main()  {
	connecton, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(connecton)

	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')

	nameInput = nameInput[:len(nameInput)-1]
	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')

		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s: %s\n",nameInput,msg[:len(msg)-1])
		connecton.Write([]byte(msg))

	}
	fmt.Println(nameInput)
}