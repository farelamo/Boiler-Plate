package logger

import "fmt"

func (l *Log) Log(message ...any) {
	if len(message) > 0 {
		message[0] = "📢 " + fmt.Sprintf("%v", message[0])
	}
	l.stdout.Println(message...)
}
