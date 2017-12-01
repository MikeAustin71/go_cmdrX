package common

import (
	"time"
	"fmt"
)

const (
	srcFileNameStartUpParms = "startupparameters.go"
	errBlockNoStartUpParms  = int64(351230000)
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
	StartTimeUTC						time.Time
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


func (sUp * StartupParameters) AssembleAppPath(appPathAndFileName string, parent []ErrBaseInfo) (FileHelper, SpecErr) {
	se := startUpBaseErrConfig(parent, "AssembleAppPath")

	fh := FileHelper{}

	appPathAndFileName = fh.AdjustPathSlash(appPathAndFileName)

	appP, err := fh.MakeAbsolutePath(appPathAndFileName)
	
	if err != nil {
		s := fmt.Sprintf("sUp.AssembleAppPath() Error. fh.MakeAbsolutePath() Failed! App File Path = %v", appPathAndFileName)
		return FileHelper{}, se.New(s, err, true, 101)
		
	}

	if !fh.DoesFileExist(appP)  {
		s := fmt.Sprintf("sUp.AssembleAppPath() Error. App File Path = %v Does NOT EXIST!", appP)
		e := fmt.Errorf("Error. File: %v DOES NOT EXIST!", appP)
		return FileHelper{}, se.New(s, e, true, 102)
	}

	fh2, err := fh.GetPathFileNameElements(appP)

	if err != nil {
		s:= fmt.Sprintf("sUp.AssembleAppPath() Error - Failed to Get Path File Name Elements for App File: %v.", appPathAndFileName)
		return FileHelper{}, se.New(s, err, true, 103)
	}

	if !fh2.AbsolutePathIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleAppPath() Error - Failed to Populate Absolute Path for App File: %v.", appP)
		e:= fmt.Errorf("Absolute App Path is NOT Populated! App Path File Name = %v", appPathAndFileName)
		return FileHelper{}, se.New(s, e, true, 104)
	}

	if !fh2.FileNameExtIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleAppPath() Error - Failed to Populate File Name and Extension Path for App File: %v.", appP)
		e:= fmt.Errorf("File Name and Extension is NOT Populated! App Path File Name = %v", appPathAndFileName)
		return FileHelper{}, se.New(s, e, true, 105)
	}

	return fh2, se.SignalNoErrors()
}

func (sUp * StartupParameters) AssembleCmdPath(cmdPathAndFileName string, parent []ErrBaseInfo) (FileHelper, SpecErr) {

	se := startUpBaseErrConfig(parent, "AssembleCmdPath")

	fh := FileHelper{}

	cmdPathAndFileName = fh.AdjustPathSlash(cmdPathAndFileName)

	cmdPath, err := fh.MakeAbsolutePath(cmdPathAndFileName)

	if err != nil {
		s := fmt.Sprintf("sUp.AssembleCmdPath() Error. fh.MakeAbsolutePath() Failed! Cmd File Path = %v", cmdPathAndFileName)
		return FileHelper{}, se.New(s, err, true, 201)

	}

	if !fh.DoesFileExist(cmdPath)  {
		s := fmt.Sprintf("sUp.AssembleCmdPath() Error. Cmd File Path = %v Does NOT EXIST!", cmdPath)
		e := fmt.Errorf("Error. Cmd File: %v DOES NOT EXIST!", cmdPath)
		return FileHelper{}, se.New(s, e, true, 202)
	}

	fh2, err := fh.GetPathFileNameElements(cmdPath)

	if err != nil {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Get Path File Name Elements for Cmd File: %v.", cmdPathAndFileName)
		return FileHelper{}, se.New(s, err, true, 203)
	}

	if !fh2.AbsolutePathIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Populate Absolute Path for Cmd File: %v.", cmdPath)
		e:= fmt.Errorf("Absolute App Path is NOT Populated! App Path File Name = %v", cmdPathAndFileName)
		return FileHelper{}, se.New(s, e, true, 204)
	}

	if !fh2.FileNameExtIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Populate File Name and Extension Path for Cmd File: %v.", cmdPath)
		e:= fmt.Errorf("File Name and Extension is NOT Populated! App Path File Name = %v", cmdPathAndFileName)
		return FileHelper{}, se.New(s, e, true, 205)
	}

	return fh2, se.SignalNoErrors()
}

func (sUp * StartupParameters) AssembleLogPath(lgPath string, parent []ErrBaseInfo) (FileHelper, SpecErr) {

	se := startUpBaseErrConfig(parent, "AssembleLogPath")

	fh := FileHelper{}

	lgPath = fh.AdjustPathSlash(lgPath)

	logPath, err := fh.MakeAbsolutePath(lgPath)

	if err != nil {
		s := fmt.Sprintf("sUp.AssembleCmdPath() Error. fh.MakeAbsolutePath() Failed! App File Path = %v", lgPath)
		return FileHelper{}, se.New(s, err, true, 301)
	}

	fh2, err := fh.GetPathFileNameElements(logPath)

	if err != nil {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Get Path lements for Log File Path: %v.", lgPath)
		return FileHelper{}, se.New(s, err, true, 302)
	}

	if !fh2.AbsolutePathIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Populate Absolute Path for Cmd File: %v.", lgPath)
		e:= fmt.Errorf("Absolute App Path is NOT Populated! App Path File Name = %v", logPath)
		return FileHelper{}, se.New(s, e, true, 303)
	}

	return fh2, se.SignalNoErrors()
}

// baseLogErrConfig - Used internally by LogJobGroup
// methods to set up error messages.
func startUpBaseErrConfig(parent []ErrBaseInfo, funcName string) SpecErr {

	bi := ErrBaseInfo{}.New(srcFileNameStartUpParms, funcName, errBlockNoStartUpParms)

	return SpecErr{}.InitializeBaseInfo(parent, bi)
}
