package main

import (
	"fmt"
	"log"
	"rpctest/client"
	"rpctest/server"
	"time"
)

func main() {
	done := make(chan int)
	go server.RegisterAndServe() //先启动服务端
	time.Sleep(1e9)              //sleep 1s，因为服务端启动是异步的，所以等一等
	go func() {                  //启动客户端
		result, err := client.CallEcho("hello world")
		if err != nil {
			log.Fatal("error calling", err)
		} else {
			fmt.Println("call echo:", result)
		}
		done <- 1
	}()
	<-done //阻塞等待客户端结束
}
