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
	}
}
