package router

import (
	"github.com/gin-gonic/gin"
	"userop-web/api/address"
	"userop-web/middlewares"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuth(), address.List)          // 用户地址列表
		AddressRouter.DELETE("/:id", middlewares.JWTAuth(), address.Delete) // 删除用户地址
		AddressRouter.POST("", middlewares.JWTAuth(), address.New)       //新建用户地址
		AddressRouter.PUT("/:id",middlewares.JWTAuth(), address.Update) //修改用户地址
	}
}