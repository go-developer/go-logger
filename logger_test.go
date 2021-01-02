// Package logger...
//
// Description : logger_test 单元测试
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-02 4:59 下午
package logger

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

// Test_Logger ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:50 下午 2021/1/2
func Test_Logger(t *testing.T) {
	encoder := GetEncoder()
	c, err := NewRotateLogConfig("./logs", "test.log", WithTimeIntervalType(TimeIntervalTypeMinute), WithMaxAge(120*time.Second))
	if nil != err {
		panic(err)
	}
	l, err := NewLogger(zapcore.InfoLevel, true, encoder, c)
	if nil != err {
		panic(err)
	}

	for {
		l.Info("这是一条测试日志", zap.Any("lala", "不限制类型"))
		l.Debug("这是一条测试日志", zap.Any("lala", "不限制类型"))
		time.Sleep(1 * time.Second)
	}
}

// Test_FormatJson 测试json格式化输出
//
// Author : go_developer@163.com<张德满>
//
// Date : 1:08 上午 2021/1/3
func Test_FormatJson(t *testing.T) {
	data := map[string]interface{}{
		"name": "zhangdeman",
		"age":  18,
	}
	str := FormatJson(data)
	fmt.Println(str)

	boolData := true
	str = FormatJson(boolData)
	fmt.Println(str)
}
