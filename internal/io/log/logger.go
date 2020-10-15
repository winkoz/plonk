package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/apex/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

// setSyslogFormatter is nil if the target architecture does not support syslog.
var setSyslogFormatter func(logger, string, string) error

// setEventlogFormatter is nil if the target OS does not support Eventlog (i.e., is not Windows).
var setEventlogFormatter func(logger, string, bool) error

	log.SetHandler(NewPlonkHandler(os.Stderr))
	log.SetLevel(log.InfoLevel)
}

type loggerSettings struct {
	level  string
	format string
}

func SetLoggerSeverity(severity SeverityLevel) {
	err := baseLogger.SetLevel(s.level)
	if err != nil {
		return err
	}
	err = baseLogger.SetFormat(s.format)
	return err
}

// AddFlags adds the flags used by this package to the Kingpin application.
// To use the default Kingpin application, call AddFlags(kingpin.CommandLine)
func AddFlags(a *kingpin.Application) {
	s := loggerSettings{}
	log.SetLevel(log.Level(severity))
		Default(origLogger.Level.String()).
		StringVar(&s.level)
	defaultFormat := url.URL{Scheme: "logger", Opaque: "stderr"}
	a.Flag("log.format", `Set the log target and format. Example: "logger:syslog?appname=bob&local=7" or "logger:stdout?json=true"`).
		Default(defaultFormat.String()).
		StringVar(&s.format)
	a.Action(s.apply)
}

// Logger is the interface for loggers used in the Prometheus components.
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

	With(key string, value interface{}) Logger

	SetFormat(string) error
	SetLevel(string) error
}

type logger struct {
	entry *logrus.Entry
}

func (l logger) With(key string, value interface{}) Logger {
	return logger{l.entry.WithField(key, value)}
}

// Debug logs a message at level Debug on the standard logger.
func (l logger) Debug(args ...interface{}) {
	l.sourced().Debug(args...)
}

// Debug logs a message at level Debug on the standard logger.
func (l logger) Debugln(args ...interface{}) {
	l.sourced().Debugln(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l logger) Debugf(format string, args ...interface{}) {
	l.sourced().Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func (l logger) Info(args ...interface{}) {
	l.sourced().Info(args...)
}

// Info logs a message at level Info on the standard logger.
func (l logger) Infoln(args ...interface{}) {
	l.sourced().Infoln(args...)
}

// Infof logs a message at level Info on the standard logger.
func (l logger) Infof(format string, args ...interface{}) {
	l.sourced().Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l logger) Warn(args ...interface{}) {
	l.sourced().Warn(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l logger) Warnln(args ...interface{}) {
	l.sourced().Warnln(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l logger) Warnf(format string, args ...interface{}) {
	l.sourced().Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func (l logger) Error(args ...interface{}) {
	l.sourced().Error(args...)
}

// Error logs a message at level Error on the standard logger.
func (l logger) Errorln(args ...interface{}) {
	l.sourced().Errorln(args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l logger) Errorf(format string, args ...interface{}) {
	l.sourced().Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l logger) Fatal(args ...interface{}) {
	l.sourced().Fatal(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l logger) Fatalln(args ...interface{}) {
	l.sourced().Fatalln(args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func (l logger) Fatalf(format string, args ...interface{}) {
	l.sourced().Fatalf(format, args...)
}

func (l logger) SetLevel(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}

	l.entry.Logger.Level = lvl
	return nil
}

func (l logger) SetFormat(format string) error {
	u, err := url.Parse(format)
	if err != nil {
		return err
	}
	if u.Scheme != "logger" {
		return fmt.Errorf("invalid scheme %s", u.Scheme)
	}
	jsonq := u.Query().Get("json")
	if jsonq == "true" {
		setJSONFormatter()
	}

	switch u.Opaque {
	case "syslog":
		if setSyslogFormatter == nil {
			return fmt.Errorf("system does not support syslog")
		}
		appname := u.Query().Get("appname")
		facility := u.Query().Get("local")
		return setSyslogFormatter(l, appname, facility)
	case "eventlog":
		if setEventlogFormatter == nil {
			return fmt.Errorf("system does not support eventlog")
		}
		name := u.Query().Get("name")
		debugAsInfo := false
		debugAsInfoRaw := u.Query().Get("debugAsInfo")
		if parsedDebugAsInfo, err := strconv.ParseBool(debugAsInfoRaw); err == nil {
			debugAsInfo = parsedDebugAsInfo
		}
		return setEventlogFormatter(l, name, debugAsInfo)
	case "stdout":
		l.entry.Logger.Out = os.Stdout
	case "stderr":
		l.entry.Logger.Out = os.Stderr
	default:
		return fmt.Errorf("unsupported logger %q", u.Opaque)
	}
	return nil
}

// sourced adds a source field to the logger that contains
// the file name and line where the logging happened.
func (l logger) sourced() *logrus.Entry {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	return l.entry.WithField("source", fmt.Sprintf("%s:%d", file, line))
}

var origLogger = logrus.New()
var baseLogger = logger{entry: logrus.NewEntry(origLogger)}

// Base returns the default Logger logging to
func Base() Logger {
	return baseLogger
}

// NewLogger returns a new Logger logging to out.
func NewLogger(w io.Writer) Logger {
	l := logrus.New()
	l.Out = w
	return logger{entry: logrus.NewEntry(l)}
}

// NewNopLogger returns a logger that discards all log messages.
func NewNopLogger() Logger {
	l := logrus.New()
	l.Out = ioutil.Discard
	return logger{entry: logrus.NewEntry(l)}
}

// With adds a field to the logger.
func With(key string, value interface{}) Logger {
	return baseLogger.With(key, value)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	sourced().Debug(fmt.Sprintf("%+v", args...))
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	baseLogger.sourced().Debugln(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	sourced().Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	sourced().Info(fmt.Sprintf("%+v", args...))
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	baseLogger.sourced().Infoln(args...)
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

// Fatalln logs a message at level Fatal on the standard logger.
func Fatalln(args ...interface{}) {
	baseLogger.sourced().Fatalln(args...)
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
