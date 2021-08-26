package log

func (l myLogger) Error(a ...interface{}) {
	if !hasLevel(l.level, errorFlag) {
		return
	}
	l.log(a...)
}

func (l myLogger) Errorf(format string, a ...interface{}) {
	if !hasLevel(l.level, errorFlag) {
		return
	}
	l.logf(format, a...)
}
