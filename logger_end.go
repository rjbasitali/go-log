package log

import (
	"fmt"
	"time"
)

func (l myLogger) End(s ...interface{}) {
	var p []interface{} = append([]interface{}{fmt.Sprintf("END δt=%dµs", time.Now().Sub(l.begin)/1000)}, s...)
	l.log(p...)
}
