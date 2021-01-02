// Package logger...
//
// Description : logger_test 单元测试
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-02 4:59 下午
package logger

import (
	"testing"

	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

// Test_Logger ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:50 下午 2021/1/2
func Test_Logger(t *testing.T) {
	encoder := GetEncoder(false)
	c, err := NewRotateLogConfig("./logs", "test.log")
	if nil != err {
		panic(err)
	}
	l, err := NewLogger(zapcore.DebugLevel, encoder, c)
	if nil != err {
		panic(err)
	}
	l.Info("这是一条测试日志", zap.Any("lala", "不限制类型"))
}
