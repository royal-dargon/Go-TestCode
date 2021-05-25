package handler

import (
	"Bang/model"

	"github.com/gin-gonic/gin"
)

// @Summary "我的页面"
// @Description "获取用户的基本信息"
// @Tags homepage
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.UserInfo
// @Failure 404 "验证失败"
// @Failure 401 "没有查找到该用户的信息！"
// @Router /user [get]
func HomePage(c *gin.Context) {
	//var UserInfo model.UserInfo
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"meaasge": "验证失败！"})
		return
	}
	//fmt.Println(id)
	UserInfo, err1 := model.GetUserInfo(id)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "没有查找到该用户的信息！"})
		return
	}
	c.JSON(200, UserInfo)
}

// @Summary "我的发布"
// @Description "查看我发布的招募信息"
// @Tags homepage
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} []model.Requirement
// @Failure 404 "验证失败"
// @Failure 401  "获取信息失败！"
// @Router /user/myrelease [get]
func MyRelease(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "验证失败！"})
		return
	}
	Requirement, err1 := model.GetMyRequire(id)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "获取信息失败！"})
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
		c.JSON(401, gin.H{"message": "没有找到信息！"})
		return
	}
	RequireInfo, err2 := model.GetMyStore(RequireId)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "没有获取到信息！"})
		return
	}
	c.JSON(200, RequireInfo)
}
