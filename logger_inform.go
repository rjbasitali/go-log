package log

func (l myLogger) Inform(a ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, informFlag) {
		return
	}
	l.log(a...)
}

func (l myLogger) Informf(format string, a ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, informFlag) {
		return
	}
	l.logf(format, a...)
}
