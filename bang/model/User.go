package model

import (
	"fmt"
)

type UserInfo struct {
	UserId        string `json:"user_id" gorm:"AUTO_INCREMENT"`
	UserName      string `json:"user_name" gorm:"column:user_name"`
	User_password string `json:"user_password" gorm:"user_password"`
	User_school   string `json:"user_school" gorm:"column:user_school"`
	User_college  string `json:"user_college" gorm:"column:user_college"`
	Real_Name     string `json:"really_name" gorm:"column:really_name"`
	User_picture  string `json:"user_picture" gorm:"column:user_picture"`
}

//防止用户名重复创建的函数
func IfExistUserName(UserName string) (int, error) {
	//println(UserName)
	var temp UserInfo
	if err := DB.Table("users").Where("user_name=?", UserName).Find(&temp).Error; err != nil {
		//log.Println(err)
		return 1, err
	}
	return 0, nil
}

//创建一个用户的函数
func Register(User UserInfo) string {
	//println(user.User_college)
	if err := DB.Table("users").Create(&User).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(User.UserId)
	return User.UserId
}

//验证登录信息的函数
func TestInfo(Name string, PassWord string) (string, error) {
	var temp UserInfo
	if err := DB.Table("users").Where("user_name=? And user_password=?", Name, PassWord).Find(&temp).Error; err != nil {
		fmt.Println(err)
		return "", err
	}
	return temp.UserId, nil
}
