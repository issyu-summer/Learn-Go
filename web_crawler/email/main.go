package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reQQEmail = `(\d+)@qq.com`
)

func GetEmail() {
	//1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?pn=2")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	//2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	//3.字节转字符串
	pageStr := string(pageBytes)
	//4.过滤数据，找出qq邮箱
	re := regexp.MustCompile(reQQEmail)
	//-1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("email", result[0])
		fmt.Println("qq", result[1])
	}
}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func main() {
	GetEmail()
}
