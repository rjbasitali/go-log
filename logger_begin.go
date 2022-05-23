package log

import "time"

func (l myLogger) Begin(s ...interface{}) Logger {
	if hasLevel(l.level, logFlag) {
		l.log(logFlag, append([]interface{}{"BEGIN"}, s...)...)
	}
	logger := myLogger{Writer: l.Writer, ErrWriter: l.ErrWriter, prefix: l.prefix, begin: time.Now(), level: l.level, data: l.data}
	return logger
}
