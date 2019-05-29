package main

import (
	"bufio"
	"fmt"
"log"
"net"
)

var(
	conns []net.Conn
	connCh = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgCh = make(chan string)
)

func main()  {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println("new client")

		conns = append(conns, conn)
		connCh <- conn
	}

	for {
		select {
			case conn := <- connCh:


		}
	}
}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		_, err := reader.ReadString('\n')

		if err != nil {
			break
		}
	}

	closeCh <- conn
}