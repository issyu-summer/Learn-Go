package main

import "os"

func main() {
	//以文件的方式操作终端
	var buf [32]byte
	os.Stdin.Read(buf[:])
	os.Stdin.WriteString(string(buf[:]))
	os.Stdout.WriteString(string(buf[:]))
}
