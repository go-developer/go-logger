// Package logger...
//
// Description : config 日志配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-02 3:07 下午
package logger

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

// TimeIntervalType 日志时间间隔类型
type TimeIntervalType uint

const (
	// TimeIntervalTypeMinute 按分钟切割
	TimeIntervalTypeMinute = TimeIntervalType(0)
	// TimeIntervalTypeHour 按小时切割
	TimeIntervalTypeHour = TimeIntervalType(1)
	// TimeIntervalTypeDay 按天切割
	TimeIntervalTypeDay = TimeIntervalType(2)
	// TimeIntervalTypeMonth 按月切割
	TimeIntervalTypeMonth = TimeIntervalType(3)
	// TimeIntervalTypeYear 按年切割
	TimeIntervalTypeYear = TimeIntervalType(4)
)

const (
	// DefaultDivisionChar 默认的时间格式分隔符
	DefaultDivisionChar = "-"
)

// RotateLogConfig 日志切割的配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:08 下午 2021/1/2
type RotateLogConfig struct {
	TimeIntervalType TimeIntervalType // 日志切割的时间间隔类型 0 - 小时 1 - 天 2 - 月 3 - 年
	TimeInterval     time.Duration    // 日志切割的时间间隔
	LogPath          string           // 存储日志的路径
	LogFileName      string           // 日志文件名
	DivisionChar     string           // 日志文件拼时间分隔符
	FullLogFormat    string           // 完整的日志格式
	MaxAge           time.Duration    // 日志最长保存时间
}

// SetRotateLogConfigOption 设置日志切割的选项
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:13 下午 2021/1/2
type SetRotateLogConfigFunc func(rlc *RotateLogConfig)

// WithTimeIntervalType 设置日志切割时间间隔
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:34 下午 2021/1/2
func WithTimeIntervalType(timeIntervalType TimeIntervalType) SetRotateLogConfigFunc {
	return func(rlc *RotateLogConfig) {
		rlc.TimeIntervalType = timeIntervalType
	}
}

// WithDivisionChar 设置分隔符
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:49 下午 2021/1/2
func WithDivisionChar(divisionChar string) SetRotateLogConfigFunc {
	return func(rlc *RotateLogConfig) {
		rlc.DivisionChar = divisionChar
	}
}

// WithMaxAge 设置日志保存时间
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:03 下午 2021/1/2
func WithMaxAge(maxAge time.Duration) SetRotateLogConfigFunc {
	return func(rlc *RotateLogConfig) {
		rlc.MaxAge = maxAge
	}
}

// NewRotateLogConfig 生成日志切割的配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:53 下午 2021/1/2
func NewRotateLogConfig(logPath string, logFile string, option ...SetRotateLogConfigFunc) (*RotateLogConfig, error) {
	if len(logPath) == 0 || len(logFile) == 0 {
		return nil, LogPathEmptyError()
	}
	c := &RotateLogConfig{
		TimeIntervalType: 0,
		LogPath:          logPath,
		LogFileName:      logFile,
		DivisionChar:     "",
	}

	for _, o := range option {
		o(c)
	}

	if err := formatConfig(c); nil != err {
		return nil, err
	}

	return c, nil
}

// formatConfig 格式化配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:23 下午 2021/1/2
func formatConfig(c *RotateLogConfig) error {

	if len(c.DivisionChar) == 0 {
		c.DivisionChar = DefaultDivisionChar
	}
	// 格式化路径
	logPathByte := []byte(c.LogPath)
	if string(logPathByte[len(logPathByte)-1]) != "/" {
		c.LogPath = c.LogPath + "/"
	}
	// 检测路径是否存在,不存在自动创建
	if _, err := os.Stat(c.LogPath); nil != err {
		if !os.IsNotExist(err) {
			// 异常不是路径不存在,抛异常
			return DealLogPathError(err, c.LogPath)
		}
		if err := os.Mkdir(c.LogPath, os.ModePerm); nil != err {
			return DealLogPathError(err, "创建日志目录失败")
		}
	}

	// 生成格式化日志全路径
	switch c.TimeIntervalType {
	case TimeIntervalTypeMinute:
		c.TimeInterval = time.Minute
		c.FullLogFormat = c.LogPath + "%Y" + c.DivisionChar + "%m" + c.DivisionChar + "%d" + c.DivisionChar + "%H" + c.DivisionChar + "%M" + c.DivisionChar + c.LogFileName
	case TimeIntervalTypeHour:
		c.TimeInterval = time.Hour
		c.FullLogFormat = c.LogPath + "%Y" + c.DivisionChar + "%m" + c.DivisionChar + "%d" + c.DivisionChar + "%H" + c.DivisionChar + c.LogFileName
	case TimeIntervalTypeDay:
		c.TimeInterval = time.Hour * 24
		c.FullLogFormat = c.LogPath + "%Y" + c.DivisionChar + "%m" + c.DivisionChar + "%d" + c.DivisionChar + c.LogFileName
	case TimeIntervalTypeMonth:
		c.TimeInterval = time.Hour * 24 * 30
		c.FullLogFormat = c.LogPath + "%Y" + c.DivisionChar + "%m" + c.DivisionChar + c.LogFileName
	case TimeIntervalTypeYear:
		c.TimeInterval = time.Hour * 24 * 365
		c.FullLogFormat = c.LogPath + "%Y" + c.DivisionChar + c.LogFileName
	default:
		return LogSplitTypeError(c.TimeIntervalType)
	}

	return nil
}

