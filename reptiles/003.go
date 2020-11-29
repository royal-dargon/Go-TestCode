package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/anaskhan96/soup"
)

func main() {
	request := "http://www.ccnu.edu.cn"
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
	subDocs := doc.FindAll()
	for _, subDoc := range subDocs {
		link := subDoc.Find("a")
		fmt.Println(link.Text())
	}
}

