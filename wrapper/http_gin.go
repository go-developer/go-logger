// Package wrapper...
//
// Description : http_gin 使用gin框架时的，记录日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-03 3:43 下午
package wrapper

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewGinWrapperLogger 使用gin框架记录日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:45 下午 2021/1/3
func NewGinWrapperLogger(loggerLevel zapcore.Level, consoleOutput bool, encoder zapcore.Encoder, splitConfig *logger.RotateLogConfig, extractFieldList []string) (*GinWrapper, error) {
	var (
		err error
		l   *zap.Logger
	)
	if l, err = logger.NewLogger(loggerLevel, consoleOutput, encoder, splitConfig); nil != err {
		return nil, err
	}

	return &GinWrapper{
		loggerInstance:   l,
		extractFieldList: extractFieldList,
	}, nil
}

// GinWrapper 包装gin实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:59 下午 2021/1/3
type GinWrapper struct {
	loggerInstance   *zap.Logger  // zap 的日志实例
	extractFieldList []string     // 从gin中抽取的字段
	ginCtx           *gin.Context // gin 实例
}

// GetLogger 为每一次请求生成不同的日志实例,包含独立的gin上下文
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:02 下午 2021/1/3
func (gw *GinWrapper) GetLogger(ginCtx *gin.Context) *GinWrapper {
	return &GinWrapper{
		loggerInstance:   gw.loggerInstance,
		extractFieldList: gw.extractFieldList,
		ginCtx:           ginCtx,
	}
}

// formatFieldList 格式化日志field列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:13 下午 2021/1/3
func (gw *GinWrapper) formatFieldList(inputFieldList []zap.Field) []zap.Field {
	if nil == inputFieldList {
		inputFieldList = make([]zap.Field, 0)
	}
	if nil != gw.ginCtx {
		// 自动扩充抽取字段,字段不存在的话,忽略掉
		for _, extractField := range gw.extractFieldList {
			if v, exist := gw.ginCtx.Get(extractField); exist {
				byteData, _ := json.Marshal(v)
				inputFieldList = append(inputFieldList, zap.String(extractField, string(byteData)))
			}
		}
	}
	return inputFieldList
}

// Debug 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:14 下午 2021/1/3
func (gw *GinWrapper) Debug(msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(field)
	gw.loggerInstance.Debug(msg, fieldList...)
}

// Info 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:28 下午 2021/1/3
func (gw *GinWrapper) Info(msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(field)
	gw.loggerInstance.Info(msg, fieldList...)
}

// Warn 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:29 下午 2021/1/3
func (gw *GinWrapper) Warn(msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(field)
	gw.loggerInstance.Warn(msg, fieldList...)
}

// Error 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:29 下午 2021/1/3
func (gw *GinWrapper) Error(msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(field)
	gw.loggerInstance.Error(msg, fieldList...)
}

// Panic 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:29 下午 2021/1/3
func (gw *GinWrapper) Panic(msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(field)
	gw.loggerInstance.Panic(msg, fieldList...)
}

// DPanic 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:30 下午 2021/1/3
func (gw *GinWrapper) DPanic(msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(field)
	gw.loggerInstance.DPanic(msg, fieldList...)
}

// GetZapLoggerInstance 获取zap日志实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021/01/03 22:56:47
func (gw *GinWrapper) GetZapLoggerInstance() *zap.Logger {
	return gw.loggerInstance
}
