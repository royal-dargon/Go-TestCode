package main

import (
	"fmt"
	"net/http"
)

//创建一个结构体
type user struct {
	username string
	password string
}

//稀里糊涂的混子处理
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你好！", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求后的查询字符是:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息：", r.Header)
}

var a user

//注册系统
func register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "register")
	a.username = r.FormValue("username")
	a.password = r.FormValue("password")
	fmt.Fprintln(w, "注册成功！")
}

//登录系统
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "您好，现在是登录界面")
	b := new(user)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	if a.username == b.username && a.password == b.password {
		fmt.Fprintln(w, "登录成功！")
	} else {
		fmt.Fprintf(w, "非法登录！")
	}
}

//修改用户密码

func changepassword(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "您好，现在是修改密码阶段!")
	b := new(user)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	if a.username == b.username && a.password == b.password {
		a.password = r.FormValue("chaangep")
		a.username = b.username
		fmt.Fprintln(w, "修改成功!")
	} else {
		fmt.Fprintln(w, "非法修改！")
	}
}

//修改用户名
func changeusername(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "现在是修改用户名阶段")
	b := new(user)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	if a.username == b.username && a.password == b.password {
		a.username = r.FormValue("changename")
		fmt.Fprintln(w, "修改成功！")
	} else {
		fmt.Fprintln(w, "非法修改！")
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println(a)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/changepassword", changepassword)
	http.HandleFunc("/changeusername", changeusername)
	http.ListenAndServe(":10086", nil)
}
