package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// func fn Hello(request, *response) error

type HelloService struct {
}

// RPC 函数签名
func (s *HelloService) Hello(request string, response *string) error {
	// request :tom
	// resone: hello,tom

	*response = fmt.Sprintf("RPC: test.%s", request)
	return nil
}

// 暴露为一个RPC HelloService.Hello
// 把HelloService 这个对象给RPC框架,RPC会把这个对象的方法暴露到网络中
func main() {
	//注册对象到rpc框架
	rpc.RegisterName("HelloService", &HelloService{})

	//启动rpc服务
	//然后我们建立一个唯一的tcp链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	// 没Accept一个请求，就创建一个goroutie进行处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 前面都是tcp的知识, 到这个RPC就接管了
		// 因此 你可以认为 rpc 帮我们封装消息到函数调用的这个逻辑,
		// 提升了工作效率, 逻辑比较简洁，可以看看他代码
		go rpc.ServeConn(conn)
	}
}
