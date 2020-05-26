package log

import (
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func Init() {

	logger, _ := zap.NewProduction()
	//defer logger.Sync() // flushes buffer, if any
	Sugar = logger.Sugar()

}
