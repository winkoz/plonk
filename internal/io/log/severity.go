package log

import (
	"errors"
	"fmt"
	"strings"
)

// SeverityLevel defines the level of logging that we desire for the application
type SeverityLevel int

// ErrInvalidLevel is returned if the severity level is invalid.
var ErrInvalidLevel = errors.New("invalid level")

// Severity sets the level of verbosity for the logs; by default the desired importance is to log info or higher
var Severity SeverityLevel = InfoLevel

// Log levels.
const (
	InvalidLevel SeverityLevel = iota - 1
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var levelStrings = map[string]SeverityLevel{
	"debug":   DebugLevel,
	"info":    InfoLevel,
	"warn":    WarnLevel,
	"warning": WarnLevel,
	"error":   ErrorLevel,
	"fatal":   FatalLevel,
}

// Set implements pflag.Value
func (s *SeverityLevel) Set(value string) error {
	level, err := parseLevel(value)
	if err != nil {
		return err
	}

	// Configure logging level
	SetLoggerSeverity(level)

	Debugf("Severity Value: %s", Severity)
	return nil
}

// String implements pflag.Value
func (s SeverityLevel) String() string {
	return fmt.Sprintf("%d", s)
}

// Type implements pflag.Value
func (s SeverityLevel) Type() string {
	return fmt.Sprintf("%d", s)
}

// parseLevel parses level string.
func parseLevel(s string) (SeverityLevel, error) {
	l, ok := levelStrings[strings.ToLower(s)]
	if !ok {
		return InvalidLevel, ErrInvalidLevel
	}

	return l, nil
}
