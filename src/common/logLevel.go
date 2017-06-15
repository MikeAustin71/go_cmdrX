package common

// LogLevel - Holds the message level indicating the relative importance of a specific log message.
type LogLevel int

func (level LogLevel) String() string {
	return LogLevelNames[level]
}

const (
	// LogDEBUG - message
	LogDEBUG LogLevel = iota
	// LogOPERROR - 1 Message is an Error Message
	LogOPERROR
	// LogFATAL - 2 Message is a Fatal Error Message
	LogFATAL
	// LogINFO - 3 Message is an Informational Message
	LogINFO
	// LogWARN - 4 Message is a warning Message
	LogWARN
)

// LogLevelNames - string array containing names of Log Levels
var LogLevelNames = [...]string{"DEBUG", "OPERROR", "FATAL", "INFO", "WARN"}
