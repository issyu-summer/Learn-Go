package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wr() {
	//打开文件，如果不存在则创建
	file, err := os.OpenFile("./file2_test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,", err)
		return
	}
	defer file.Close()
	//获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	//刷新缓冲区
	writer.Flush()
}

func re() {
	file, err := os.Open("./file2_test.txt")
	if err != nil {
		fmt.Println("open file failed,", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read data failed,", err)
			return
		}
		fmt.Println(string(line))
	}
}

func main() {
	wr()
	re()
}
