package handler

import (
	"Bang/model"

	"github.com/gin-gonic/gin"
)

// @Summary "大赛初始页面"
// @Description "获取大赛页面的信息"
// @Tags contest
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Contest
// @Failure 404 "验证失败"
// @Failure 401 "获取信息失败"
// @Router /match [get]
func Match(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "验证失败！"})
		return
	}
	MatchInfo, err1 := model.GetMatchInfo()
	if err1 != nil {
		c.JSON(401, gin.H{"message": "获取页面信息失败！"})
		return
	}
	c.JSON(200, MatchInfo)
}

// @Summary "获取信息"
// @Description "获取比赛的信息"
// @Tags contest
// @Accept json
// @Produce json
// @Param contest_id path string true "contest_id"
// @Success 200 {object} model.Contest
// @Failure 401 "获取信息失败"
// @Router /match/{contest_id} [get]
func MatchInfo(c *gin.Context) {
	id := c.Param("content_id")
	ContestInfo, err := model.GetContestInfo(id)
	if err != nil {
		c.JSON(401, gin.H{"messgae": "获取信息失败"})
		return
	}
	c.JSON(200, ContestInfo)
}

// @Summary "搜索比赛"
// @Description "通过比赛名称对比赛进行搜索"
// @Tags contest
// @Accept json
// @Produce json
// @Param temp body model.SearchContent true "temp"
// @Success 200 {object} model.Contest
// @Failure 404 "请重试"
// @Failure 401 "查询失败"
// @Router /match/search [post]
func Searcher(c *gin.Context) {
	var temp model.SearchContent
	if err := c.BindJSON(&temp); err != nil {
		c.JSON(404, gin.H{"message": "请重试"})
		return
	}
	result, err1 := model.SearchOK(temp.Content)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "查询失败"})
		return
	}
	c.JSON(200, result)
}
