package log

import "context"

type bypassLogger struct {
	rootLogger Logger

	multiWriter *MultiWriter

	Logger
}

func NewBypassLogger(rootLogger Logger) BypassLogger {
	l := &bypassLogger{
		rootLogger:  rootLogger,
		multiWriter: &MultiWriter{},
	}
	l.Logger = newModelToLogger(l)
	return l
}

func (l *bypassLogger) sprint(level Level, a ...any) string {
	return l.rootLogger.sprint(level, a...)
}

func (l *bypassLogger) sprintf(level Level, format string, a ...any) string {
	return l.rootLogger.sprintf(level, format, a...)
}

func (l *bypassLogger) print(level Level, a ...any) {
	l.rootLogger.print(level, a...)
	if l.multiWriter.Len() > 0 {
		str := l.rootLogger.sprint(level, a...)
		l.multiWriter.Write([]byte(str))
	}
}

func (l *bypassLogger) printf(level Level, format string, a ...any) {
	l.rootLogger.printf(level, format, a...)
	if l.multiWriter.Len() > 0 {
		str := l.rootLogger.sprintf(level, format, a...)
		l.multiWriter.Write([]byte(str))
	}
}

func (l *bypassLogger) RootLogger() Logger {
	return l.rootLogger
}

func (l *bypassLogger) MultiWriter() *MultiWriter {
	return l.multiWriter
}

type bypassContextLogger struct {
	rootLogger ContextLogger

	multiWriter *MultiWriter

	ContextLogger
}

func NewBypassContextLogger(rootLogger ContextLogger) BypassContextLogger {
	l := &bypassContextLogger{
		rootLogger:  rootLogger,
		multiWriter: &MultiWriter{},
	}
	l.ContextLogger = newModelToContextLogger(l)
	return l
}

func (l *bypassContextLogger) sprintContext(ctx context.Context, level Level, a ...any) string {
	return l.rootLogger.sprintContext(ctx, level, a...)
}

func (l *bypassContextLogger) sprintfContext(ctx context.Context, level Level, format string, a ...any) string {
	return l.rootLogger.sprintfContext(ctx, level, format, a...)
}

func (l *bypassContextLogger) printContext(ctx context.Context, level Level, a ...any) {
	l.rootLogger.printContext(ctx, level, a...)
	if l.multiWriter.Len() > 0 {
		str := l.rootLogger.sprintContext(ctx, level, a...)
		l.multiWriter.Write([]byte(str))
	}
}

func (l *bypassContextLogger) printfContext(ctx context.Context, level Level, format string, a ...any) {
	l.rootLogger.printfContext(ctx, level, format, a...)
	if l.multiWriter.Len() > 0 {
		str := l.rootLogger.sprintfContext(ctx, level, format, a...)
		l.multiWriter.Write([]byte(str))
	}
}

func (l *bypassContextLogger) RootLogger() Logger {
	return l.rootLogger
}

func (l *bypassContextLogger) RootContextLogger() ContextLogger {
	return l.rootLogger
}

func (l *bypassContextLogger) MultiWriter() *MultiWriter {
	return l.multiWriter
}
