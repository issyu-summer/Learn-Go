package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	//标准输出流就是控制台
	_, err := fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	if err != nil {
		return
	}
	//perm：0八进制标识  -普通文件 r 4 w 2 x 1
	file, err := os.OpenFile("./file_test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败,err:", err)
		return
	}
	name := "hello,go"
	_, err = fmt.Fprintf(file, "使用Fprint()向文件中写入:%s", name)
	if err != nil {
		return
	}
}
