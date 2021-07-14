package log

func (l myLogger) Alert(a ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, alertFlag) {
		return
	}
	l.log(a...)
}

func (l myLogger) Alertf(format string, a ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, alertFlag) {
		return
	}
	l.logf(format, a...)
}
