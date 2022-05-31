package log

import (
	"fmt"
	"strings"
)

func (l myLogger) Prefix(p ...string) Logger {
	var buffer strings.Builder
	buffer.WriteString(l.prefix)
	for _, prefix := range p {
		buffer.WriteString(fmt.Sprintf("%s: ", prefix))
	}
	logger := myLogger{Writer: l.Writer, ErrWriter: l.ErrWriter, prefix: buffer.String(), data: l.data}
	return logger
}
