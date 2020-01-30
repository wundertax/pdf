package pdf

import "log"

var Log Logger = &NoopLogger{}

type Logger interface {
	Warn(message string)
	Error(message string)
}

type NoopLogger struct{}

func (instance NoopLogger) Warn(string)  {}
func (instance NoopLogger) Error(string) {}

type DefaultLogger struct{}

func (instance DefaultLogger) Warn(message string)  { log.Print("WARN ", message) }
func (instance DefaultLogger) Error(message string) { log.Print("ERROR", message) }
