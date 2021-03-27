package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/tidwall/gjson"
)

//var wg sync.WaitGroup

type Result struct {
	Channel_id  int    `gorm:"column:channel_id"`
	Vname       string `gorm:"column:Vname"`
	View_count  string `gorm:"column:view_count"`
	Like_count  string `gorm:"columnlike_count"`
	Author_name string `gorm:"column:author_name"`
	Author_id   string `gorm:"column:author_id"`
	Bvid        string `gorm:"column:bvid"`
}

func WorkerOne(i int, wg sync.WaitGroup) {
	client := &http.Client{}
	temp := strconv.Itoa(i)
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/web/channel/multiple/list?channel_id="+temp+"&sort_type=hot&page_size=30", nil)
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

	for count := 0; count <= 30; count++ {
		num := strconv.Itoa(count)
		var res Result
		res.Channel_id = i
		res.Vname = gjson.Get(string(result), "data.list."+num+".name").String()
		res.View_count = gjson.Get(string(result), "data.list."+num+".view_count").String()
		res.Like_count = gjson.Get(string(result), "data.list."+num+".like_count").String()
		res.Author_name = gjson.Get(string(result), "data.list."+num+".author_name").String()
		res.Author_id = gjson.Get(string(result), "data.list."+num+".author_id").String()
		res.Bvid = gjson.Get(string(result), "data.list."+num+".bvid").String()
		if res.Vname != "" {
			fmt.Println(res)
			if err1 := DB.Table("content").Select("channel_id", "Vname", "view_count", "like_count", "author_name", "author_id", "bvid").Create(&res).Error; err1 != nil {
				fmt.Println(err1)
			}
		}
	}
	wg.Done()
}

func WorkerTwo(Url1 chan int) {
	//i := <-Url1
	for i := range Url1 {
		client := &http.Client{}
		temp := strconv.Itoa(i)
		req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/web/channel/multiple/list?channel_id="+temp+"&sort_type=hot&page_size=30", nil)
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

		for count := 0; count <= 30; count++ {
			num := strconv.Itoa(count)
			var res Result
			res.Channel_id = i
			res.Vname = gjson.Get(string(result), "data.list."+num+".name").String()
			res.View_count = gjson.Get(string(result), "data.list."+num+".view_count").String()
			res.Like_count = gjson.Get(string(result), "data.list."+num+".like_count").String()
			res.Author_name = gjson.Get(string(result), "data.list."+num+".author_name").String()
			res.Author_id = gjson.Get(string(result), "data.list."+num+".author_id").String()
			res.Bvid = gjson.Get(string(result), "data.list."+num+".bvid").String()
			if res.Vname != "" {
				fmt.Println(res)
				if err1 := DB.Table("content").Select("channel_id", "Vname", "view_count", "like_count", "author_name", "author_id", "bvid").Create(&res).Error; err1 != nil {
					fmt.Println(err1)
				}
			}
		}
	}
	//wg.Done()
}
