// Package logger ...
//
// Author: go_developer@163.com<张德满>
//
// Description: 日志相关配置及相关默认值配置
//
// File: config.go
//
// Version: 1.0.0
//
// Date: 2020/07/19 01:42:53
package logger

import (
	"time"

	"github.com/go-developer/go-util/util"
	"go.uber.org/zap/zapcore"
)

//定义相关key的默认值
const (
	// DefaultTimeKey 默认时间key
	DefaultTimeKey = "time"
	// DefaultLevelKey 默认level key
	DefaultLevelKey = "level"
	// DefaultLogKey 默认 log key
	DefaultNameKey = "log"
	// 默认 caller key
	DefaultCallerKey = "caller"
	// 默认 message key
	DefaultMessageKey = "message"
	// 默认 stacktrace key
	DefaultStacktraceKey = "trace"
)

// LogConfig 日志的配置
//
// Author : go_developer@163.com<张德满>
type LogConfig struct {
	AppName        string                  `json:"app_name"`        //应用名称
	Develop        bool                    `json:"develop"`         //是否开发模式
	LogLevel       zapcore.Level           `json:"log_level"`       //日志级别
	Encoding       string                  `json:"encoding"`        //编码
	LogFile        string                  `json:"log_file"`        //日志文件
	LogKeyInfo     *KeyInfo                `json:"log_key_info"`    //日志key相关信息
	EncodeTime     zapcore.TimeEncoder     `json:"encode_time"`     //时间戳的各式还函数
	LineEnding     string                  `json:"line_ending"`     //换行符
	EncodeLevel    zapcore.LevelEncoder    `json:"encode_level"`    //level 编码器
	EncodeDuration zapcore.DurationEncoder `json:"encode_duration"` //duration 编码
	EncodeCaller   zapcore.CallerEncoder   `json:"caller_encoder"`  //调用的编码
}

// KeyInfo 定义相关的 key 信息
//
// Author : go_developer@163.com<张德满>
type KeyInfo struct {
	Time       string `json:"time"`       //时间戳的字段名
	Level      string `json:"level"`      //输出日志级别的key
	Caller     string `json:"caller"`     //调用信息key
	Message    string `json:"message"`    //信息字段的key
	Stacktrace string `json:"stacktrace"` //堆栈key
	Name       string `json:"name"`       // name key
}

// DefaultFormatEncodeTime 默认的时间格式化函数
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/27 15:07:18
func DefaultFormatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(util.TimeUtil.GetFormatCurrentNanoTime(time.Now().UnixNano()))
}

// appName string, develop bool, logLevel zapcore.Level, encodeing string, logFile string, errorLogFIle string

// BuildLogConfig 构建日志配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 01:52:36
func BuildLogConfig(appName string, develop bool, logLevel zapcore.Level, encoding string, logFile string, keyInfo *KeyInfo) *LogConfig {
	keyInfo = formatKeyInfo(keyInfo)
	return &LogConfig{
		AppName:        appName,
		Develop:        develop,
		LogLevel:       logLevel,
		Encoding:       encoding,
		LogFile:        logFile,
		LogKeyInfo:     keyInfo,
		EncodeTime:     DefaultFormatEncodeTime,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// formatKeyInfo 格式化时间的方法
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 01:59:25
func formatKeyInfo(keyInfo *KeyInfo) *KeyInfo {
	if nil == keyInfo {
		keyInfo = &KeyInfo{}
	}
	if len(keyInfo.Caller) == 0 {
		keyInfo.Caller = DefaultCallerKey
	}
	if len(keyInfo.Level) == 0 {
		keyInfo.Level = DefaultLevelKey
	}
	if len(keyInfo.Message) == 0 {
		keyInfo.Message = DefaultMessageKey
	}
	if len(keyInfo.Stacktrace) == 0 {
		keyInfo.Message = DefaultStacktraceKey
	}
	if len(keyInfo.Time) == 0 {
		keyInfo.Time = DefaultTimeKey
	}
	if len(keyInfo.Name) == 0 {
		keyInfo.Name = DefaultNameKey
	}
	return keyInfo
}
