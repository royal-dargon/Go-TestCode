package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/howeyc/gopass"
)

func main() {
	requestUrl := "https://accounts.douban.com/j/mobile/login/basic"

	// 输入账号和密码
	var name string
	fmt.Print("name：")
	_, _ = fmt.Scanln(&name)
	fmt.Print("输入密码：")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err)
	}

	data := url.Values{}
	data.Set("name", name)
	data.Set("password", string(password))
	data.Set("remember", "false")
	data.Set("ck", "")
	data.Set("ticket", "")

	payload := strings.NewReader(data.Encode())

	request, err := http.NewRequest("POST", requestUrl, payload)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Origin", "http://work.muxi-tech.xyz")
	request.Header.Add("Referer", "http://work.muxi-tech.xyz/landing/?username=baron&token=de6bee888238beb706f8bd19d23d37240a1facb6e03cfb0d1dc1ba6a1fabdd6e102c0152de63a272ce8dfe5b4b2989b015ad80cff72ca6e7d36ff814a3cbd13144&id=345")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)

	fmt.Println(string(body))
}
