package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//初始化一个客户端
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 然后通过client.Call调用具体的RPC方法
	// 在调用client.Call时:
	// 		第一个参数是用点号链接的RPC服务名字和方法名字，
	// 		第二个参数是 请求参数
	//      第三个是请求响应, 必须是一个指针, 有底层rpc服务帮你赋值
	var reply string
	err = client.Call("HelloService.Hello", "测试rpc", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

/*
1. 启动server  go run server/main.go
2. go run client/main.go
*/
