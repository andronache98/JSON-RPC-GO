package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	go StartServer()

	var reply int
	conn, err := net.Dial("tcp", "localhost:4040")
	if err != nil {
		log.Fatal(err)
	}
	client := jsonrpc.NewClient(conn)

	args := ArgsSuma{5, 5}
	if err := client.Call("MyServer.Suma", args, &reply); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}

	var reply2 string
	path := "string.txt"
	args2 := &ArgsScris{2, path}
	if err := client.Call("MyServer.Scris", args2, &reply2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply2)

	}

	var reply3 int
	args3 := &ArgsCitit{path}
	if err := client.Call("MyServer.Citit", args3, &reply3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply3)
	}
}
