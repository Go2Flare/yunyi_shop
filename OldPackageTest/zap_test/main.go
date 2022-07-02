package main

import (
	"go.uber.org/zap"
	"time"
)
var (
	url = "https://www.baidu.com"
	logger *zap.Logger
)

func init(){
	//生产环境打印json
	//logger, _ = zap.NewProduction()
	//开发环境打印日志
	//logger, _ = zap.NewDevelopment()

	//输出日志文件
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./zap_log_file/myproject.log",
		"stderr", //红色打印
		"stdout", //标准打印
	}
	logger, _ = cfg.Build()
}

func UseSugarLogger(){
	defer logger.Sync() // 最后刷新缓存
	sugar := logger.Sugar()//使用糖后，相当于简单打印，类似fmt
	//info write
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
		"key","value",
	)
	//info format
	sugar.Infof("Failed to fetch URL: %s", url)
}

func UseLogger(){
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		// 如果指明了类型，就不会触发go的反射机制，节省判断类型时间
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Info("Failed to fetch URL:", zap.String("",url))
}

func main() {
	UseSugarLogger()
	UseLogger()
}
