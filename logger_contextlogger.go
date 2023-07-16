package log

import (
	"context"
	"fmt"
	"time"
)

type contextLogger struct {
	rootLogger Logger

	colorLayer ColorLayer

	ContextLogger
}

func NewContextLogger(rootLogger Logger) ContextLogger {
	l := &contextLogger{
		rootLogger: rootLogger,
	}
	l.ContextLogger = newModelToContextLogger(l)
	return l
}

func (l *contextLogger) InitColor() {
	if l.colorLayer == nil {
		l.colorLayer = newColorLayer()
	}
}

func (l *contextLogger) ColorLayer() ColorLayer {
	return l.colorLayer
}

func (l *contextLogger) printContext(ctx context.Context, level Level, a ...any) {
	ctxMsg := GetContextMessage(ctx)
	if ctxMsg == nil {
		l.rootLogger.print(level, a...)
		return
	}
	id := ctxMsg.GetID()
	duration := time.Since(ctxMsg.GetAddTime()).String()
	if cl := l.colorLayer; cl != nil {
		id = cl.Print(id)
		duration = cl.Print(duration)
	}
	_a := []any{fmt.Sprintf("[%s %s] ", id, duration)}
	_a = append(_a, a...)
	l.rootLogger.print(level, _a...)
}

func (l *contextLogger) printfContext(ctx context.Context, level Level, format string, a ...any) {
	ctxMsg := GetContextMessage(ctx)
	if ctxMsg == nil {
		l.rootLogger.printf(level, format, a...)
		return
	}
	id := ctxMsg.GetID()
	duration := time.Since(ctxMsg.GetAddTime()).String()
	if cl := l.colorLayer; cl != nil {
		id = cl.Print(id)
		duration = cl.Print(duration)
	}
	format = fmt.Sprintf("[%s %s] %s", id, duration, format)
	l.rootLogger.printf(level, format, a...)
}

func (l *contextLogger) sprintContext(ctx context.Context, level Level, a ...any) string {
	ctxMsg := GetContextMessage(ctx)
	if ctxMsg == nil {
		return l.rootLogger.sprint(level, a...)
	}
	id := ctxMsg.GetID()
	duration := time.Since(ctxMsg.GetAddTime()).String()
	if cl := l.colorLayer; cl != nil {
		id = cl.Print(id)
		duration = cl.Print(duration)
	}
	_a := []any{fmt.Sprintf("[%s %s] ", id, duration)}
	_a = append(_a, a...)
	return l.rootLogger.sprint(level, _a...)
}

func (l *contextLogger) sprintfContext(ctx context.Context, level Level, format string, a ...any) string {
	ctxMsg := GetContextMessage(ctx)
	if ctxMsg == nil {
		return l.rootLogger.sprintf(level, format, a...)
	}
	id := ctxMsg.GetID()
	duration := time.Since(ctxMsg.GetAddTime()).String()
	if cl := l.colorLayer; cl != nil {
		id = cl.Print(id)
		duration = cl.Print(duration)
	}
	format = fmt.Sprintf("[%s %s] %s", id, duration, format)
	return l.rootLogger.sprintf(level, format, a...)
}

func (l *contextLogger) RootLogger() Logger {
	return l.rootLogger
}
