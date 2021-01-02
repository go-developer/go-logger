// Package logger...
//
// Description : logger 日志文件
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-02 5:04 下午
package logger

import (
	"io"
	"os"

	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// NewLogger 获取日志实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:05 下午 2021/1/2
func NewLogger(loggerLevel zapcore.Level, consoleOutput bool, encoder zapcore.Encoder, splitConfig *RotateLogConfig) (*zap.Logger, error) {
	loggerLevelDeal := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= loggerLevel
	})
	l := &Logger{
		splitConfig: splitConfig,
		encoder:     encoder,
	}
	var (
		err          error
		loggerWriter io.Writer
	)
	// 获取 日志实现
	if loggerWriter, err = l.getWriter(); nil != err {
		return nil, err
	}

	fileHandlerList := []zapcore.Core{
		zapcore.NewCore(encoder, zapcore.AddSync(loggerWriter), loggerLevelDeal),
	}

	// 设置控制台输出
	if consoleOutput {
		fileHandlerList = append(fileHandlerList, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), loggerLevelDeal))
	}

	// 最后创建具体的Logger
	core := zapcore.NewTee(fileHandlerList...)

	log := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	return log, nil
}

type Logger struct {
	splitConfig *RotateLogConfig
	encoder     zapcore.Encoder
}

// getWriter 获取日志实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:08 下午 2021/1/2
func (l *Logger) getWriter() (io.Writer, error) {
	option := make([]rotatelogs.Option, 0)
	option = append(option, rotatelogs.WithRotationTime(l.splitConfig.TimeInterval))
	if l.splitConfig.MaxAge > 0 {
		option = append(option, rotatelogs.WithMaxAge(l.splitConfig.MaxAge))
	}
	var (
		hook *rotatelogs.RotateLogs
		err  error
	)
	if hook, err = rotatelogs.New(l.splitConfig.FullLogFormat, option...); nil != err {
		return nil, CreateIOWriteError(err)
	}

	return hook, nil
}
