package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/anaskhan96/soup"
)

func main() {
	request := "https://www.jianshu.com"
	rsp, err := http.Get(request)
	if err != nil {
		log.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	content := string(body)
	defer rsp.Body.Close()

	doc := soup.HTMLParse(content)
	subDocs := doc.FindAll("div","class","content")
	for _, subDoc := range subDocs {
		link := subDoc.Find("a")
		fmt.Println(link.Text())
	}
}

