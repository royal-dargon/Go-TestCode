package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
	//"strconv"
)

type result struct {
	channel_id  int
	name        string
	view_count  string
	like_count  string
	author_name string
	author_id   string
	bvid        string
}

func main() {
	t1 := time.Now()
	for Id := 1; Id <= 1000; Id++ {
		Worker(Id)
	}
	t2 := time.Since(t1)
	fmt.Println(t2)
}

func Worker(i int) {
	var res result
	client := &http.Client{}
	p := strconv.Itoa(i)
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/web/channel/multiple/list?channel_id="+p+"&sort_type=hot&page_size=30", nil)
	if err != nil {
		panic(err)
		log.Println(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	result, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err2)
	}
	for a := 0; a <= 30; a++ {
		num := strconv.Itoa(a)
		res.channel_id = i
		res.name = gjson.Get(string(result), "data.list."+num+".name").String()
		res.view_count = gjson.Get(string(result), "data.list."+num+".view_count").String()
		res.like_count = gjson.Get(string(result), "data.list."+num+".like_count").String()
		res.author_name = gjson.Get(string(result), "data.list."+num+".author_name").String()
		res.author_id = gjson.Get(string(result), "data.list."+num+".author_id").String()
		res.bvid = gjson.Get(string(result), "data.list."+num+".bvid").String()
		if res.name != "" {
			fmt.Println(res)
		}
	}

}
