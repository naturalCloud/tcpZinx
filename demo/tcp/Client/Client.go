package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {

	dial, err := net.Dial("tcp", ":8889")
	if err != nil {
		log.Println(err)
		return
	}

	last := 0

	rand.Seed(time.Now().Unix())
	for {

		last = randomAZ(last)
		dial.Write([]byte(string(last)))

		rb := make([]byte, 512)
		read, _ := dial.Read(rb)


		fmt.Println("读到服务端消息--", string(rb[:read]))

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
