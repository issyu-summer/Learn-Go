package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("./file1_test.txt")
	if err != nil {
		fmt.Println("open file failed,", err)
		return
	}
	defer file.Close()
	//定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			//读取结束
			break
		}
		if err != nil {
			fmt.Println("read file failed,", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))

}
