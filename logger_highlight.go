package log

func (l myLogger) Highlight(a ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, highlightFlag) {
		return
	}
	l.log(a...)
}

func (l myLogger) Highlightf(format string, a ...interface{}) {
	if l.Writer == nil {
		return
	}
	if !hasLevel(l.level, highlightFlag) {
		return
	}
	l.logf(format, a...)
}
