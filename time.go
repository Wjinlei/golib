package golib

import "time"

// GetNowTime 获取当前时间的字符串表示形式, 格式化为 2006-01-02 15:04:05
func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FormatNowTime 获取当前时间的字符串表示形式, 并自己传入格式化字符串
func FormatNowTime(format string) string {
	return time.Now().Format(format)
}
