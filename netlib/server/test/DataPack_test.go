package test

import (
	"fmt"
	"io"
	"net"
	"netLearn/netlib/server"
	"testing"
)

func TestDataPack(t *testing.T) {

	listen, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		t.Error("链接创建错误", err)
	}

	go func() {

		for {
			conn, err := listen.Accept()
			if err != nil {
				t.Error("accept err", err)
			}

			go func(conn net.Conn) {
				dp := server.NewDataPack()
				for {
					//先读取head
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						t.Error("读取head 头部数据错误", err)
					}

					//解压header数据
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						t.Error("unpack error", err)
					}
					if msgHead.GetDataLen() > 0 {
						//有数据,需要进行二次读取
						//第二次从conn读取数据,根据head中data len
						msg := msgHead.(*server.Message)
						msg.Data = make([]byte, msg.GetDataLen())
						//根据data len 长度再次从io流中读取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							t.Error("server unpack data err :", err)
						}

						fmt.Println("data--->", string(msg.Data))

					}

					t.Log("--------> recv msgId ", msgHead.GetMsgId(), "data len", msgHead.GetDataLen())
				}

			}(conn)

		}
	}()

	//客户端
	conn, err := net.Dial("tcp", "127.0.0.1:6666")

	if err != nil {
		t.Error("client err", err)
		return
	}

	//创建封包对象
	pack := server.NewDataPack()

	data := "我是中国人"
	//封装第一个msg包
	msg1 := &server.Message{
		MessageId: 1,
		DataLen:   uint32(len(data)),
		Data:      []byte(data),
	}

	data2 := "hello word"
	msg2 := &server.Message{
		MessageId: 2,
		DataLen:   uint32(len(data2)),
		Data:      []byte(data2),
	}

	bytes1, _ := pack.Pack(msg1)
	bytes2, _ := pack.Pack(msg2)
	//合并两个包
	sendData := append(bytes1, bytes2...)
	conn.Write(sendData)
	select {}
}
