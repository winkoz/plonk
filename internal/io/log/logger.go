package log

import (
	"github.com/prometheus/common/log"
)

// Logger is the interface for loggers
type Logger interface {
	Debug(...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infoln(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnln(...interface{})
	Warnf(string, ...interface{})

	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})

	SetLevel(string) error
}

type logger struct{}

var baseLogger logger

// Debug logs a message at level Debug on the standard log.
func (l logger) Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debug logs a message at level Debug on the standard log.
func (l logger) Debugln(args ...interface{}) {
	log.Debugln(args...)
}

// Debugf logs a message at level Debug on the standard log.
func (l logger) Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Info logs a message at level Info on the standard log.
func (l logger) Info(args ...interface{}) {
	log.Info(args...)
}

// Info logs a message at level Info on the standard log.
func (l logger) Infoln(args ...interface{}) {
	log.Infoln(args...)
}

// Infof logs a message at level Info on the standard log.
func (l logger) Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard log.
func (l logger) Warn(args ...interface{}) {
	log.Warn(args...)
}

// Warn logs a message at level Warn on the standard log.
func (l logger) Warnln(args ...interface{}) {
	log.Warnln(args...)
}

// Warnf logs a message at level Warn on the standard log.
func (l logger) Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Error logs a message at level Error on the standard log.
func (l logger) Error(args ...interface{}) {
	log.Error(args...)
}

// Error logs a message at level Error on the standard log.
func (l logger) Errorln(args ...interface{}) {
	log.Errorln(args...)
}

// Errorf logs a message at level Error on the standard log.
func (l logger) Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard log.
func (l logger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatal logs a message at level Fatal on the standard log.
func (l logger) Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}

// Fatalf logs a message at level Fatal on the standard log.
func (l logger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func (l logger) SetLevel(level string) error {
	if err := log.Base().SetLevel(level); err != nil {
		return err
	}

	return nil
}

// Debug logs a message at level Debug on the standard log.
func Debug(args ...interface{}) {
	baselog.Debug(args...)
}

// Debugln logs a message at level Debug on the standard log.
func Debugln(args ...interface{}) {
	baselog.Debugln(args...)
}

// Debugf logs a message at level Debug on the standard log.
func Debugf(format string, args ...interface{}) {
	baselog.Debugf(format, args...)
}

// Info logs a message at level Info on the standard log.
func Info(args ...interface{}) {
	baselog.Info(args...)
}

// Infoln logs a message at level Info on the standard log.
func Infoln(args ...interface{}) {
	baselog.Infoln(args...)
}

// Infof logs a message at level Info on the standard log.
func Infof(format string, args ...interface{}) {
	baselog.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard log.
func Warn(args ...interface{}) {
	baselog.Warn(args...)
}

// Warnln logs a message at level Warn on the standard log.
func Warnln(args ...interface{}) {
	baselog.Warnln(args...)
}

// Warnf logs a message at level Warn on the standard log.
func Warnf(format string, args ...interface{}) {
	baselog.Warnf(format, args...)
}

// Error logs a message at level Error on the standard log.
func Error(args ...interface{}) {
	baselog.Error(args...)
}

// Errorln logs a message at level Error on the standard log.
func Errorln(args ...interface{}) {
	baselog.Errorln(args...)
}

// Errorf logs a message at level Error on the standard log.
func Errorf(format string, args ...interface{}) {
	baselog.Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard log.
func Fatal(args ...interface{}) {
	baselog.Fatal(args...)
}

// Fatalln logs a message at level Fatal on the standard log.
func Fatalln(args ...interface{}) {
	baselog.Fatalln(args...)
}

// Fatalf logs a message at level Fatal on the standard log.
func Fatalf(format string, args ...interface{}) {
	baselog.Fatalf(format, args...)
}

// SetLevel sets the verbosity level for the logs
func SetLevel(level string) error {
	return baselog.SetLevel(level)
}

func init() {
	baseLogger = logger{}
}
