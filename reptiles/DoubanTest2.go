package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://book.douban.com/tag/%25E5%2593%25B2%25E5%25AD%25A6?start=20&type=T"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "bid=mv8ww0YyRhY")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