// ============== 以下为zap相关配置

const (
	// defaultMessageKey 默认的message key
	defaultMessageKey = "message"
	// defaultLevelKey 默认的level
	defaultLevelKey = "level"
	// defaultTimeKey 默认时间key
	defaultTimeKey = "time"
	// defaultCallerKey 默认的文件key
	defaultCallerKey = "file"
	// defaultUserShortCaller 是否使用短的文件调用格式
	defaultUseShortCaller = true
	// defaultUseJsonFormat 日志默认使用json格式
	defaultUseJsonFormat = true
)

// defaultTimeEncoder 默认的时间处理
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:53 下午 2021/1/2
func defaultTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// defaultEncodeDuration 默认的原始时间处理
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:56 下午 2021/1/2
func defaultEncodeDuration(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(int64(d) / 1000000)
}

// OptionLogger 日志配置的选项
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:41 下午 2021/1/2
type OptionLogger struct {
	UseJsonFormat  bool                    // 日志使用json格式
	MessageKey     string                  // message 字段
	LevelKey       string                  // level 字段
	TimeKey        string                  // 时间字段
	CallerKey      string                  // 记录日志的文件的代码行数
	UseShortCaller bool                    // 使用短的调用文件格式
	TimeEncoder    zapcore.TimeEncoder     // 格式化时间的函数
	EncodeDuration zapcore.DurationEncoder // 原始时间信息
}

// 设置日志配置
type SetLoggerOptionFunc func(o *OptionLogger)

// WithUseJsonFormat 日志是否使用json格式数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:30 上午 2021/1/3
func WithUseJsonFormat(isJsonFormat bool) SetLoggerOptionFunc {
	return func(o *OptionLogger) {
		o.UseJsonFormat = isJsonFormat
	}
}

// WithMessageKey 使用message key
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:32 上午 2021/1/3
func WithMessageKey(messageKey string) SetLoggerOptionFunc {
	return func(o *OptionLogger) {
		messageKey = strings.Trim(messageKey, " ")
		if len(messageKey) == 0 {
			return
		}
		o.MessageKey = messageKey
	}
}

// WithLevelKey 设置level key
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:33 上午 2021/1/3
func WithLevelKey(levelKey string) SetLoggerOptionFunc {
	return func(o *OptionLogger) {
		levelKey = strings.Trim(levelKey, " ")
		if len(levelKey) == 0 {
			return
		}
		o.LevelKey = levelKey
	}
}

// WithTimeKey 设置time key ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:34 上午 2021/1/3
func WithTimeKey(timeKey string) SetLoggerOptionFunc {
	return func(o *OptionLogger) {
		timeKey = strings.Trim(timeKey, " ")
		if len(timeKey) == 0 {
			return
		}
		o.TimeKey = timeKey
	}
}

// WithCallerKey 设置caller key
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:37 上午 2021/1/3
func WithCallerKey(callerKey string) SetLoggerOptionFunc {
	return func(o *OptionLogger) {
		callerKey = strings.Trim(callerKey, " ")
		if len(callerKey) == 0 {
			return
		}
		o.CallerKey = callerKey
	}
}

// WithShortCaller 是否使用短caller格式
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:39 上午 2021/1/3
func WithShortCaller(useShortCaller bool) SetLoggerOptionFunc {
	return func(o *OptionLogger) {
		o.UseShortCaller = useShortCaller
	}
}

// WithTimeEncoder 设置格式化时间方法
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:41 上午 2021/1/3
func WithTimeEncoder(encoder zapcore.TimeEncoder) SetLoggerOptionFunc {
	return func(o *OptionLogger) {
		if nil == encoder {
			return
		}
		o.TimeEncoder = encoder
	}
}

// WithEncodeDuration 原始时间
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:42 上午 2021/1/3
func WithEncodeDuration(encoder zapcore.DurationEncoder) SetLoggerOptionFunc {
	return func(o *OptionLogger) {
		if nil == encoder {
			return
		}
		o.EncodeDuration = encoder
	}
}

// GetEncoder 获取空中台输出的encoder
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:24 下午 2021/1/2
func GetEncoder(option ...SetLoggerOptionFunc) zapcore.Encoder {
	ol := &OptionLogger{
		UseJsonFormat:  defaultUseJsonFormat,
		MessageKey:     defaultMessageKey,
		LevelKey:       defaultLevelKey,
		TimeKey:        defaultTimeKey,
		TimeEncoder:    defaultTimeEncoder,
		CallerKey:      defaultCallerKey,
		EncodeDuration: defaultEncodeDuration,
		UseShortCaller: defaultUseShortCaller,
	}
	for _, o := range option {
		o(ol)
	}
	ec := zapcore.EncoderConfig{
		MessageKey:     ol.MessageKey,
		LevelKey:       ol.LevelKey,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		TimeKey:        ol.TimeKey,
		EncodeTime:     ol.TimeEncoder,
		CallerKey:      ol.CallerKey,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: ol.EncodeDuration,
	}
	if !ol.UseShortCaller {
		ec.EncodeCaller = zapcore.FullCallerEncoder
	}
	if !ol.UseJsonFormat {
		return zapcore.NewConsoleEncoder(ec)
	}
	return zapcore.NewJSONEncoder(ec)
}
