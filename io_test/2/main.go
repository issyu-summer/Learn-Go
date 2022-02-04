package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan bool, 1)
	go func() {
		wg.Add(1)
		//打开文件
		file, err := os.OpenFile("./file_test.txt", os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("open file failed,", err)
			return
		}
		//处理完成后关闭文件
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("close file failed,", err)
			}
			ch <- true
			wg.Done()
		}(file)

		for i := 0; i < 5; i++ {
			file.WriteString("ab\n")
			file.Write([]byte("cd\n"))
		}
	}()

	go func() {
		wg.Add(1)
		//创建文件
		file, err := os.Create("./file1_test.txt")
		if err != nil {
			fmt.Println("create file failed,", err)
			return
		}
		//处理完成后关闭文件
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("close file failed,", err)
			}
			wg.Done()
		}(file)
		//写入nr
		for i := 0; i < 5; i++ {
			file.WriteString("ab\n")
			file.Write([]byte("cd\n"))
		}
	}()

	go func() {
		wg.Add(1)

		var fileCreated = false
		var ok bool
		for {
			fileCreated, ok = <-ch //1、通道关闭后，再取值ok=false
			if ok {
				close(ch)
				break
			}
		}
		if fileCreated {
			//删除文件
			err := os.Remove("./file_test.txt")
			if err != nil {
				fmt.Println("remove file failed,", err)
			} else {
				fmt.Println("./file_text.txt has been removed")
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
