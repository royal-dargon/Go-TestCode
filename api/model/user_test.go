package model

import (
	"fmt"
	"testing"
)

func TestAdduser(t *testing.T) {
	fmt.Println("添加用户：")
	user := &User{}
	//调用添加用户的方法
	//user.AddUser()
	//user.AddUser2()
	us, _ := user.GetUsers()
	//遍历切片
	for k, v := range us {
		fmt.Printf("第%v个用户是%v：\n", k+1, v)
	}
}
