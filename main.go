package main

import "CNCServer"

func main() {
	server := CNCServer.CNCServer{}
	server.Initialize()
	server.Run(":8081")
}
