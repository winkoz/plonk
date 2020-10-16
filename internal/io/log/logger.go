package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/apex/log"
)

// initialises a logger and adds the flag vars to the terminal and configures their default values
func init() {
	log.SetHandler(NewPlonkHandler(os.Stderr))
	log.SetLevel(log.InfoLevel)
}

func SetLoggerSeverity(severity SeverityLevel) {
	log.SetLevel(log.Level(severity))
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	sourced().Debug(fmt.Sprintf("%+v", args...))
}

func Debugf(format string, args ...interface{}) {
	sourced().Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	sourced().Info(fmt.Sprintf("%+v", args...))
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	sourced().Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	sourced().Warn(fmt.Sprintf("%+v", args...))
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	sourced().Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	sourced().Error(fmt.Sprintf("%+v", args...))
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	sourced().Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	sourced().Fatal(fmt.Sprintf("%+v", args...))
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	sourced().Fatalf(format, args...)
}

// StartTrace logs an info message and returns a signal to be used when calling stop
func StarTrace(message string) interface{} {
	return sourced().Trace(message)
}

// StopTrace uses the passed in signal to calculate the duration of the trace. Adds the error if any.
func StopTrace(signal interface{}, err error) {
	signal.(*log.Entry).Stop(&err)
}

func sourced() *log.Entry {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}

	return log.WithFields(log.Fields{
		"file": file,
		"line": line,
	})
}
