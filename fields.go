package log

func (l myLogger) WithFields(fields Fields) Logger {
	if l.data != nil && len(l.data) > 0 {
		for k, v := range fields {
			l.data[k] = v
		}
	} else {
		l.data = fields
	}
	return l
}
