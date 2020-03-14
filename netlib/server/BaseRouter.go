package server

import "netLearn/netlib/sInterface"

type BasRouter struct{}

//
//基类Router
//
//
//处理业务之前钩子
func (b *BasRouter) PreHandle(request sInterface.Request) {
}

//处理业务中前钩子
func (b *BasRouter) Handle(request sInterface.Request) {
}

//处理业务后钩子
func (b *BasRouter) PostHandle(request sInterface.Request) {
}
