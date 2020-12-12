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

var userDate = make(map[string]user)

//稀里糊涂的混子处理
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你好！", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求后的查询字符是:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息：", r.Header)
}

//注册系统
func register(w http.ResponseWriter, r *http.Request) {
	var a user
	fmt.Fprintln(w, "register")
	a.username = r.FormValue("username")
	a.password = r.FormValue("password")
	userDate[a.username] = a
	fmt.Fprintln(w, "注册成功！")
}

//登录系统
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "您好，现在是登录界面")
	b := new(user)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	//验证账号是否正确
	use, ok := userDate[b.username]
	if !ok {
		fmt.Fprintln(w, "非法登录！")
		return
	}
	if use.password != b.password {
		fmt.Fprintf(w, "非法登录！")
		return
	} else {
		fmt.Fprintln(w, "登录成功!")
	}
}

//修改用户密码

func changepassword(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "您好，现在是修改密码阶段!")
	b := new(user)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	use, ok := userDate[b.username]
	if !ok {
		fmt.Fprintln(w, "非法修改！")
		return
	}

	if use.password != b.password {
		fmt.Fprintln(w, "非法修改！")
		return
	} else {
		use.password = r.FormValue("chaangep")
		userDate[b.username] = use
		fmt.Fprintln(w, "修改成功!")
	}
}

//修改用户名
func changeusername(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "现在是修改用户名阶段")
	b := new(user)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	use, ok := userDate[b.username]
	if !ok {
		fmt.Fprintln(w, "非法修改!")
		return
	}

	if use.password != b.password {
		fmt.Fprintln(w, "修改失败!")
		return
	} else {
		use.username = r.FormValue("changename")
		userDate[use.username] = use
		fmt.Fprintln(w, "修改成功！")
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/changepassword", changepassword)
	http.HandleFunc("/changeusername", changeusername)
	http.ListenAndServe(":10086", nil)
}
