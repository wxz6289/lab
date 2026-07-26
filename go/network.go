package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func server() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	listen, err := net.ListenTCP("tcp",addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		fmt.Println(conn.RemoteAddr())
		conn.Write([]byte("Hello world"))
		time.Sleep(3*time.Second)
		conn.Close()
	}
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		byteData := make([] byte, 1024)
		n, err := conn.Read(byteData)
		if err == io.EOF {
			break
		}
		fmt.Println(string(byteData[:n]))
	}
}