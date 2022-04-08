package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	DefaultTimeFormat = time.RFC3339
)

type myLogger struct {
	Writer io.Writer
	prefix string
	begin  time.Time
	level  uint8
}

func logPrefix(level uint8) string {
	return fmt.Sprintf("time=%q level=%s function=%s", time.Now().Format(DefaultTimeFormat), unmarshalLevel(level), funcName())
}

func (l myLogger) log(level uint8, s ...interface{}) {
	if l.Writer == nil {
		return
	}
	f := fmt.Sprintf("%s%s msg=%q", logPrefix(level), l.prefix, fmt.Sprint(s...))
	fmt.Fprintln(l.Writer, f)
}

func (l myLogger) logf(level uint8, format string, s ...interface{}) {
	if l.Writer == nil {
		return
	}
	format = fmt.Sprintf("%%s %%s msg=\"%s\"\n", format)
	s = append([]interface{}{logPrefix(level), l.prefix}, s...)

	fmt.Fprintf(l.Writer, format, s...)
}

func funcName() string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	for _, p := range pc {
		fn := runtime.FuncForPC(p)
		if !strings.HasPrefix(fn.Name(), "github.com/rjbasitali/go-log") {
			fname := fn.Name()
			i := strings.LastIndexByte(fname, byte('/'))
			if i == -1 {
				return fname
			}
			return fname[i+1:]
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
		level:  traceFlag,
	}
}
