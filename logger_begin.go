package log

import "time"

func (l myLogger) Begin(s ...interface{}) Logger {
	logger := myLogger{Writer: l.Writer, prefix: l.prefix, begin: time.Now(), level: l.level}
	l.log(append([]interface{}{"BEGIN"}, s...)...)
	return logger
}
