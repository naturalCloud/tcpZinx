package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {

	dial, err := net.Dial("tcp", ":8868")
	if err != nil {
		log.Println(err)
		return
	}
	strArray := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
	}

	leng := len(strArray) - 1
	for {

		int63 := rand.Intn(leng)

		dial.Write([]byte(strArray[int63]))

		rb := make([]byte, 512)
		read, _ := dial.Read(rb)

		fmt.Println("读到服务端消息--", string(rb[:read]))

		time.Sleep(1 * time.Second)

	}
}
