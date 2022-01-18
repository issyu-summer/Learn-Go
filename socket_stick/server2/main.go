package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"src/learnGo/socket_stick/proto"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader:= bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err==io.EOF{
			return
		}
		if err!=nil{
			fmt.Println("decode msg failed,err:",err)
			break
		}
		fmt.Println("msg accessed from client:",msg)
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
