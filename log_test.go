package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	log, _ := NewDefaultLoggerConfig("ginx", true, 1, "json", "./test.log", "./test.log")
	log.Error("测试日志", zap.Any("test", "这是一条message===="))
}
