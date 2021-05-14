package handler

import (
	"Bang/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.UserInfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误格式错误"})
		return
	}
	if a, _ := model.IfExistUserName(user.UserName); a != 1 {
		c.JSON(200, gin.H{
			"message": "该用户名已经存在，请重新创建。"})
		return
	}
	user_id := model.Register(user)
	if user_id == "" {
		c.JSON(400, gin.H{
			"message": "注册失败"})
		return
	}
	fmt.Println(user.UserId)
	c.JSON(200, gin.H{
		"user": user_id,
	})
}

func Login(c *gin.Context) {
	var User model.UserInfo
	if err := c.BindJSON(&User); err != nil {
		c.JSON(400, gin.H{
			"message": "输入的信息格式不符要求！"})
		return
	}
	if Id, err1 := model.TestInfo(User.UserName, User.User_password); err1 != nil {
		c.JSON(200, gin.H{
			"message": "密码或者昵称输入错误！"})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "登录成功",
			"token":   model.CreatToken(User.UserName, Id)})
	}
}
