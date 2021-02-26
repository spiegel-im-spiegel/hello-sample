package hello

import (
	"go.uber.org/zap"
)

func Hello() {
	logger, _ := zap.NewDevelopment()
	logger.Info("Hello, World!")
}
