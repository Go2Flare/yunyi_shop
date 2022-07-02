package global

import (
	ut "github.com/go-playground/universal-translator"
	"user-web/config"
	"user-web/proto"
)

var (
	//错误翻译器
	Trans ut.Translator

	//服务器 配置
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	//nacos 配置
	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	//grpc proto
	UserSrvClient proto.UserClient

	//默认生产环境
	Debug bool
)
