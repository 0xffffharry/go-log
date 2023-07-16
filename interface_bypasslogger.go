package log

type MultiWriterLogger interface {
	MultiWriter() *MultiWriter
}

type BypassLogger interface {
	Logger
	MultiWriterLogger
}

type BypassContextLogger interface {
	ContextLogger
	MultiWriterLogger
}
