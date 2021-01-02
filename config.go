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
	"time"
)

// TimeIntervalType 日志时间间隔类型
type TimeIntervalType uint

const (
	// TimeIntervalTypeHour 按小时切割
	TimeIntervalTypeHour = TimeIntervalType(0)
	// TimeIntervalTypeDay 按天切割
	TimeIntervalTypeDay = TimeIntervalType(1)
	// TimeIntervalTypeMonth 按月切割
	TimeIntervalTypeMonth = TimeIntervalType(2)
	// TimeIntervalTypeYear 按年切割
	TimeIntervalTypeYear = TimeIntervalType(3)
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
