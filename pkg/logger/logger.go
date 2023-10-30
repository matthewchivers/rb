package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	once         sync.Once
	globalLogger Logger
)

const (
	// LevelInfo is the log level for info messages
	LevelInfo = iota
	// LevelDebug is the log level for debug messages
	LevelDebug = iota
)

// Logger is the interface that wraps the basic logging methods
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	SetLogLevel(level int)
	SetWriter(w io.Writer)
}

// DefaultLogger is the default logger
type DefaultLogger struct {
	logLevel int
	writer   io.Writer
}

// GetLogger returns the global logger
func GetLogger() Logger {
	once.Do(func() {
		globalLogger = newLogger(LevelInfo)
	})
	return globalLogger
}

// NewLogger returns a new DefaultLogger
func newLogger(level int) *DefaultLogger {
	return &DefaultLogger{logLevel: level, writer: io.Writer(os.Stdout)}
}

// Debugf logs a debug message
func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	if l.logLevel >= LevelDebug {
		fmt.Fprintf(l.writer, "DEBUG: "+format, args...)
		fmt.Print("\n")
	}
}

// Infof logs an info message
func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	fmt.Fprintf(l.writer, "INFO: "+format, args...)
	fmt.Print("\n")
}

// Warnf logs a warning message
func (l *DefaultLogger) Warnf(format string, args ...interface{}) {
	fmt.Fprintf(l.writer, "WARN: "+format, args...)
	fmt.Print("\n")
}

// Errorf logs an error message
func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(l.writer, "ERROR: "+format, args...)
	fmt.Print("\n")
}

// SetLogLevel sets the log level
func (l *DefaultLogger) SetLogLevel(level int) {
	l.logLevel = level
}

// SetWriter sets the writer
func (l *DefaultLogger) SetWriter(w io.Writer) {
	l.writer = w
}
