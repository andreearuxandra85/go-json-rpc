package main

import (
    "fmt"
    "log"
    "net/rpc/jsonrpc"
	"net/rpc"
)

var path = "../ioFile.txt"

type ArgsSum struct {
    Item1, Item2 int
}

type ArgsWrite struct {
	Item int
	FilePath string
} 

type ArgsRead struct {
	FilePath string
}

func ConnectClient(reply *rpc.Client) {
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }
	*reply = *client
	return
}


func main() {


    client, err := jsonrpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }
	

	
	// MyServer.Sum()
	args := ArgsSum{2, 3}	
	var reply int
	
	err = client.Call("MyServer.Sum", args, &reply)
	if err != nil {
		log.Fatal("MyServer.Sum error:", err)
	}
	fmt.Printf("MyServer.Sum: %d + %d = %v\n", args.Item1, args.Item2, reply) 
	
	
	// MyServer.Write()
	argsWrite := ArgsWrite{reply, path}
	var replyWrite string
	fmt.Printf("MyServer.Write: writing value %d written to file: %s \n", reply, path)
	err = client.Call("MyServer.Write", argsWrite, &replyWrite)
	if err != nil {
		log.Fatal("MyServer.Write error:", err)
	}
	
	
	// MyServer.Read()
	argsRead := ArgsRead{path}
	var replyRead int
	err = client.Call("MyServer.Read", argsRead, &replyRead)
	if err != nil {
		log.Fatal("MyServer.Read error:", err)
	}
	fmt.Printf("MyServer.Read: file: '%s' contains Sum result: %d \n", argsRead.FilePath, replyRead);
	  
}
