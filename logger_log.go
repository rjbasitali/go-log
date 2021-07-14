package log

func (l myLogger) Log(s ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, logFlag) {
		return
	}
	l.log(s...)
}

func (l myLogger) Logf(format string, s ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, logFlag) {
		return
	}
	l.logf(format, s...)
}
