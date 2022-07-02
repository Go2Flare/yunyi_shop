package initialize

import "go.uber.org/zap"

// InitLogger 初始化zap log
func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
