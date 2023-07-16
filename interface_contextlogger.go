package log

import "context"

type _ContextLogger interface {
	PrintContext(ctx context.Context, level Level, a ...any)
	PrintfContext(ctx context.Context, level Level, format string, a ...any)
	DebugContext(ctx context.Context, a ...any)
	DebugfContext(ctx context.Context, format string, a ...any)
	InfoContext(ctx context.Context, a ...any)
	InfofContext(ctx context.Context, format string, a ...any)
	WarnContext(ctx context.Context, a ...any)
	WarnfContext(ctx context.Context, format string, a ...any)
	ErrorContext(ctx context.Context, a ...any)
	ErrorfContext(ctx context.Context, format string, a ...any)
	FatalContext(ctx context.Context, a ...any)
	FatalfContext(ctx context.Context, format string, a ...any)
	PanicContext(ctx context.Context, a ...any)
	PanicfContext(ctx context.Context, format string, a ...any)
}

type _ModelContextLogger interface {
	sprintContext(ctx context.Context, level Level, a ...any) string
	sprintfContext(ctx context.Context, level Level, format string, a ...any) string
	printContext(ctx context.Context, level Level, a ...any)
	printfContext(ctx context.Context, level Level, format string, a ...any)
}

type ContextLogger interface {
	_ModelLogger
	_ModelContextLogger
	_Logger
	_ContextLogger
}

type modelToContextLogger struct {
	_ModelContextLogger
}

func newModelToContextLogger(modelContextLogger _ModelContextLogger) ContextLogger {
	l := &modelToContextLogger{
		_ModelContextLogger: modelContextLogger,
	}
	return l
}

func (l *modelToContextLogger) sprint(level Level, a ...any) string {
	return l._ModelContextLogger.sprintContext(context.Background(), level, a...)
}

func (l *modelToContextLogger) sprintf(level Level, format string, a ...any) string {
	return l._ModelContextLogger.sprintfContext(context.Background(), level, format, a...)
}

func (l *modelToContextLogger) print(level Level, a ...any) {
	l._ModelContextLogger.printContext(context.Background(), level, a...)
}

func (l *modelToContextLogger) printf(level Level, format string, a ...any) {
	l._ModelContextLogger.printfContext(context.Background(), level, format, a...)
}

func (l *modelToContextLogger) Print(level Level, a ...any) {
	l.print(level, a...)
}

func (l *modelToContextLogger) Printf(level Level, format string, a ...any) {
	l.printf(level, format, a...)
}

func (l *modelToContextLogger) Debug(a ...any) {
	l.print(LevelDebug, a...)
}

func (l *modelToContextLogger) Debugf(format string, a ...any) {
	l.printf(LevelDebug, format, a...)
}

func (l *modelToContextLogger) Info(a ...any) {
	l.print(LevelInfo, a...)
}

func (l *modelToContextLogger) Infof(format string, a ...any) {
	l.printf(LevelInfo, format, a...)
}

func (l *modelToContextLogger) Warn(a ...any) {
	l.print(LevelWarn, a...)
}

func (l *modelToContextLogger) Warnf(format string, a ...any) {
	l.printf(LevelWarn, format, a...)
}

func (l *modelToContextLogger) Error(a ...any) {
	l.print(LevelError, a...)
}

func (l *modelToContextLogger) Errorf(format string, a ...any) {
	l.printf(LevelError, format, a...)
}

func (l *modelToContextLogger) Fatal(a ...any) {
	l.print(LevelFatal, a...)
}

func (l *modelToContextLogger) Fatalf(format string, a ...any) {
	l.printf(LevelFatal, format, a...)
}

func (l *modelToContextLogger) Panic(a ...any) {
	l.print(LevelPanic, a...)
}

func (l *modelToContextLogger) Panicf(format string, a ...any) {
	l.printf(LevelPanic, format, a...)
}

func (l *modelToContextLogger) printContext(ctx context.Context, level Level, a ...any) {
	l._ModelContextLogger.printContext(ctx, level, a...)
}

func (l *modelToContextLogger) printfContext(ctx context.Context, level Level, format string, a ...any) {
	l._ModelContextLogger.printfContext(ctx, level, format, a...)
}

func (l *modelToContextLogger) PrintContext(ctx context.Context, level Level, a ...any) {
	l.printContext(ctx, level, a...)
}

func (l *modelToContextLogger) PrintfContext(ctx context.Context, level Level, format string, a ...any) {
	l.printfContext(ctx, level, format, a...)
}

func (l *modelToContextLogger) DebugContext(ctx context.Context, a ...any) {
	l.printContext(ctx, LevelDebug, a...)
}

func (l *modelToContextLogger) DebugfContext(ctx context.Context, format string, a ...any) {
	l.printfContext(ctx, LevelDebug, format, a...)
}

func (l *modelToContextLogger) InfoContext(ctx context.Context, a ...any) {
	l.printContext(ctx, LevelInfo, a...)
}

func (l *modelToContextLogger) InfofContext(ctx context.Context, format string, a ...any) {
	l.printfContext(ctx, LevelInfo, format, a...)
}

func (l *modelToContextLogger) WarnContext(ctx context.Context, a ...any) {
	l.printContext(ctx, LevelWarn, a...)
}

func (l *modelToContextLogger) WarnfContext(ctx context.Context, format string, a ...any) {
	l.printfContext(ctx, LevelWarn, format, a...)
}

func (l *modelToContextLogger) ErrorContext(ctx context.Context, a ...any) {
	l.printContext(ctx, LevelError, a...)
}

func (l *modelToContextLogger) ErrorfContext(ctx context.Context, format string, a ...any) {
	l.printfContext(ctx, LevelError, format, a...)
}

func (l *modelToContextLogger) FatalContext(ctx context.Context, a ...any) {
	l.printContext(ctx, LevelFatal, a...)
}

func (l *modelToContextLogger) FatalfContext(ctx context.Context, format string, a ...any) {
	l.printfContext(ctx, LevelFatal, format, a...)
}

func (l *modelToContextLogger) PanicContext(ctx context.Context, a ...any) {
	l.printContext(ctx, LevelPanic, a...)
}

func (l *modelToContextLogger) PanicfContext(ctx context.Context, format string, a ...any) {
	l.printfContext(ctx, LevelPanic, format, a...)
}
