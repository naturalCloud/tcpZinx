package server

import (
	"errors"
	"fmt"
	"io"
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
	//消息管理 msgId 和 对应处理的程序
	MsgHandler sInterface.MessageHandle
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
		//buf := make([]byte, util.ServerConf.MaxBufSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Printf(" %d 读取数据错误  v% \n", c.ConnId, err)
		//	continue
		//}
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTcpConnection(), headData)
		if err != nil {
			fmt.Println("read msg error", err)
			break
		}

		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack err", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
			}

		}

		msg.SetData(data)
		req := Request{
			conn: c,
			msg:  msg,
		}

		//从路由中找到对应的路由处理程序
		go func(request *Request) {
			c.MsgHandler.DoMessageHandle(request)
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
func (c *Connection) SendMsg(msgId uint32, data []byte) error {

	if c.IsClosed {
		return errors.New("connection closed when send msg")
	}
	//将数据封包
	dp := NewDataPack()
	bmsg, err := dp.Pack(NewMessage(msgId, data)) //二进制数据
	if err != nil {
		return errors.New("pack data error")
	}
	if _, err := c.GetTcpConnection().Write(bmsg); err != nil {
		return errors.New("send msg error")
	}

	return nil
}

//获取当前链接id
func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

func NewConnection(conn *net.TCPConn, connId uint32, msgHandler sInterface.MessageHandle) *Connection {
	return &Connection{
		Conn:       conn,
		ConnId:     connId,
		IsClosed:   false,
		MsgHandler: msgHandler,
		ExitChan:   make(chan bool),
	}
}
