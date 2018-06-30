package main_test

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

var path = "../mockIoFile.txt"

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

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}




func TestSumWriteToAndReadFromFile(t *testing.T) {
	// arrange
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }
	
	args := ArgsSum{2, 3}	
	expectedResult := 5
	expectedResultAsString := "5"
	var reply int
	
	// act
	errSum := client.Call("MyServer.Sum", args, &reply)
	if errSum != nil {
		log.Fatal("MyServer.Sum error:", errSum)
	}


	// MyServer.Write()
	argsWrite := ArgsWrite{reply, path}
	var replyWrite string
	
	errWrite := client.Call("MyServer.Write", argsWrite, &replyWrite)
	if errWrite != nil {
		log.Fatal("MyServer.Write error:", errWrite)
	}

	// MyServer.Read()
	argsRead := ArgsRead{path}
	var replyRead int
	errRead := client.Call("MyServer.Read", argsRead, &replyRead)
	if errRead != nil {
		log.Fatal("MyServer.Read error:", errRead)
	} 

	// assert Sum
	assert.Nil(t, errSum)
	assert.Equal(t, expectedResult, reply)
	
	// assert Write
	assert.Nil(t, errWrite)
	assert.Equal(t, expectedResultAsString, replyWrite)
	
	// aseert Read
	assert.Nil(t, errRead)
	assert.Equal(t, expectedResult, replyRead)

}

