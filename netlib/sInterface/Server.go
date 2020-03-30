package sInterface

type Server interface {
	//启动服务
	Start()
	//停止服务
	Stop()
	//启动serv
	Serve()

	//添加路由,给当前服务注册路由,什么消息id,对应什么router
	AddRouter(uint32, Router)
	//获取当前链接测管理
	GetConnMgr() ConnectionManage

	//注册OnConnStart 钩子函数
	SetOnConnStart(func(connection Connection))

	//注册OnConnStop 钩子函数
	SetOnConnStop(func(connection Connection))

	//调用 OnConnStart 钩子函数
	CallOnConnStart(connection Connection)

	//调用 OnConnStop 钩子函数
	CallOnConnStop(connection Connection)
}
