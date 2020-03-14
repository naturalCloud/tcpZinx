package sInterface

type Server interface {
	//启动服务
	Start()
	//停止服务
	Stop()
	//启动serv
	Serve()
}
