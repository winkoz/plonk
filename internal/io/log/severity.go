package log

// SeverityLevel defines the level of logging that we desire for the application
type SeverityLevel string

// Severity sets the level of verbosity for the logs; by default the desired importance is to log info or higher
var Severity SeverityLevel = "INFO"

// Set implements pflag.Value
func (s *SeverityLevel) Set(value string) error {
	Severity = SeverityLevel(value)
	// Configure logging level
	if err := SetLevel(value); err != nil {
		panic(err.Error())
	}

	Debugf("Severity Value: %s", Severity)
	return nil
}

// String implements pflag.Value
func (s SeverityLevel) String() string {
	return string(s)
}

// Type implements pflag.Value
func (s SeverityLevel) Type() string {
	return string(s)
}
