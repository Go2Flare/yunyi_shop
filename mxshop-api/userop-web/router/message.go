package router

import (
	"github.com/gin-gonic/gin"
	"userop-web/api/message"
	"userop-web/middlewares"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middlewares.JWTAuth())
	{
		MessageRouter.GET("", message.List)          // 用户留言列表
		MessageRouter.POST("", message.New)       //新建用户留言
		MessageRouter.DELETE("/:id", message.Delete)
	}
}