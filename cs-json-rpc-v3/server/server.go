package server

import (
    "fmt"
	"os"
	"strconv"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
	"io/ioutil"
)

var path = "../ioFile.txt"

type ArgsSum struct {
    Item1, Item2 int
}

type MyServer int

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
	*reply = "" + s
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

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}


func StartServer() {
	myServer := new(MyServer)
	
	rpc.Register(myServer)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        jsonrpc.ServeConn(conn)
    }
}

func main() {
	go StartServer()
}


