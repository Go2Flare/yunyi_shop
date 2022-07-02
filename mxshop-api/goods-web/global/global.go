package global

import (
	ut "github.com/go-playground/universal-translator"
	"goods-web/config"
	"goods-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	Debug bool
)
