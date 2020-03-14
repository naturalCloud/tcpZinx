package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	//服务器名称
	Name string
	//版本
	IPVersion string
	//端口
	Port int
}

//开启服务
func (s *Server) Start() {
	fmt.Println("server will start")

	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", "127.0.0.1", s.Port))

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

		Connection := NewConnection(conn, cid, func(conn *net.TCPConn, bytes []byte, i int) error {
			fmt.Println("回显回调函数----")
			if _, err2 := conn.Write(bytes[:i]); err2 != nil {
				fmt.Println("回显错误")

				return nil
			}

			return nil
		})
		cid++
		go Connection.Start()

	}

}

//停止服务
func (s *Server) Stop() {

}

//运行服务
func (s *Server) Serve() {
	s.Start()
	select {}
}

func New(Name string) *Server {

	return &Server{
		Name:      Name,
		IPVersion: "tcp4",
		Port:      8868,
	}

}
