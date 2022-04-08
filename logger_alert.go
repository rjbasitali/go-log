package log

func (l myLogger) Alert(a ...interface{}) {
	if !hasLevel(l.level, alertFlag) {
		return
	}
	l.log(alertFlag, a...)
}

func (l myLogger) Alertf(format string, a ...interface{}) {
	if !hasLevel(l.level, alertFlag) {
		return
	}
	l.logf(alertFlag, format, a...)
}
