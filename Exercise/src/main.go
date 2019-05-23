package main

import (
	"flag"
	//server "pkg/Server"
)
var (
	port = flag.String("port", "9000", "port used for ws connection")
)
func main()  {
	flag.Parse()
	//log.Fatal(server.server())
}
