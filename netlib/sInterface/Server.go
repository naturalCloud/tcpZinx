package sInterface

type Server interface {
	//启动服务
	Start()
	//停止服务
	Stop()
	//启动serv
	Serve()

	//添加路由,给当前服务注册路由,什么消息id,对应什么router
	AddRouter(uint32,Router)
}
