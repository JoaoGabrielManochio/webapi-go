package main

import (
	"github.com/JoaoGabrielManochio/webapi-go/server"
)

func main() {

	server := server.NewServer()

	server.Run()
}
