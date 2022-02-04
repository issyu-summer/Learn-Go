package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开源文件
	srcFile, err := os.Open("./file1_test.txt")
	if err != nil {
		fmt.Println("open file failed,", err)
		return
	}
	//创建新文件
	dstFile, err2 := os.Create("./file1_test_copy.txt")
	if err2 != nil {
		fmt.Println("create file failed,", err)
		return
	}
	//处理完成后关闭文件
	defer srcFile.Close()
	defer dstFile.Close()
	//缓冲读取
	//buf := make([]byte, 1024)
	buf := make([]byte, 2)
	for {
		//如果buf不为空，则会覆盖其中的内容
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("error,", err)
			break
		}
		dstFile.Write(buf[:n])
	}
}
