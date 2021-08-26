package log

func (l myLogger) Highlight(a ...interface{}) {
	if !hasLevel(l.level, highlightFlag) {
		return
	}
	l.log(a...)
}

func (l myLogger) Highlightf(format string, a ...interface{}) {
	if !hasLevel(l.level, highlightFlag) {
		return
	}
	l.logf(format, a...)
}
