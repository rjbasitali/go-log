package log

func (l myLogger) Warn(a ...interface{}) {
	if !hasLevel(l.level, warnFlag) {
		return
	}
	l.log(a...)
}

func (l myLogger) Warnf(format string, a ...interface{}) {
	if !hasLevel(l.level, warnFlag) {
		return
	}
	l.logf(format, a...)
}
