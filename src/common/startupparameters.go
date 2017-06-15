package common

import (
	"time"
	"fmt"
)

// StartupParameters - This data structure is
// used by the calling method to initialize a
// LogJobGroup using the 'New(...)' method.
type StartupParameters struct {
	AppName                 string
	AppExeFileName          string
	AppPath									string
	BaseStartDir						string
	AppVersion              string
	CommandFileName         string
	StartTime               time.Time
	AppLogPath              string
	AppLogFileName          string
	NoOfJobs                int
	LogFileRetentionInDays  int
	CmdExeDir               string
	KillAllJobsOnFirstError bool
	IanaTimeZone            string
	LogMode                 LoggingMode
	Dtfmt										* DateTimeFormatUtility
}


func (sUp * StartupParameters) AssembleAppPath(cmdFilePath string) (FileHelper, error) {
	fh := FileHelper{}

	if ! fh.DoesFileExist(cmdFilePath) {
		return fh, fmt.Errorf("sUp.AssembleAppPath() Error. Cmd File Path = %v Does NOT EXIST!", cmdFilePath, )
	}

	fh2, err := fh.GetPathFileNameElements(cmdFilePath)

	if err != nil {
		return fh, fmt.Errorf("sUp.AssembleAppPath() Error.Failed to Get Path File Name Elements for Command File: %v.  Error = ", cmdFilePath, err.Error())
	}

	return fh2, nil
}

