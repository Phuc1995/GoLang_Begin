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

	go func() {
		for {
			conn, err := server.Accept()
			if err != nil{
				log.Fatal(err)
			}
			fmt.Println("new client")

			conns = append(conns, conn)


			connCh <- conn
		}
		
	}()

	for {
		select {
			case conn := <- connCh:
				go onMessage(conn)

			case msg := <- msgCh:
				fmt.Println(msg)

			case conn := <- closeCh:
				fmt.Println("client exit",conn)
				removeConn(conn)
		}
	}
}

func publishMsg(conn net.Conn, msg string)  {
	for i := range conns{
		if conns[i] != conn{
			conns[i].Write([]byte(msg))
		}
	}
}

func removeConn(conn net.Conn)  {
	var i int
	for i = range conns{
		if conns[i] == conn{
			break
		}
	}

	conns = append(conns[i:], conns[:i+1]...)
}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')

		if err != nil {
			break
		}
		msgCh <- msg
		publishMsg(conn, msg)
	}

	closeCh <- conn
}