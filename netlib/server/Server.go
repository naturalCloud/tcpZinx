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
}

//开启服务
func (s *Server) Start() {

	fmt.Printf("server %s Host %s Port %d start", s.Name, s.Host, s.Port)

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

		Connection := NewConnection(conn, cid, s.MsgHandler)
		cid += 1
		go Connection.Start()

	}

}

//停止服务
func (s *Server) Stop() {

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
	}

}
