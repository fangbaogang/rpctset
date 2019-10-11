package server

import (
	"log"
	"net/http"
	"net/rpc"
)

type EchoService struct{}

func (service EchoService) Echo(arg string, result *string) error {
	*result = arg //在这里直接将第二个参数（也就是实际的返回值）赋值为arg
	return nil    //error返回nil，也就是没有发生异常
}
func RegisterAndServe() {
	err := rpc.Register(&EchoService{}) //注册并不是注册方法，而是注册EchoService的一个实例
	if err != nil {
		log.Fatal("error registering", err)
		return
	}
	rpc.HandleHTTP()                        //rpc通信协议设置为http协议
	err = http.ListenAndServe(":9999", nil) //端口设置为9999
	if err != nil {
		log.Fatal("error listening", err)
	}
}
