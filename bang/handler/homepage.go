package handler

import (
	"Bang/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	//var UserInfo model.UserInfo
	toekn := c.Request.Header.Get("token")
	id, err := model.VerifyToken(toekn)
	if err != nil {
		c.JSON(404, gin.H{"meaasge": "验证失败！"})
		return
	}
	fmt.Println(id)
	UserInfo, err1 := model.GetUserInfo(id)
	if err1 != nil {
		c.JSON(404, gin.H{"message": "没有查找到该用户的信息！"})
		return
	}
	c.JSON(200, UserInfo)
}
