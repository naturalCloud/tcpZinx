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
	//路由
	Router sInterface.Router
	Host   string
}

//开启服务
func (s *Server) Start() {

	fmt.Printf("server %s Host %s Port %d start",s.Name,s.Host,s.Port)

	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Host, s.Port))

	if err != nil {
		log.Fatal(err)
		return
	}

	tcp, _ := net.ListenTCP(s.IPVersion, addr)

	for {
		conn, err := tcp.AcceptTCP()
		if err != nil {
			log.Println(err, "错误")
		}

		var cid uint32

		Connection := NewConnection(conn, cid, s.Router)
		cid++
		go Connection.Start()

	}

}

//停止服务
func (s *Server) Stop() {

}

//运行服务
func (s *Server) Serve() {
	if s.Router == nil {
		fmt.Println("路由未设置,终止..")
		return
	}
	s.Start()
	select {}
}

//添加router
func (s *Server) AddRouter(router sInterface.Router) {
	s.Router = router
}

func New() *Server {

	return &Server{
		Name:      util.ServerConf.Name,
		IPVersion: "tcp4",
		Port:      util.ServerConf.Port,
		Router:    nil,
		Host:      util.ServerConf.Host,
	}

}
