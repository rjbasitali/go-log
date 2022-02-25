package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type myLogger struct {
	Writer io.Writer
	prefix string
	begin  time.Time
	level  uint8
}

func (l myLogger) log(s ...interface{}) {
	if l.Writer == nil {
		return
	}
	s = append([]interface{}{l.prefix, funcName()}, s...)
	fmt.Fprintln(l.Writer, s...)
}

func (l myLogger) logf(format string, s ...interface{}) {
	if l.Writer == nil {
		return
	}
	format = fmt.Sprintf("%%s %%s %s\n", format)
	s = append([]interface{}{l.prefix, funcName()}, s...)

	fmt.Fprintf(l.Writer, format, s...)
}

func funcName() string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	for _, p := range pc {
		fn := runtime.FuncForPC(p)
		if !strings.HasPrefix(fn.Name(), "github.com/rjbasitali/go-log") {
			return fn.Name()
		}
	}
	return ""
}

func NewLogger(w io.Writer) Logger {
	if w == nil {
		w = os.Stdout
	}
	return myLogger{
		Writer: w,
		begin:  time.Now(),
		level:  63,
	}
}
