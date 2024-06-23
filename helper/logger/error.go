package logger

import (
	"fmt"
	"runtime/debug"
)

func (l *Log) Error(message ...any) {
	if len(message) > 0 {
		message[0] = "⚠️ " + fmt.Sprintf("%v", message[0])
		message = append(message, "\n\n", string(debug.Stack()))
	}
	l.stderr.Println(message...)
}
