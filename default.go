package log

var (
	DefaultLogger        Logger        = NewSimpleLogger()
	DefaultContextLogger ContextLogger = NewContextLogger(DefaultLogger)
)

func init() {
	simpleLogger, _ := DefaultLogger.(*SimpleLogger)
	simpleLogger.SetErrDone(true)
}

var (
	Print  = DefaultLogger.Print
	Printf = DefaultLogger.Printf
	Info   = DefaultLogger.Info
	Infof  = DefaultLogger.Infof
	Warn   = DefaultLogger.Warn
	Warnf  = DefaultLogger.Warnf
	Error  = DefaultLogger.Error
	Errorf = DefaultLogger.Errorf
	Fatal  = DefaultLogger.Fatal
	Fatalf = DefaultLogger.Fatalf
	Panic  = DefaultLogger.Panic
	Panicf = DefaultLogger.Panicf
)

var (
	PrintContext  = DefaultContextLogger.PrintContext
	PrintfContext = DefaultContextLogger.PrintfContext
	InfoContext   = DefaultContextLogger.InfoContext
	InfofContext  = DefaultContextLogger.InfofContext
	WarnContext   = DefaultContextLogger.WarnContext
	WarnfContext  = DefaultContextLogger.WarnfContext
	ErrorContext  = DefaultContextLogger.ErrorContext
	ErrorfContext = DefaultContextLogger.ErrorfContext
	FatalContext  = DefaultContextLogger.FatalContext
	FatalfContext = DefaultContextLogger.FatalfContext
	PanicContext  = DefaultContextLogger.PanicContext
	PanicfContext = DefaultContextLogger.PanicfContext
)
