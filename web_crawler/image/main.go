package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	//存放图片链接的管道
	chanImageUrls chan string
	wg            sync.WaitGroup
	//监控协程的管道
	chanTask chan string
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.get.url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "resp.body")
	filename = "D:\\img\\" + filename
	fmt.Println("目录：", filename)
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

// DownloadImg 下载图片
func DownloadImg() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	wg.Done()
}

// GetFilenameFromUrl 截取url的名字
func GetFilenameFromUrl(url string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

func CheckOut() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 5 {
			close(chanImageUrls)
			break
		}
	}
	wg.Done()
}

//爬取图片链接到管道
//url是传的整页链接
func getImgUrls(url string) {
	urls := getImages(url)
	//遍历切片中的所有链接，存入数据管道
	for _, url := range urls {
		chanImageUrls <- url
	}
	//标识当前协程完成任务
	//没完成一个任务，写一条数据
	//用于监控协程知道已经完成了几个任务
	chanTask <- url
	wg.Done()
}

// getImages 获取当前页的图片链接
func getImages(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n\n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

// GetPageStr 根据Url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get.url")
	defer resp.Body.Close()
	//读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	//字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

func main() {
	//1.初始化管道
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 5)
	//2.爬虫协程
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	//3.任务统计协程，统计26个任务是否都完成，完成则关闭管道
	wg.Add(1)
	go CheckOut()
	//4.下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}
