package model

import (
	"fmt"

	"github.com/royal-dargon/go1/api/utils"
)

type User struct {
	ID       int
	Username string
	Password string
}

//add user 第一招
func (user *User) AddUser() error {
	//写SQL语句
	sqlStr := "insert into users(username,password) values(?,?)"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
	}
	//执行
	_, err2 := inStmt.Exec("汤先宁", "666")
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err
	}
	return nil
}

//add user 第二招
func (user *User) AddUser2() error {
	//写SQL语句
	sqlStr := "insert into users(username,password) values(?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, "张旷", "123")
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//GetUserId 根据用户的ID从表中查询信息
func (user *User) GetUserId() (*User, error) {
	//写SQL语句
	sqlStr := "select username,password from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, user.ID)
	//声明
	var id int
	var username string
	var password string
	err := row.Scan(&username, &password)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
	}
	return u, nil
}

//GetUsers 获取数据库中所有的记录
func (user *User) GetUsers() ([]*User, error) {
	//写SQL语句
	sqlStr := "select id,username,password from users"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	utils.Db.Query(sqlStr)
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
func (user *User) updatePassword {
	//写SQL语句
	sqlStr := "update users set password = '' where username = '' "
	//执行
	_, err := utils.Db.
}
