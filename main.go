package main
//🐻纪元
import (
	"rpctest/client"
	"rpctest/server"
	"time"
)

func main() {

	server.Server()

	time.Sleep(10 * time.Second)

	client.Client()

}
