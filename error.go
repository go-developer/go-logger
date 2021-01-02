// Package logger...
//
// Description : error 定义日志处理过程中的各种错误
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-02 2:44 下午
package logger

import "github.com/pkg/errors"

// CreateLogFileError 创建日志文件失败
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:55 下午 2021/1/2
func CreateLogFileError(err error, logFilePath string) error {
	return errors.Wrapf(err, "创建日志文件失败,日志文件路径 : %s", logFilePath)
}

// LogPathEmptyError 日志路径为空
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:03 下午 2021/1/2
func LogPathEmptyError() error {
	return errors.Wrap(errors.New("日志存储路径或者日志文件名为空"), "日志存储路径或者日志文件名为空")
}

// CustomTimeIntervalError 自定义日志切割时间间隔错误
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:11 下午 2021/1/2
func CustomTimeIntervalError() error {
	return errors.Wrap(errors.New("自定义时间间隔错误,必须是大于0的值"), "自定义时间间隔错误,必须是大于0的值")
}

// DealLogPathError 日志路径处理异常
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:31 下午 2021/1/2
func DealLogPathError(err error, logPath string) error {
	return errors.Wrapf(err, "日志路径检测处理异常, 日志路径 : %s", logPath)
}

// LogSplitTypeError 日志切割类型错误
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:50 下午 2021/1/2
func LogSplitTypeError(splitType TimeIntervalType) error {
	return errors.Wrapf(errors.New("日志切割时间类型错误"), "日志切割时间类型错误, 传入类型 : %v", splitType)
}

// CreateIOWriteError 创建日志实例失败
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:20 下午 2021/1/2
func CreateIOWriteError(err error) error {
	return errors.Wrapf(err, "创建日志实例失败")
}
