package sInterface

//把客户端请求的链接信息 请求数据
type Request interface {
	//获取链接
	GetConn() Connection
	//获取数据
	GetData() []byte
}
