package sInterface

//链接管理模块
type ConnectionManage interface {
	//添加链接
	AddConn(conn Connection)

	//删除链接
	RemoveConn(conn Connection)

	//查询链接
	GetConn(connId uint32) (Connection, error)

	//当前连接数
	Len() int

	//清除链接
	ClearConn()
}
