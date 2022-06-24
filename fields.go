package log

func (l myLogger) WithFields(fields Fields) Logger {
	if l.data != nil && len(l.data) > 0 {
		f := Fields{}
		for k, v := range l.data {
			f[k] = v
		}
		for k, v := range fields {
			f[k] = v
		}
		l.data = f
	} else {
		l.data = fields
	}
	return l
}
