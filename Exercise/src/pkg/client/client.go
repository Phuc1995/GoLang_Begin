package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"math/rand"
	"os"
	"time"
)

type Message struct {
	Text string "text"
	Mod int
}
var (

	port = flag.String("port", "9000", "port used for ws connection")
)

func connect() (*websocket.Conn, error) {
	return websocket.Dial(fmt.Sprintf("ws://localhost:%s", *port), "", mockedIP())
}

func mockedIP() string {
	var arr [4]int
	for i := 0; i < 4; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(256)
	}
	return fmt.Sprintf("http://%d.%d.%d.%d", arr[0], arr[1], arr[2], arr[3])
}

func main() {
	flag.Parse()
	// connect
	ws, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// receive
	var m Message
	var mod int

	go func() {
		for {
			err := websocket.JSON.Receive(ws, &m)
			if err != nil {
				fmt.Println("Error receiving message: ", err.Error())
				break
			}
			fmt.Println("Message: ", m)
		}
	}()
	// send

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choice mod(1:mod, 2:nor): ")
	fmt.Scan(&mod)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		m := Message{
			Text: text,
		}
		err = websocket.JSON.Send(ws, m)
		if err != nil {
			fmt.Println("Error sending message: ", err.Error())
			break
		}
	}
}