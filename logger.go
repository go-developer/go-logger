// Package logger ...
//
// File : logger.go
//
// Decs : 基于 zap 包装的 logger 库
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/27 14:06:03
package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewDefaultLoggerConfig 可以直接使用的一套基础配置实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/27 14:28:21
func NewDefaultLoggerConfig(appName string, develop bool, logLevel zapcore.Level, encodeing string, logFile string, errorLogFIle string) (*zap.Logger, error) {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel), //日志级别
		Development:       true,                           //是否是开发环境。如果是开发模式，对DPanicLevel进行堆栈跟踪
		Encoding:          encodeing,                      //编码类型，目前两种json 和 console【按照空格隔开】,常用json
		DisableCaller:     false,                          //禁止使用调用函数的文件名和行号来注释日志
		DisableStacktrace: true,                           //是否禁用堆栈跟踪捕获。默认对Warn级别以上和生产error级别以上的进行堆栈跟踪。
		EncoderConfig: zapcore.EncoderConfig{ //生成格式的一些配置--TODO
			TimeKey:        "time",  //输出时间的key名
			LevelKey:       "level", //输出日志级别的key名
			NameKey:        "log",
			CallerKey:      "caller",
			MessageKey:     "message", //输入信息的key名
			StacktraceKey:  "trace",
			LineEnding:     zapcore.DefaultLineEnding,      //每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
			EncodeLevel:    zapcore.LowercaseLevelEncoder,  //基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
			EncodeTime:     DefaultFormatEncodeTime,        //输出的时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder, //一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
			EncodeCaller:   zapcore.ShortCallerEncoder,     //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
		},
		OutputPaths:      []string{logFile},      //日志写入文件的地址
		ErrorOutputPaths: []string{errorLogFIle}, //将系统内的error记录到文件的地址
		InitialFields: map[string]interface{}{ //加入一些初始的字段数据
			"app": appName,
		},
	}
	return cfg.Build()
}

// DefaultFormatEncodeTime 默认的时间格式化函数
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/27 15:07:18
func DefaultFormatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%s.%d", t.Format("2006-01-02 15:04:05"), (t.UnixNano()-t.Unix()*1e9)/1e6))
}
