package log

func (l myLogger) Trace(a ...interface{}) {
	if !hasLevel(l.level, traceFlag) {
		return
	}
	l.log(a...)
}

func (l myLogger) Tracef(format string, a ...interface{}) {
	if !hasLevel(l.level, traceFlag) {
		return
	}
	l.logf(format, a...)
}
