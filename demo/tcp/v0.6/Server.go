package main

import (
	"fmt"
	"netLearn/netlib/sInterface"
	"netLearn/netlib/server"
	"netLearn/netlib/util"
)

const (
	Ping uint32 = iota
	Hello
)

type PingRouter struct {
	server.BasRouter
}

//处理业务之前
//func (p *PingRouter) PreHandle(request sInterface.Request) {
//	fmt.Println("call Router PreHandle ", string(request.GetData()))
//	_, err := request.GetConn().GetTcpConnection().Write(request.GetData())
//	if err != nil {
//		fmt.Println("ping err")
//	}
//}

// Handle 处理中
func (p *PingRouter) Handle(request sInterface.Request) {
	fmt.Println("call Router Handle")
	fmt.Println("recv from client: msgId", request.GetMsgId(), "data ", string(request.GetData()))
	err := request.GetConn().SendMsg(request.GetMsgId(), []byte("ping...ping...ping..."))
	if err != nil {
		fmt.Println("发送消息", err)
	}

}

type HelloRouter struct {
	server.BasRouter
}

func (h *HelloRouter) Handle(request sInterface.Request) {
	fmt.Println("call Router Handle")
	fmt.Println("recv from client: msgId", request.GetMsgId(), "data ", string(request.GetData()))
	err := request.GetConn().SendMsg(request.GetMsgId(), []byte("hello ...hello...hello ..."))
	if err != nil {
		fmt.Println("发送消息", err)
	}

}

//处理之后
//func (p *PingRouter) PostHandle(request sInterface.Request) {
//	fmt.Println("call Router PostHandle...")
//	err := request.GetConn().SendMsg(4,[]byte("over"))
//	if err != nil {
//		fmt.Println("over err")
//	}
//}

func main() {

	s := server.New()
	s.AddRouter(Ping, &PingRouter{})
	s.AddRouter(Hello, &HelloRouter{})
	s.Serve()

}

func init() {
	util.Init("/demo/tcp/v0.6")

}
