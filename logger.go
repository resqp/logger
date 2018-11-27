package logger

import (
	"fmt"
	"go/build"
	"log"
	"os"
)

type Logger struct {
	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
}

// New - creates a new Logger
func New(projectName string) *Logger {
	path := fmt.Sprintf("%s/%s/%s", build.Default.GOPATH, projectName, "error.log")
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return &Logger{
		trace:   log.New(file, "Trace: ", log.LstdFlags|log.Lshortfile),
		info:    log.New(file, "Info: ", log.LstdFlags|log.Lshortfile),
		warning: log.New(file, "Warning: ", log.LstdFlags|log.Lshortfile),
		error:   log.New(file, "Error: ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Trace(format string, dest ...interface{}) {
	l.trace.Printf(format, dest...)
}

func (l *Logger) Info(format string, dest ...interface{}) {
	l.info.Printf(format, dest...)
}

func (l *Logger) Warning(format string, dest ...interface{}) {
	l.warning.Printf(format, dest...)
}

func (l *Logger) Error(format string, dest ...interface{}) {
	l.error.Printf(format, dest...)
}
