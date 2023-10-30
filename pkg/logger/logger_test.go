package logger

import (
	"bytes"
	"io"
	"testing"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := GetLogger()
	setLoggerWriter(logger, &buf)

	tests := []struct {
		name     string
		logLevel int
		logFunc  func(l Logger)
		expected string
	}{
		{
			name:     "Error Message - Level None",
			logLevel: LevelNone,
			logFunc: func(l Logger) {
				l.Errorf("test %s", "message")
			},
			expected: "ERROR: test message",
		},
		{
			name:     "Error Message - Level Info",
			logLevel: LevelInfo,
			logFunc: func(l Logger) {
				l.Errorf("test %s", "message")
			},
			expected: "ERROR: test message",
		},
		{
			name:     "Error Message - Level Debug",
			logLevel: LevelDebug,
			logFunc: func(l Logger) {
				l.Errorf("test %s", "message")
			},
			expected: "ERROR: test message",
		},
		{
			name:     "Warn Message - Level None",
			logLevel: LevelNone,
			logFunc: func(l Logger) {
				l.Warnf("test %s", "message")
			},
			expected: "WARN: test message",
		},
		{
			name:     "Warn Message - Level Info",
			logLevel: LevelInfo,
			logFunc: func(l Logger) {
				l.Warnf("test %s", "message")
			},
			expected: "WARN: test message",
		},
		{
			name:     "Warn Message - Level Debug",
			logLevel: LevelDebug,
			logFunc: func(l Logger) {
				l.Warnf("test %s", "message")
			},
			expected: "WARN: test message",
		},
		{
			name:     "Info Message - Level None",
			logLevel: LevelNone,
			logFunc: func(l Logger) {
				l.Infof("test %s", "message")
			},
			expected: "",
		},
		{
			name:     "Info Message - Level Info",
			logLevel: LevelInfo,
			logFunc: func(l Logger) {
				l.Infof("test %s", "message")
			},
			expected: "INFO: test message",
		},
		{
			name:     "Info Message - Level Debug",
			logLevel: LevelDebug,
			logFunc: func(l Logger) {
				l.Infof("test %s", "message")
			},
			expected: "INFO: test message",
		},
		{
			name:     "Debug Message - Level None",
			logLevel: LevelNone,
			logFunc: func(l Logger) {
				l.Debugf("test %s", "message")
			},
			expected: "",
		},
		{
			name:     "Debug Message - Level Info",
			logLevel: LevelInfo,
			logFunc: func(l Logger) {
				l.Debugf("test %s", "message")
			},
			expected: "",
		},
		{
			name:     "Debug Message - Level Debug",
			logLevel: LevelDebug,
			logFunc: func(l Logger) {
				l.Debugf("test %s", "message")
			},
			expected: "DEBUG: test message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			logger.SetLogLevel(tt.logLevel)
			tt.logFunc(logger)

			got := buf.String()
			if got != tt.expected {
				t.Errorf("Expected: %s, got: %s", tt.expected, got)
			}
		})
	}
}

// Assuming you have a function like this to set the writer, or adjust the logger package to expose the writer setting
func setLoggerWriter(logger Logger, writer io.Writer) {
	if dl, ok := logger.(*DefaultLogger); ok {
		dl.writer = writer
	}
}
