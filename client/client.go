package client

import "net/rpc"

func CallEcho(arg string) (result string, err error) {
	var client *rpc.Client
	client, err = rpc.DialHTTP("tcp", ":9999") //通过rpc.DialHTTP创建一个client
	if err != nil {
		return "", err
	}
	err = client.Call("EchoService.Echo", arg, &result) //通过类型加方法名指定要调用的方法
	if err != nil {
		return "", err
	}
	return result, err
}
