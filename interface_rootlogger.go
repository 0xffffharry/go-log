package log

type RootLogger interface {
	RootLogger() Logger
}

type RootContextLogger interface {
	RootContextLogger() ContextLogger
}
