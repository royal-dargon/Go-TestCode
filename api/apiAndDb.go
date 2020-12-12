package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/royal-dargon/go1/api/utils"
)

//创建一个结构体
type users struct {
	username string
	password string
}

type User struct {
	ID       int
	Username string
	Password string
}

//注册系统
func register(w http.ResponseWriter, r *http.Request) {
	var a users
	fmt.Fprintln(w, "register")
	a.username = r.FormValue("username")
	a.password = r.FormValue("password")
	user := &User{}
	user.AddUser(a.username, a.password)
	fmt.Fprintln(w, "注册成功！")
}

//登录系统
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "您好，现在是登录界面")
	b := new(users)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	//验证账号是否正确
	user := &User{}
	use, _ := user.GetUsers()
	for _, k := range use {
		if k.Password != b.password || k.Username != b.username {
			fmt.Fprintf(w, "非法登录！")
		} else {
			fmt.Fprintln(w, "登录成功!")
			return
		}
	}
	return
}

//修改用户密码

func changepassword(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "您好，现在是修改密码阶段!")
	b := new(users)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	user := &User{}
	use, _ := user.GetUsers()
	for _, k := range use {
		if k.Password != b.password || k.Username != b.username {
			fmt.Fprintf(w, "非法修改！")
		} else {
			c := new(users)
			c.password = r.FormValue("changep")
			user := &User{}
			user.UpdatePassword(c.password, b.username, b.password)
			fmt.Fprintln(w, "修改成功!")
		}
	}
	return
}

//修改用户名
func changeusername(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "现在是修改用户名阶段")
	b := new(users)
	b.username = r.FormValue("username")
	b.password = r.FormValue("password")
	user := &User{}
	use, _ := user.GetUsers()
	for _, k := range use {
		if k.Password != b.password || k.Username != b.username {
			fmt.Fprintf(w, "非法修改！")
		} else {
			c := new(users)
			c.username = r.FormValue("changen")
			user := &User{}
			user.Updatename(c.username, b.username, b.password)
			fmt.Fprintln(w, "修改成功!")
		}
	}
	return
}

func main() {
	//开启一个又一个handler
	//注册
	http.HandleFunc("/register", register)
	//登录
	http.HandleFunc("/login", login)
	http.HandleFunc("/changepassword", changepassword)
	http.HandleFunc("/changeusername", changeusername)
	http.ListenAndServe(":10086", nil)
}

//插入demo
func (user *User) AddUser(userName string, passWord string) error {
	//写SQL语句
	sqlStr := "insert into users(username,password) values(?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, userName, passWord)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//GetUsers 获取数据库中所有的记录
func (user *User) GetUsers() ([]*User, error) {
	//写SQL语句
	sqlStr := "select id,username,password from users"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建user切片
	var users []*User
	for rows.Next() {
		//声明
		var id int
		var username string
		var password string
		err := rows.Scan(&id, &username, &password)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
		}
		users = append(users, u)
	}
	return users, nil
}

//updatepassword
func (user *User) UpdatePassword(password, username, password1 string) error {
	//写SQL语句
	sqlStr := "update users set password = ? where username = ? and password = ? "
	//执行
	_, err := utils.Db.Exec(sqlStr, password, username, password1)
	if err != nil {
		fmt.Println("执行出现异常:", err)
		return err
	}
	return nil
}

//更新用户名demo
func (user *User) Updatename(username, username1, password string) error {
	//写SQL语句
	sqlStr := "update users set username = ? where username = ? and password = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, username1, password)
	if err != nil {
		fmt.Println("执行出现错误:", err)
		return err
	}
	return nil
}
