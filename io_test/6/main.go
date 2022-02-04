package main

import (
	"fmt"
	"io/ioutil"
)

func wr() {
	err := ioutil.WriteFile("./file3_test.txt", []byte("hello,world"), 0666)
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
}

func re() {
	content, err := ioutil.ReadFile("./file3_test.txt")
	if err != nil {
		fmt.Println("access data failed,", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	wr()
	re()
}
