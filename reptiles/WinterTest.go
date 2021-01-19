package main

import (
	"fmt"
	"io"
	"net/http"
)

//爬取页面操作
func working(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页\n", start, end)
	for i := start; i <= end; i++ {
		url := ""
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err:", err)
			continue //该页出错仍然可以读取其他页
		}
		//fmt.Println("result=",result)
		//将读取的整网页保存成一个文件
		f, err := os.Creat("第" + strconv.itoa(i) + "页" + "html")
		if err != nil {
			fmt.Println("Create err :", err)
			continue
		}
		f.WriteString(result)
		f.Close() // 保存好一个， 存一个
	}
}

func HttpGet(url string) (result string) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//循环读取网页数据，传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加每一次读取的数据
		result += string(buf[:n])
	}
	return
}

func main() {
	//指定爬取起始，终止页
	var start, end int
	fmt.Printf("输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Printf("输入终止页（>=start）:")
	fmt.Scan(&end)

	working(start, end)
}
