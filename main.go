package main

import (
    "fmt"
	"os"
	"strconv"
    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
	"io/ioutil"
)

var path = "../ioFile.txt"

type ArgsSum struct {
    Item1, Item2 int
}

type MyServer struct { } 

type ArgsWrite struct {
	Item int
	FilePath string
} 

type ArgsRead struct {
	FilePath string
}


func (srv *MyServer) Write(args ArgsWrite, reply *string) error {
	file, err := os.Create(args.FilePath)
    if err != nil {
        fmt.Println(err.Error())
    }
    defer file.Close() 
	
	s := strconv.Itoa(args.Item)
    file.WriteString(s)
	return nil
}


func (srv *MyServer) Read(args ArgsRead, reply *int) error {

	file, err := os.Open(args.FilePath)
    if err != nil {
        fmt.Println(err.Error())
    }
    defer file.Close()
	
	
	content, err := ioutil.ReadFile(args.FilePath)
	if err != nil {
		fmt.Println("Error reading file :", err)
	}
	
	intFileContent, err := strconv.Atoi(string(content))
	if err != nil {
		fmt.Println("Atoi conversion error")
	}
	
	*reply = intFileContent
	return nil
}


func (srv *MyServer) Sum(args ArgsSum, reply *int) error {
	*reply = args.Item1 + args.Item2
    return nil
}


func startServer() {
	
	server := rpc.NewServer()
	srv := new(MyServer)
	server.Register(srv)

    l, e := net.Listen("tcp", "localhost:8333")
    if e != nil {
        log.Fatal("listen error:", e)
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go server.ServeCodec(jsonrpc.NewServerCodec(conn))
    }
}



func main() {

    go startServer()
	
	conn, err := net.Dial("tcp", "localhost:8333")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    client := jsonrpc.NewClient(conn)

	
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