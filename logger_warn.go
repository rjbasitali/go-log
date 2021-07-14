package log

func (l myLogger) Warn(a ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, warnFlag) {
		return
	}
	l.log(a...)
}

func (l myLogger) Warnf(format string, a ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, warnFlag) {
		return
	}
	l.logf(format, a...)
}
