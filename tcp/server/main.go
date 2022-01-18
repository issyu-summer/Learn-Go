package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader:=bufio.NewReader(conn)
		var buf[128]byte
		n,err:=reader.Read(buf[:])
		if err!=nil{
			fmt.Println("read from client failed,err：",err)
			break
		}
		receiverStr:=string(buf[:n])
		fmt.Println("msg accessed from client:",receiverStr)
		conn.Write([]byte(receiverStr))
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20001")
	if err!=nil{
		fmt.Println("listen failed,err",err)
		return
	}
	for{
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("accepted failed,err",err)
			continue
		}
		go process(conn)
	}
}
