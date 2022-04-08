package log

func (l myLogger) Trace(a ...interface{}) {
	if !hasLevel(l.level, traceFlag) {
		return
	}
	l.log(traceFlag, a...)
}

func (l myLogger) Tracef(format string, a ...interface{}) {
	if !hasLevel(l.level, traceFlag) {
		return
	}
	l.logf(traceFlag, format, a...)
}
