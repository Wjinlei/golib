package golib

import "testing"

func TestGetNowTime(t *testing.T) {
	t.Log(GetNowTime())
}

func TestFormatNowTime(t *testing.T) {
	t.Log(FormatNowTime("2006-01-02 15:04:05"))
}
