package main

import "netLearn/netlib/server"

func main() {

	s := server.New("tcp1.00")
	s.Serve()



}
