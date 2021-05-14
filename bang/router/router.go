package router

import (
	"Bang/handler"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/registe", handler.Register)
	r.POST("/login", handler.Login)
}
