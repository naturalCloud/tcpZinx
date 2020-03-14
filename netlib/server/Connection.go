package server

import (
	"fmt"
	"net"
	"netLearn/netlib/sInterface"
)

type Connection struct {
	//当前链接socket TCP套接字
	Conn *net.TCPConn
	//链接id
	ConnId uint32
	//当期链接状态
	IsClosed bool
	//退出的channel
	ExitChan chan bool

	//当前链接Router
	Router sInterface.Router
}

//启动链接
func (c *Connection) Start() {
	fmt.Println("conn start connId ", c.ConnId)

	go c.StartReader()

}

//读取任务携程
func (c *Connection) StartReader() {

	fmt.Printf("reader is running connId = %d , addr = %s \n", c.ConnId, c.RemoteAddr().String())
	defer fmt.Printf(" connId = %d 关闭 \n")
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf(" %d 读取数据错误  v% \n", c.ConnId, err)
			continue
		}

		req := Request{
			conn: c,
			data: buf,
		}
		go func(request *Request) {
			fmt.Println("Router will run")
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}

}

//停止服务
func (c *Connection) Stop() {
	fmt.Printf("conn stop connId %d \n", c.ConnId)
	//链接已经关闭
	if c.IsClosed {
		return
	}
	defer close(c.ExitChan)
	defer c.Conn.Close()
	c.IsClosed = true
}

//获取当前链接
func (c *Connection) GetTcpConnection() *net.TCPConn {

	return c.Conn
}

//获取当前链接远程地址
func (c *Connection) RemoteAddr() net.Addr {

	return c.Conn.RemoteAddr()
}

//发送数据
func (c *Connection) Send(data []byte) error {

	return nil
}

//获取当前链接id
func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

func NewConnection(conn *net.TCPConn, connId uint32, router sInterface.Router) *Connection {
	return &Connection{
		Conn:     conn,
		ConnId:   connId,
		IsClosed: false,
		Router:   router,
		ExitChan: make(chan bool),
	}
}
