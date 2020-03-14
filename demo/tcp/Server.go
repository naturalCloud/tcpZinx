package main

import (
	"fmt"
	"netLearn/netlib/sInterface"
	"netLearn/netlib/server"
)

type PingRouter struct {
	server.BasRouter
}

//处理业务之前
func (p *PingRouter) PreHandle(request sInterface.Request) {
	fmt.Println("call Router PreHandle ", string(request.GetData()))
	_, err := request.GetConn().GetTcpConnection().Write(request.GetData())
	if err != nil {
		fmt.Println("ping err")
	}
}

//处理中
func (p *PingRouter) Handle(request sInterface.Request) {
	fmt.Println("call Router Handle")

}

//处理之后
func (p *PingRouter) PostHandle(request sInterface.Request) {
	fmt.Println("call Router PostHandle...")
	err := request.GetConn().Send([]byte("over"))
	if err != nil {
		fmt.Println("over err")
	}
}

func main() {

	s := server.New("tcp1.00")
	s.AddRouter(&PingRouter{})
	s.Serve()

}
