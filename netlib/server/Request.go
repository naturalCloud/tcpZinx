package server

import "netLearn/netlib/sInterface"

type Request struct {
	//当前面链接
	conn sInterface.Connection
	//客户端请求数据
	data []byte
}

func (r *Request) GetConn() sInterface.Connection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
