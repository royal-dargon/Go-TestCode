package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/web/channel/multiple/list?channel_id=5417&sort_type=hot&page_size=30", nil)
	if err != nil {
		panic(err)
		fmt.Println(err)
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
	//fmt.Printf("%s", result)
	for i := 0; i < 30; i++ {
		num := strconv.Itoa(i)
		name := gjson.Get(string(result), "data.list."+num+".name")
		view_count := gjson.Get(string(result), "data.list."+num+".view_count")
		like_count := gjson.Get(string(result), "data.list."+num+".like_count")
		author_name := gjson.Get(string(result), "data.list."+num+".author_name")
		author_id := gjson.Get(string(result), "data.list."+num+".author_id")
		bvid := gjson.Get(string(result), "data.list."+num+".bvid")
		fmt.Println(name, view_count, like_count, author_name, author_id, bvid)
	}

}
