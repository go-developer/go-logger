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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger 可以直接使用的一套基础配置实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/27 14:28:21
func NewZapLogger(cfg *LogConfig) (*zap.Logger, error) {
	zapCfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(cfg.LogLevel), //日志级别
		Development:       true,                               //是否是开发环境。如果是开发模式，对DPanicLevel进行堆栈跟踪
		Encoding:          cfg.Encoding,                       //编码类型，目前两种json 和 console【按照空格隔开】,常用json
		DisableCaller:     false,                              //禁止使用调用函数的文件名和行号来注释日志
		DisableStacktrace: true,                               //是否禁用堆栈跟踪捕获。默认对Warn级别以上和生产error级别以上的进行堆栈跟踪。
		EncoderConfig: zapcore.EncoderConfig{ //生成格式的一些配置--TODO
			TimeKey:        cfg.LogKeyInfo.Time,       //输出时间的key名
			LevelKey:       cfg.LogKeyInfo.Level,      //输出日志级别的key名
			NameKey:        cfg.LogKeyInfo.Name,       // name
			CallerKey:      cfg.LogKeyInfo.Caller,     // caller
			MessageKey:     cfg.LogKeyInfo.Message,    //输入信息的key名
			StacktraceKey:  cfg.LogKeyInfo.Stacktrace, //stack
			LineEnding:     cfg.LineEnding,            //每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
			EncodeLevel:    cfg.EncodeLevel,           //基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
			EncodeTime:     cfg.EncodeTime,            //输出的时间格式
			EncodeDuration: cfg.EncodeDuration,        //一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
			EncodeCaller:   cfg.EncodeCaller,          //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
		},
		OutputPaths: []string{cfg.LogFile}, //日志写入文件的地址
		//ErrorOutputPaths: []string{errorLogFIle}, //将系统内的error记录到文件的地址
		InitialFields: map[string]interface{}{ //加入一些初始的字段数据
			"app": cfg.AppName,
		},
	}
	return zapCfg.Build()
}
