package log

import (
	"fmt"
	"io"
)

type Logger struct {
	Writer io.Writer
}

func (l Logger) Log(s ...interface{}) {
	if l.Writer == nil {
		return
	}
	fmt.Fprintln(l.Writer, s...)
}

func (l Logger) Logf(format string, s ...interface{}) {
	if l.Writer == nil {
		return
	}
	format = fmt.Sprintf("%s%s", format, "\n")
	fmt.Fprintf(l.Writer, format, s...)
}

func (l Logger) Begin(s ...interface{}) {
	var p []interface{} = append([]interface{}{"BEGIN "}, s...)
	l.Log(p...)
}

func (l Logger) End(s ...interface{}) {
	var p []interface{} = append([]interface{}{"END "}, s...)
	l.Log(p...)
}

func (l Logger) Error(s string, err error) {
	l.Log(s, err)
}
