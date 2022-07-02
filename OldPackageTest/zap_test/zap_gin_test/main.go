package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)
var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context){
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误,: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":"生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath": b64s,
	})
}

func InitBaseRouter(Router *gin.RouterGroup){
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", GetCaptcha)
	}
}

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"code":http.StatusOK,
			"success":true,
		})
	})

	//配置跨域
	//Router.Use(middlewares.Cors())

	//ApiGroup := Router.Group("/u/v1")
	return Router
}

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

func main(){
	InitLogger()
	router := Routers()
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	port := 8041
	zap.S().Infof("启动服务器端，端口：%v",port)
	if err := router.Run(fmt.Sprintf("%d", port)); err != nil{
		zap.S().Panic("启动服务器端失败，err：%v",err.Error())
	}
}
