package handler

import (
	"Bang/model"

	"github.com/gin-gonic/gin"
)

// @Summary "招募页面"
// @Description "获取十条招募信息并返回"
// @Tags requirements
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} []model.Requirement
// @Failure 404 "验证失败！"
// @Failure 401 "获取失败！"
// @Router /requirements [get]
func Requirements(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err1 := model.VerifyToken(token)
	if err1 != nil {
		c.JSON(404, gin.H{"message": "验证失败！"})
		return
	}
	temp, err := model.RequirePage()
	if err != nil {
		c.JSON(401, gin.H{"message": "获取失败！"})
		return
	}
	c.JSON(200, temp)
}

// @Summary "招募搜索"
// @Description "通过招募的名称进行搜索返回多组结果"
// @Tags requirements
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param temp body model.SearchContent true "temp"
// @Success 200 {object} []model.Requirement
// @Failure 404 "验证失败！"
// @Failure 401 "请重试！"
// @Failure 401 "查询失败！"
// @Router /requirements/search [post]
func Searcher2(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "验证失败！"})
		return
	}
	var temp model.SearchContent
	if err1 := c.BindJSON(&temp); err1 != nil {
		c.JSON(401, gin.H{"message": "请重试！"})
		return
	}
	Result, err2 := model.GetResult2(temp.Content)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "查询失败"})
		return
	}
	c.JSON(200, Result)
}

// @Summary "招募信息"
// @Description "通过招募id查看详细信息"
// @Tags requirements
// @Accept json
// @Produce json
// @Param token header string true "true"
// @Param require_id path string true "require_id"
// @Success 200 {object} model.Requirement model.UserInfo
// @Failure 404 "验证失败"
// @Failure 401 "获取失败"
// @Router /requirements/{require_id} [get]
func RequirePage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "验证失败！"})
		return
	}
	require_id := c.Param("require_id")
	temp, err1 := model.GetRequireInfo(require_id)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "打开失败！"})
		return
	}
	User, err2 := model.GetUserInfo(temp.UserId)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "获取失败"})
		return
	}
	c.JSON(200, gin.H{"RequireInfo": temp, "UserInfo": User})
}

/*
func Release(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "验证失败"})
		return
	}
	temp, err1 := model.GetUserInfo(id)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "查找失败！"})
	}
}
*/
