package router

import (
	"Bang/handler"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/registe", handler.Register)
	r.POST("/login", handler.Login)
	//我的页面
	v1 := r.Group("/user")
	{
		v1.GET("/", handler.HomePage)
		v1.GET("/myrelease", handler.MyRelease)
		v1.GET("/mystore", handler.Mystore)
		v1.GET("/myrelease/:require_id", handler.GetRequirement)
		v1.GET("/mystore/:require_id", handler.GetRequirement2)
	}
	v2 := r.Group("/match")
	{
		v2.GET("/", handler.Match)
		v2.GET("/:contest_id", handler.MatchInfo) //获取比赛的详细信息
		v2.POST("/search", handler.Searcher)      // 搜索的功能
	}
}
