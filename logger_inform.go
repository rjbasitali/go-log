package log

func (l myLogger) Inform(a ...interface{}) {
	if !hasLevel(l.level, informFlag) {
		return
	}
	l.log(informFlag, a...)
}

func (l myLogger) Informf(format string, a ...interface{}) {
	if !hasLevel(l.level, informFlag) {
		return
	}
	l.logf(informFlag, format, a...)
}
