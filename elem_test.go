package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

var client *rpc.Client

func TestStartCon(t *testing.T) {
	go StartServer()
	conn, err := net.Dial("tcp", "localhost:4040")
	if err != nil {
		log.Fatal(err)
	}
	client = jsonrpc.NewClient(conn)
}

func TestSuma(t *testing.T) {

	var replySuma int
	var replyScris string
	var replyCitit int

	path := "testReadWrite.txt"
	expected := 5

	args := &ArgsSuma{2, 3}
	if err := client.Call("MyServer.Suma", args, &replySuma); err != nil {
		t.Error(err)
	} else if replySuma != expected {
		t.Errorf("Expected : <%+v>,received: <%+v>", expected, replySuma)
	}

	args2 := &ArgsScris{replySuma, path}
	if err := client.Call("MyServer.Scris", args2, &replyScris); err != nil {
		t.Error(err)
	} else if replyScris != "works" {
		t.Errorf("does not work")
	}

	args3 := &ArgsCitit{path}
	if err := client.Call("MyServer.Citit", args3, &replyCitit); err != nil {
		t.Error(err)
	} else if replyCitit != expected {
		t.Errorf("Expected : <%+v>,received: <%+v>", expected, replyCitit)
	}

}
