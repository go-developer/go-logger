// Package logger...
//
// Description : logger...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-03 12:58 上午
package logger

import (
	"bytes"
	"encoding/json"
)

// FormatJson 格式化输出json
//
// Author : go_developer@163.com<张德满>
//
// Date : 1:06 上午 2021/1/3
func FormatJson(src interface{}) string {
	byteData, _ := json.Marshal(src)

	var str bytes.Buffer
	_ = json.Indent(&str, byteData, "", "    ")
	return str.String()
}
