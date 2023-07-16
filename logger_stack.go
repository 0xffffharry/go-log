package log

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
)

type stackLogger struct {
	rootLogger Logger
	skip       int

	Logger
}

func NewStackLogger(rootLogger Logger) Logger {
	l := &stackLogger{
		rootLogger: rootLogger,
		skip:       3,
	}
	l.Logger = newModelToLogger(l)
	return l
}

func (l *stackLogger) print(level Level, a ...any) {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		_a := []any{fmt.Sprintf("[%s:%d] ", filepath.Base(file), line)}
		_a = append(_a, a...)
		a = _a
	}
	l.rootLogger.print(level, a...)
}

func (l *stackLogger) printf(level Level, format string, a ...any) {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		format = fmt.Sprintf("[%s:%d] %s", filepath.Base(file), line, format)
	}
	l.rootLogger.printf(level, format, a...)
}

func (l *stackLogger) sprint(level Level, a ...any) string {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		_a := []any{fmt.Sprintf("[%s:%d] ", filepath.Base(file), line)}
		_a = append(_a, a...)
		a = _a
	}
	return l.rootLogger.sprint(level, a...)
}

func (l *stackLogger) sprintf(level Level, format string, a ...any) string {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		format = fmt.Sprintf("[%s:%d] %s", filepath.Base(file), line, format)
	}
	return l.rootLogger.sprintf(level, format, a...)
}

func (l *stackLogger) RootLogger() Logger {
	return l.rootLogger
}

type stackContextLogger struct {
	rootContextLogger ContextLogger
	skip              int

	ContextLogger
}

func NewStackContextLogger(rootContextLogger ContextLogger) ContextLogger {
	l := &stackContextLogger{
		rootContextLogger: rootContextLogger,
		skip:              3,
	}
	l.ContextLogger = newModelToContextLogger(l)
	return l
}

func (l *stackContextLogger) printContext(ctx context.Context, level Level, a ...any) {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		_a := []any{fmt.Sprintf("[%s:%d] ", filepath.Base(file), line)}
		_a = append(_a, a...)
		a = _a
	}
	l.rootContextLogger.printContext(ctx, level, a...)
}

func (l *stackContextLogger) printfContext(ctx context.Context, level Level, format string, a ...any) {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		format = fmt.Sprintf("[%s:%d] %s", filepath.Base(file), line, format)
	}
	l.rootContextLogger.printfContext(ctx, level, format, a...)
}

func (l *stackContextLogger) sprintContext(ctx context.Context, level Level, a ...any) string {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		_a := []any{fmt.Sprintf("[%s:%d] ", filepath.Base(file), line)}
		_a = append(_a, a...)
		a = _a
	}
	return l.rootContextLogger.sprintContext(ctx, level, a...)
}

func (l *stackContextLogger) sprintfContext(ctx context.Context, level Level, format string, a ...any) string {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		format = fmt.Sprintf("[%s:%d] %s", filepath.Base(file), line, format)
	}
	return l.rootContextLogger.sprintfContext(ctx, level, format, a...)
}

func (l *stackContextLogger) RootLogger() Logger {
	if ml, ok := l.rootContextLogger.(RootLogger); ok {
		return ml.RootLogger()
	}
	return nil
}

func (l *stackContextLogger) RootContextLogger() ContextLogger {
	return l.rootContextLogger
}
