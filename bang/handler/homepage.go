package handler

import (
	"Bang/model"

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
	//fmt.Println(id)
	UserInfo, err1 := model.GetUserInfo(id)
	if err1 != nil {
		c.JSON(404, gin.H{"message": "没有查找到该用户的信息！"})
		return
	}
	c.JSON(200, UserInfo)
}

func MyRelease(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "验证失败！"})
		return
	}
	Requirement, err1 := model.GetMyRequire(id)
	if err1 != nil {
		c.JSON(404, gin.H{"message": "获取信息失败！"})
		return
	}
	c.JSON(200, Requirement)
}

func Mystore(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "验证失败"})
		return
	}
	RequireId, err1 := model.GetMyStoreId(id)
	if err1 != nil {
		c.JSON(404, gin.H{"message": "没有找到信息！"})
		return
	}
	RequireInfo, err2 := model.GetMyStore(RequireId)
	if err2 != nil {
		c.JSON(404, gin.H{"message": "没有获取到信息！"})
		return
	}
	c.JSON(200, RequireInfo)
}
