package server

import "netLearn/netlib/sInterface"

type Request struct {
	//当前面链接
	conn sInterface.Connection
	//客户端请求数据
	msg  sInterface.Message
}

func (r *Request) GetConn() sInterface.Connection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//获取请求消息的Id
func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
