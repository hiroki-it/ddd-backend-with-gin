package infrastructure

import "go.uber.org/zap"

// NewLogger コンストラクタ
func NewLogger() (*zap.Logger, error) {
	log, err := zap.NewDevelopment()
	return log, err
}
