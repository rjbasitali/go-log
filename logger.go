package log

type Logger interface {
	Alert(a ...interface{})
	Alertf(format string, a ...interface{})

	Error(a ...interface{})
	Errorf(format string, a ...interface{})

	Highlight(a ...interface{})
	Highlightf(format string, a ...interface{})

	Inform(a ...interface{})
	Informf(format string, a ...interface{})

	Log(a ...interface{})
	Logf(format string, a ...interface{})

	Trace(a ...interface{})
	Tracef(format string, a ...interface{})

	Warn(a ...interface{})
	Warnf(format string, a ...interface{})

	Prefix(...string) Logger

	Begin(a ...interface{}) Logger

	End(a ...interface{})

	Level(uint8) Logger

	WithFields(fields Fields) Logger
}
