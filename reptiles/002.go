package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"github.com/anaskhan96/soup"
)

func main() {
	requestUr1 := "http://www.liaoxuefeng.com/"
	rsp, err := http.Get(requestUr1)
	if err != nil {
		log.Println(err.Error())
		return
	}
//	take := string(rsp)

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	content := string(body)
//	fmt.Println(content)
	defer rsp.Body.Close()
	doc := soup.HTMLParse(content)
	subDocs := doc.FindAll("div","class","uk-margin")
	for _,subDoc := range subDocs {
		link := subDoc.Find("a")
		fmt.Println(link.Text())
	}
}
