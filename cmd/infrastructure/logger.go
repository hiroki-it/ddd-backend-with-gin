package infrastructure

import (
	"go.uber.org/zap"
)

type Logger struct {
	Log *zap.Logger
}

// NewLogger コンストラクタ
func NewLogger() *Logger {

	log, _ := zap.NewDevelopment()

	return &Logger{
		Log: log,
	}
}
