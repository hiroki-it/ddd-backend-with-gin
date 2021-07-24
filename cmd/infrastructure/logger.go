package infrastructure

import (
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.Logger
}

// NewLogger コンストラクタ
func NewLogger() *Logger {
	log, _ := zap.NewDevelopment()

	return &Logger{
		log: log,
	}
}

// Log logを返却します．
func (l *Logger) Log() *zap.Logger {
	return l.log
}
