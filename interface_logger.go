package log

type _Logger interface {
	Print(level Level, a ...any)
	Printf(level Level, format string, a ...any)
	Debug(a ...any)
	Debugf(format string, a ...any)
	Info(a ...any)
	Infof(format string, a ...any)
	Warn(a ...any)
	Warnf(format string, a ...any)
	Error(a ...any)
	Errorf(format string, a ...any)
	Fatal(a ...any)
	Fatalf(format string, a ...any)
	Panic(a ...any)
	Panicf(format string, a ...any)
}

type Logger interface {
	_ModelLogger
	_Logger
}

type _ModelLogger interface {
	sprint(level Level, a ...any) string
	sprintf(level Level, format string, a ...any) string
	print(level Level, a ...any)
	printf(level Level, format string, a ...any)
}

type modelToLogger struct {
	_ModelLogger
}

func newModelToLogger(modelLogger _ModelLogger) Logger {
	return &modelToLogger{
		_ModelLogger: modelLogger,
	}
}

func (l *modelToLogger) sprint(level Level, a ...any) string {
	return l._ModelLogger.sprint(level, a...)
}

func (l *modelToLogger) sprintf(level Level, format string, a ...any) string {
	return l._ModelLogger.sprintf(level, format, a...)
}

func (l *modelToLogger) print(level Level, a ...any) {
	l._ModelLogger.print(level, a...)
}

func (l *modelToLogger) printf(level Level, format string, a ...any) {
	l._ModelLogger.printf(level, format, a...)
}

func (l *modelToLogger) Print(level Level, a ...any) {
	l.print(level, a...)
}

func (l *modelToLogger) Printf(level Level, format string, a ...any) {
	l.printf(level, format, a...)
}

func (l *modelToLogger) Debug(a ...any) {
	l.print(LevelDebug, a...)
}

func (l *modelToLogger) Debugf(format string, a ...any) {
	l.printf(LevelDebug, format, a...)
}

func (l *modelToLogger) Info(a ...any) {
	l.print(LevelInfo, a...)
}

func (l *modelToLogger) Infof(format string, a ...any) {
	l.printf(LevelInfo, format, a...)
}

func (l *modelToLogger) Warn(a ...any) {
	l.print(LevelWarn, a...)
}

func (l *modelToLogger) Warnf(format string, a ...any) {
	l.printf(LevelWarn, format, a...)
}

func (l *modelToLogger) Error(a ...any) {
	l.print(LevelError, a...)
}

func (l *modelToLogger) Errorf(format string, a ...any) {
	l.printf(LevelError, format, a...)
}

func (l *modelToLogger) Fatal(a ...any) {
	l.print(LevelFatal, a...)
}

func (l *modelToLogger) Fatalf(format string, a ...any) {
	l.printf(LevelFatal, format, a...)
}

func (l *modelToLogger) Panic(a ...any) {
	l.print(LevelPanic, a...)
}

func (l *modelToLogger) Panicf(format string, a ...any) {
	l.printf(LevelPanic, format, a...)
}
