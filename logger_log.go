package log

func (l myLogger) Log(s ...interface{}) {
	if !hasLevel(l.level, logFlag) {
		return
	}
	l.log(logFlag, s...)
}

func (l myLogger) Logf(format string, s ...interface{}) {
	if !hasLevel(l.level, logFlag) {
		return
	}
	l.logf(logFlag, format, s...)
}
