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
	"time"

	"go.uber.org/zap/zapcore"
)

// Test_Logger ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:50 下午 2021/1/2
func Test_Logger(t *testing.T) {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	c, err := NewRotateLogConfig("./logs", "test.log")
	if nil != err {
		panic(err)
	}
	l, err := NewLogger(zapcore.DebugLevel, encoder, c)
	if nil != err {
		panic(err)
	}
	l.Info("这是一条测试日志")
}
