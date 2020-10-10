package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strconv"
)

type MyServer struct{}

func (srv *MyServer) Suma(args ArgsSuma, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func (srv *MyServer) Scris(args ArgsScris, reply *string) error {
	file, err := os.Create(args.FilePath)
	if err != nil {
		return err
	}
	fmt.Fprint(file, args.A)
	*reply = "works"
	return nil
}

func (srv *MyServer) Citit(args ArgsCitit, reply *int) error {
	file, err := ioutil.ReadFile(args.FilePath)
	if err != nil {
		return err
	}
	number, err := strconv.Atoi(string(file))
	*reply = number
	return nil
}

func StartServer() {
	fmt.Println("Start server")

	srv := new(MyServer)

	server := rpc.NewServer()

	server.Register(srv)

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := listener.Accept()
		fmt.Println("Connection established")
		if err != nil {
			log.Fatal("Connection error")
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(connection))

		defer connection.Close()

	}

}
