package log

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type SimpleLogger struct {
	output     io.Writer
	errOutput  io.Writer
	formatFunc FormatFunc
	level      Level
	errDone    bool

	colorLayer ColorLayer

	Logger
}

func NewSimpleLogger() *SimpleLogger {
	l := &SimpleLogger{
		output:     os.Stdout,
		formatFunc: DefaultFormatFunc,
		level:      LevelDebug,
	}
	l.errOutput = l.output
	l.Logger = newModelToLogger(l)
	return l
}

func (l *SimpleLogger) SetOutput(output io.Writer) {
	l.output = output
}

func (l *SimpleLogger) SetErrOutput(errOutput io.Writer) {
	l.errOutput = errOutput
}

func (l *SimpleLogger) SetFormatFunc(formatFunc FormatFunc) {
	l.formatFunc = formatFunc
}

func (l *SimpleLogger) SetLevel(level Level) *SimpleLogger {
	l.level = level
	return l
}

func (l *SimpleLogger) SetErrDone(errDone bool) {
	l.errDone = errDone
}

func (l *SimpleLogger) InitColor() {
	if l.colorLayer == nil {
		l.colorLayer = newColorLayer()
	}
}

func (l *SimpleLogger) ColorLayer() ColorLayer {
	return l.colorLayer
}

func (l *SimpleLogger) sprint(level Level, a ...any) string {
	if l.level.level() > level.level() {
		return ""
	}
	message := fmt.Sprint(a...)
	levelStr := level.String()
	if cl := l.ColorLayer(); cl != nil {
		levelStr = cl.PrintWithColor(level.color(), levelStr)
	}
	messageFormat := l.formatFunc(levelStr, message)
	messageFormat = strings.TrimRightFunc(messageFormat, func(r rune) bool {
		return r == '\r' || r == '\n'
	})
	messageFormat = fmt.Sprintln(messageFormat)
	return messageFormat
}

func (l *SimpleLogger) sprintf(level Level, format string, a ...any) string {
	if l.level.level() > level.level() {
		return ""
	}
	message := fmt.Sprintf(format, a...)
	levelStr := level.String()
	if cl := l.ColorLayer(); cl != nil {
		levelStr = cl.PrintWithColor(level.color(), levelStr)
	}
	messageFormat := l.formatFunc(levelStr, message)
	messageFormat = strings.TrimRightFunc(messageFormat, func(r rune) bool {
		return r == '\r' || r == '\n'
	})
	messageFormat = fmt.Sprintln(messageFormat)
	return messageFormat
}

func (l *SimpleLogger) print(level Level, a ...any) {
	if l.level.level() > level.level() {
		return
	}
	message := fmt.Sprint(a...)
	levelStr := level.String()
	if cl := l.ColorLayer(); cl != nil {
		levelStr = cl.PrintWithColor(level.color(), levelStr)
	}
	messageFormat := l.formatFunc(levelStr, message)
	messageFormat = strings.TrimRightFunc(messageFormat, func(r rune) bool {
		return r == '\r' || r == '\n'
	})
	messageFormat = fmt.Sprintln(messageFormat)
	if l.errDone {
		switch {
		case level.Compare(LevelFatal):
			defer func() {
				os.Exit(1)
			}()
		case level.Compare(LevelPanic):
			defer func() {
				panic(message)
			}()
		}
	}
	if level.errOutput() && l.errOutput != nil {
		l.errOutput.Write([]byte(messageFormat))
		return
	}
	l.output.Write([]byte(messageFormat))
}

func (l *SimpleLogger) printf(level Level, format string, a ...any) {
	if l.level.level() > level.level() {
		return
	}
	message := fmt.Sprintf(format, a...)
	levelStr := level.String()
	if cl := l.ColorLayer(); cl != nil {
		levelStr = cl.PrintWithColor(level.color(), levelStr)
	}
	messageFormat := l.formatFunc(levelStr, message)
	messageFormat = strings.TrimRightFunc(messageFormat, func(r rune) bool {
		return r == '\r' || r == '\n'
	})
	messageFormat = fmt.Sprintln(messageFormat)
	if l.errDone {
		switch {
		case level.Compare(LevelFatal):
			defer func() {
				os.Exit(1)
			}()
		case level.Compare(LevelPanic):
			defer func() {
				panic(message)
			}()
		}
	}
	if level.errOutput() && l.errOutput != nil {
		l.errOutput.Write([]byte(messageFormat))
		return
	}
	l.output.Write([]byte(messageFormat))
}
