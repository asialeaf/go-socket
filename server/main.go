package main

import (
	"fmt"
	"gosocket/config"
	"gosocket/util"
	"io"
	"net"
	"time"
)

func handleConnect(conn net.Conn) {
	fmt.Printf("client %v connected\n", conn.RemoteAddr())
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(2))) //设置读取超时时间
		if _, err := util.Read(conn); err != nil {
			if err == io.EOF {
				fmt.Printf("client %v closed\n", conn.RemoteAddr())
				break
			} else {
				fmt.Printf("read error:%v\n", err.Error())
			}
		} else {
			util.Write(conn, "welcome")
		}
	}
}

func main() {
	fmt.Printf("server start %v\n", config.ServerAddress)
	listener, err := net.Listen(config.ServerNetworkType, config.ServerAddress)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Printf("waiting client connect...\n")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnect(conn)
	}
}
