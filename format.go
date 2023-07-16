package log

import (
	"fmt"
	"time"
)

type FormatFunc func(level string, message string) string

var DefaultFormatFunc FormatFunc = TimestampFormatFunc

func TimestampFormatFunc(level string, message string) string {
	return fmt.Sprintf("[%s] [%s] %s", time.Now().Format(time.DateTime), level, message)
}

func SimpleFormatFunc(level string, message string) string {
	return fmt.Sprintf("[%s] %s", level, message)
}
