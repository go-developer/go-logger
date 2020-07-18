package logger

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/go-util/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	traceLogIDField = "trace_log_id"
)

// SetTraceLogIDField 更换trace_id字段名
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 00:42:33
func SetTraceLogIDField(field string) {
	traceLogIDField = field
}

// NewWrapperLogger 获取日志实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 01:10:49
func NewWrapperLogger(cfg *LogConfig) *WrapperLogger {
	var (
		zapLogger *zap.Logger
		err       error
		l         *WrapperLogger
	)
	if zapLogger, err = NewZapLogger(cfg); nil != err {
		panic("获取zap日志实例失败, 失败原因 : " + err.Error())
	}
	l = &WrapperLogger{
		zapLogger: zapLogger,
	}
	return l
}

// WrapperLogger 包装 zap logger 和 gin 框架绑定
//
// Author : go_developer@163.com<张德满>
type WrapperLogger struct {
	zapLogger *zap.Logger
}

// Debug 记录debug级别日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 00:39:02
func (wl *WrapperLogger) Debug(ctx *gin.Context, msg string, fieldList ...zap.Field) {
	wl.zapLogger.Debug(msg, wl.getLoggerInfo(ctx, fieldList)...)
}

// Info 记录 info 级别日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 01:07:41
func (wl *WrapperLogger) Info(ctx *gin.Context, msg string, fieldList ...zap.Field) {
	wl.zapLogger.Info(msg, wl.getLoggerInfo(ctx, fieldList)...)
}

// Warn 记录 Warn 级别日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 01:07:41
func (wl *WrapperLogger) Warn(ctx *gin.Context, msg string, fieldList ...zap.Field) {
	wl.zapLogger.Warn(msg, wl.getLoggerInfo(ctx, fieldList)...)
}

// Error 记录 Error 级别日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 01:07:41
func (wl *WrapperLogger) Error(ctx *gin.Context, msg string, fieldList ...zap.Field) {
	wl.zapLogger.Error(msg, wl.getLoggerInfo(ctx, fieldList)...)
}

// Panic 记录 Panic 级别日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 01:07:41
func (wl *WrapperLogger) Panic(ctx *gin.Context, msg string, fieldList ...zap.Field) {
	wl.zapLogger.Panic(msg, wl.getLoggerInfo(ctx, fieldList)...)
}

// getLoggerInfo 获取日志记录的相关信息
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 00:44:03
func (wl *WrapperLogger) getLoggerInfo(ctx *gin.Context, fieldList []zap.Field) []zap.Field {
	var (
		traceID    string
		ctxTraceID interface{}
		exist      bool
	)
	if nil == fieldList {
		fieldList = make([]zapcore.Field, 0)
	}
	if ctxTraceID, exist = ctx.Get(traceLogIDField); !exist {
		traceID = util.ProjectUtil.GetTraceID()
		ctx.Set(traceLogIDField, traceID)
	} else {
		traceID = fmt.Sprintf("%v", ctxTraceID)
	}

	return append([]zap.Field{zap.String(traceLogIDField, traceID)}, fieldList...)
}
