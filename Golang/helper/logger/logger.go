package logger

import (
	"log"
	"os"
)

type Log struct {
	stdout *log.Logger
	stderr *log.Logger
}

func New(prefix string) *Log {
	return &Log{
		stdout: log.New(os.Stdout, "[LOG]["+prefix+"]", log.Ldate|log.Ltime),
		stderr: log.New(os.Stderr, "[ERROR]["+prefix+"]", log.Ldate|log.Ltime),
	}
}
