package log

import "time"

func (l myLogger) Begin(s ...interface{}) Logger {
	if hasLevel(l.level, logFlag) {
		l.log(append([]interface{}{"BEGIN"}, s...)...)
	}
	logger := myLogger{Writer: l.Writer, prefix: l.prefix, begin: time.Now(), level: l.level}
	return logger
}
