package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func HttpGet2(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	//循环爬取整页数据
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPage2(idx int) {
	//获取 url
	url := "https://book.douban.com/tag/%E5%93%B2%E5%AD%A6?start=" + strconv.Itoa((idx-1)*20) + "&type=T"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	if err != nil {
		fmt.Println("HttpGet2 err:", err)
		return
	}
	fmt.Println("result=", req.Body)
}

func toWork(start, end int) {
	fmt.Printf("正在爬取%d到%d页...\n", start, end)
	for i := start; i <= end; i++ {
		SpiderPage2(i)
	}
}

func main() {
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)
	toWork(start, end)
}
