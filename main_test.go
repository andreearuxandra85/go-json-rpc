package main_test

import (
	"fmt"
	"os"
	"testing"
	"strconv"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

var path = "../mockIoFile.txt"

type MockMyServer struct { } 

type ArgsWrite struct {
	Item int
	FilePath string
} 

type ArgsRead struct {
	FilePath string
}

type ArgsSum struct {
    Item1, Item2 int
}

func (srv *MockMyServer) Sum(args ArgsSum, reply *int) error {
	*reply = args.Item1 + args.Item2
    return nil
}

func (srv *MockMyServer) Write(args ArgsWrite, reply *string) error {
	file, err := os.Create(args.FilePath)
    if err != nil {
        fmt.Println(err.Error())
    }
    defer file.Close() 

	s := strconv.Itoa(args.Item)
    file.WriteString(s)
	*reply = s
	return nil
}

func (srv *MockMyServer) Read(args ArgsRead, reply *int) error {
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


func TestSum(t *testing.T) {
	// arrange
	mock := &MockMyServer{}
	
	args := ArgsSum{2, 3}	
	var reply int
	expectedResult := 5
	
	// act
	err := mock.Sum(args, &reply)
	
	// assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, reply)
}


func TestSumWriteToAndReadFromFile(t *testing.T) {
	// arrange
	mock := &MockMyServer{}
	
	args := ArgsSum{2, 3}	
	expectedResult := 5
	expectedResultAsString := "5"
	var reply int
	
	// act
	errSum := mock.Sum(args, &reply)
	
		// MyServer.Write()
	argsWrite := ArgsWrite{reply, path}
	var replyWrite string
	errWrite := mock.Write(argsWrite, &replyWrite)
	
		// MyServer.Read()
	argsRead := ArgsRead{path}
	var replyRead int
	errRead := mock.Read(argsRead, &replyRead)
	
	
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