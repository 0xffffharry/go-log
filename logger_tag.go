package log

import "fmt"

type tagLogger struct {
	rootLogger Logger
	tag        string

	colorLayer ColorLayer

	Logger
}

func NewTagLogger(rootLogger Logger, tag string) Logger {
	l := &tagLogger{
		rootLogger: rootLogger,
		tag:        tag,
	}
	l.Logger = newModelToLogger(l)
	return l
}

func (l *tagLogger) InitColor() {
	if l.colorLayer == nil {
		l.colorLayer = newColorLayer()
	}
}

func (l *tagLogger) ColorLayer() ColorLayer {
	return l.colorLayer
}

func (l *tagLogger) sprint(level Level, a ...any) string {
	tag := l.tag
	if cl := l.ColorLayer(); cl != nil {
		tag = cl.Print(tag)
	}
	_a := []any{"[", tag, "] "}
	_a = append(_a, a...)
	return l.rootLogger.sprint(level, _a...)
}

func (l *tagLogger) sprintf(level Level, format string, a ...any) string {
	tag := l.tag
	if cl := l.ColorLayer(); cl != nil {
		tag = cl.Print(tag)
	}
	format = fmt.Sprintf("[%s] %s", tag, format)
	return l.rootLogger.sprintf(level, format, a...)
}

func (l *tagLogger) print(level Level, a ...any) {
	tag := l.tag
	if cl := l.ColorLayer(); cl != nil {
		tag = cl.Print(tag)
	}
	_a := []any{"[", tag, "] "}
	_a = append(_a, a...)
	l.rootLogger.print(level, _a...)
}

func (l *tagLogger) printf(level Level, format string, a ...any) {
	tag := l.tag
	if cl := l.ColorLayer(); cl != nil {
		tag = cl.Print(tag)
	}
	format = fmt.Sprintf("[%s] %s", tag, format)
	l.rootLogger.printf(level, format, a...)
}

func (l *tagLogger) RootLogger() Logger {
	return l.rootLogger
}
