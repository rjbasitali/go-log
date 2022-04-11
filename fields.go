package log

func (l myLogger) WithFields(fields Fields) Logger {
	l.data = fields
	return l
}
