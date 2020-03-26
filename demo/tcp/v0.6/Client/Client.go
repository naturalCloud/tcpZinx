package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"netLearn/netlib/server"
	"time"
)

func main() {

	fmt.Println("client start ...")
	dial, err := net.Dial("tcp", ":8889")
	if err != nil {
		log.Println(err)
		return
	}

	last := 0

	var msgId uint32 = 0
	rand.Seed(time.Now().Unix())
	for {

		last = randomAZ(last)
		//发送封包数据
		dp := server.NewDataPack()
		msgId += 1

	   var 	msgid  uint32 = 0
		if msgId % 2 == 0 {
			msgid = 0
		}else {
			msgid = 1
		}
		binaryMsg, err := dp.Pack(server.NewMessage( msgid, []byte(string(last))))
		if err != nil {
			fmt.Println("pack error", err)
			break
		}

		if _, err := dial.Write(binaryMsg); err != nil {
			fmt.Println("write err", err)
			break
		}

		//先读取服务器返回的head
		bHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(dial, bHead); err != nil {
			fmt.Println("client read binary  msgHead err", err)
			break
		}
		//解包header
		msgHead, err := dp.UnPack(bHead)
		if err != nil {
			fmt.Println("client unpack msgHead err", err)
			break
		}

		//读取数据
		if msgHead.GetMsgLen() > 0 {
			message := msgHead.(*server.Message) //转为指针接口格式
			message.Data = make([]byte, message.GetMsgLen())
			if _, err := io.ReadFull(dial, message.Data); err != nil {
				fmt.Println("client read data error", err)
				break
			}

			fmt.Println("读到服务端消息-- 消息id--->", message.GetMsgId(), "len--->", message.GetMsgLen(), "data--->", string(message.Data))
		}

		time.Sleep(1 * time.Second)

	}
}

func randomAZ(last int) int {

	if last >= 97 && last <= 122 {
		cha := 90 - 65
		return rand.Intn(cha) + 65
	}

	cha := 122 - 97
	return rand.Intn(cha) + 97

}
