package router

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"goods-web/middlewares"

	"goods-web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup){
	GoodsRouter := Router.Group("goods").Use(middlewares.Trace()).Use(otelgin.Middleware("goods-goods"))
	{
		GoodsRouter.GET("", goods.List) //商品列表
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New) //修改接口需要管理员权限
		GoodsRouter.GET("/:id", goods.Detail) //获取商品的详情
		GoodsRouter.DELETE("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete) //删除商品
		GoodsRouter.GET("/:id/stocks", goods.Stocks) //获取商品的库存
		GoodsRouter.PUT("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
		GoodsRouter.PATCH("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus)
		GoodsRouter.GET("/hotsearchs",middlewares.JWTAuth(), goods.HotSearch)
	}
}
