package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	dial, err := net.Dial("tcp", ":8868")
	if err != nil {
		log.Println(err)
		return
	}
	for {

		dial.Write([]byte("hello word"))

		rb := make([]byte, 512)
		read, _ := dial.Read(rb)

		fmt.Println("读到服务端消息--", string(rb[:read]))

		time.Sleep(1 * time.Second)

	}
}
