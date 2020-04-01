package sInterface

import "net"

type Connection interface {

	//启动链接

	Start()
	//停止链接
	Stop()
	//获取当前链接 conn

	GetTcpConnection() *net.TCPConn

	//获取远程地址
	RemoteAddr() net.Addr
	//发送数据

	SendMsg(msgId uint32, data []byte) error
	//获取当前的链接id
	GetConnId() uint32

	//
	SetProperty(key string, value interface{})

	GetProperty(key string) (interface{}, error)

	RemoveProperty(key string)
}

//连处理函数
type HandelFun func(*net.TCPConn, []byte, int) error
