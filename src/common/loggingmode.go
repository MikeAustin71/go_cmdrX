package common

/*  'loggingmode.go' is located in source code
		repository:

		https://github.com/MikeAustin71/logopsgo.git

 */

// LoggingMode - Holds Logging mode
type LoggingMode int

func (mode LoggingMode) String() string {
	return LoggingModeNames[mode]
}

const (
	// LogMINIMAL - minimum amount of information
	// will be logged
	LogMINIMAL LoggingMode = iota
	// LogNORMAL - Typical amount of information
	// will be logged
	LogNORMAL
	// LogVERBOSE - Maximum amount of information
	// will be logged
	LogVERBOSE
)

// LoggingModeNames - string array containing names of Logging Modes
var LoggingModeNames = [...]string{"MINIMAL", "NORMAL", "VERBOSE"}
