package server

import (
	"fmt"
	"log"
	"net"
	"netLearn/netlib/sInterface"
	"netLearn/netlib/util"
)

type Server struct {
	//服务器名称
	Name string
	//版本
	IPVersion string
	//端口
	Port int

	Host string
	//当前消息的Message handler
	MsgHandler sInterface.MessageHandle

	//server 链接管理模块
	connMgr sInterface.ConnectionManage

	//server创建后的函数
	connStart func(connection sInterface.Connection)

	//server 销毁后的函数
	connStop func(connection sInterface.Connection)
}

//注册OnConnStart 钩子函数
func (s *Server) SetOnConnStart(start func(connection sInterface.Connection)) {
	s.connStart = start
}

//注册OnConnStop 钩子函数
func (s *Server) SetOnConnStop(stop func(connection sInterface.Connection)) {
	s.connStop = stop
}

//调用 OnConnStart 钩子函数
func (s *Server) CallOnConnStart(connection sInterface.Connection) {
	if s.connStart != nil {
		fmt.Println("-----> Call OnConnStart() ...")
		s.connStart(connection)
	}
}

//调用 OnConnStop 钩子函数
func (s *Server) CallOnConnStop(connection sInterface.Connection) {

	if s.connStop != nil {
		fmt.Println("-----> Call OnConnStop() ...")
		s.connStop(connection)
	}
}

//获取链接管理器
func (s *Server) GetConnMgr() sInterface.ConnectionManage {
	return s.connMgr
}

//开启服务
func (s *Server) Start() {

	fmt.Printf("server %s Host %s Port %d start \n", s.Name, s.Host, s.Port)

	s.MsgHandler.StarWorkPool()
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Host, s.Port))

	if err != nil {
		log.Fatal(err)
		return
	}

	tcp, _ := net.ListenTCP(s.IPVersion, addr)

	var cid uint32
	for {
		conn, err := tcp.AcceptTCP()
		if err != nil {
			log.Println(err, "错误")
		}

		//建立链接超过最大链接数时候丢掉链接
		if uint32(s.connMgr.Len()) >= util.ServerConf.MaxConn {
			pack, _ := NewDataPack().Pack(BuildTextMsg(4, "too many Connections"))

			_, _ = conn.Write(pack)
			_ = conn.Close()
			fmt.Println("too many Connections ,current connectionLen ---> ", s.connMgr.Len())
			continue
		}

		Connection := NewConnection(conn, cid, s.MsgHandler, s)
		cid += 1
		go Connection.Start()

	}

}

//停止服务
func (s *Server) Stop() {

	//服务资源回收
	fmt.Println("回收资源开始......")
	s.connMgr.ClearConn()
	fmt.Println("回收资源成功......")

}

//运行服务
func (s *Server) Serve() {
	if s.MsgHandler == nil {
		fmt.Println("路由未设置,终止..")
		return
	}
	s.Start()
	select {}
}

//添加router
func (s *Server) AddRouter(msgId uint32, router sInterface.Router) {
	s.MsgHandler.AddRouterMap(msgId, router)
}

func New() sInterface.Server {

	return &Server{
		Name:       util.ServerConf.Name,
		IPVersion:  "tcp4",
		Port:       util.ServerConf.Port,
		Host:       util.ServerConf.Host,
		MsgHandler: NewMessageHandler(),
		connMgr:    NewConnectionManager(),
	}

}
