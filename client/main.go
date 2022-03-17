package main

import (
	"fmt"
	"gosocket/config"
	"gosocket/util"
	"net"
)

func main() {
	fmt.Printf("client connect %v\n", config.ServerAddress)
	conn, err := net.Dial(config.ServerNetworkType, config.ServerAddress)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("client %v connected\n", conn.LocalAddr())

	util.Write(conn, "hello")
	if _, err := util.Read(conn); err != nil {
		fmt.Println(err)
	}
}
