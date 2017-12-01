package common

import "time"

// CommandJobGroup - Contains a
// slice of CommandJobInfo structures
// plus Group Header Info.
type CommandJobGroup struct {
	CmdFileVersion          string
	DefaultCmdExeDir        string
	LogPath									string
	LogFileName							string
	LogPathFileName					string
	LogFileRetention        time.Duration
	KillAllJobsOnFirstError bool
	IanaTimeZone						string
	NoOfCmdJobs							int
}

// CommandJobInfo - Contains information on
// a Command Job to be executed
type CommandJobInfo struct {
	DisplayName                   string
	Description                   string
	CommandType                   string
	ExeDir                        string
	DelayCmdStart                 time.Duration
	StartCmdDateTime              time.Time
	KillJobsOnExitCodeGreaterThan int
	KillJobsOnExitCodeLessThan    int
	CommandTimeOut                time.Duration
	ExeCommand                    string
	CmdExeStartTime               time.Time
	CmdExeEndTime                 time.Time
	CmdExeElapsedTime             time.Duration
	KillAllJobsOnFirstError       bool
	JobExitCode                   int
}
