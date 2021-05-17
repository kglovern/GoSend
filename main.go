package main

import "github.com/kglovern/GoSend/CNCServer"

func main() {
	server := CNCServer.CNCServer{}
	server.Initialize()
	server.Run(":8081")
}
