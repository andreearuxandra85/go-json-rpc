package main

import (
	"../server"
	"../client"
)


func main() {
	go server.StartServer()
	client.ConnectClient()
}

