package common

import (
	"time"
	"fmt"
)

/*  'startupparameters.go' is located in source code
		repository:

		https://github.com/MikeAustin71/go_cmdrX.git

*/


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


func (sUp *StartupParameters) AssembleAppPath(appPathAndFileName string, parentHistory [] OpsMsgContextInfo) (FileHelper, OpsMsgDto) {

	om := sUp.startUpBaseErrConfig(parentHistory, "AssembleAppPath")

	fh := FileHelper{}

	appPathAndFileName = fh.AdjustPathSlash(appPathAndFileName)

	appP, err := fh.MakeAbsolutePath(appPathAndFileName)
	
	if err != nil {
		s := fmt.Sprintf("sUp.AssembleAppPath() Error. fh.MakeAbsolutePath() Failed! App File Path = %v", appPathAndFileName)
		om.SetFatalError(s, err, 101)
		return FileHelper{}, om
		
	}

	if !fh.DoesFileExist(appP)  {
		s := fmt.Sprintf("sUp.AssembleAppPath() Error. App File Path = %v Does NOT EXIST!", appP)
		e := fmt.Errorf("Error. File: %v DOES NOT EXIST!", appP)
		om.SetFatalError(s, e, 102)
		return FileHelper{}, om
	}

	fh2, err := fh.GetPathFileNameElements(appP)

	if err != nil {
		s:= fmt.Sprintf("sUp.AssembleAppPath() Error - Failed to Get Path File Name Elements for App File: %v.", appPathAndFileName)
		om.SetFatalError(s, err, 103)
		return FileHelper{}, om
	}

	if !fh2.AbsolutePathIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleAppPath() Error - Failed to Populate Absolute Path for App File: %v.", appP)
		e:= fmt.Errorf("Absolute App Path is NOT Populated! App Path File Name = %v", appPathAndFileName)
		om.SetFatalError(s, e, 104)
		return FileHelper{}, om
	}

	if !fh2.FileNameExtIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleAppPath() Error - Failed to Populate File Name and Extension Path for App File: %v.", appP)
		e:= fmt.Errorf("File Name and Extension is NOT Populated! App Path File Name = %v", appPathAndFileName)
		om.SetFatalError(s, e, 105)
		return FileHelper{}, om
	}

	om.SetSuccessfulCompletionMessage("Finished AssembleAppPath Setup", 109)

	return fh2, om
}

func (sUp * StartupParameters) AssembleCmdPath(cmdPathAndFileName string, parent []OpsMsgContextInfo) (FileHelper, OpsMsgDto) {

	om := sUp.startUpBaseErrConfig(parent, "AssembleCmdPath")

	fh := FileHelper{}

	cmdPathAndFileName = fh.AdjustPathSlash(cmdPathAndFileName)

	cmdPath, err := fh.MakeAbsolutePath(cmdPathAndFileName)

	if err != nil {
		s := fmt.Sprintf("sUp.AssembleCmdPath() Error. fh.MakeAbsolutePath() Failed! Cmd File Path = %v", cmdPathAndFileName)
		om.SetFatalError(s, err, 201)
		return FileHelper{}, om

	}

	if !fh.DoesFileExist(cmdPath)  {
		s := fmt.Sprintf("sUp.AssembleCmdPath() Error. Cmd File Path = %v Does NOT EXIST!", cmdPath)
		e := fmt.Errorf("Error. Cmd File: %v DOES NOT EXIST!", cmdPath)
		om.SetFatalError(s, e, 202)
		return FileHelper{}, om
	}

	fh2, err := fh.GetPathFileNameElements(cmdPath)

	if err != nil {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Get Path File Name Elements for Cmd File: %v.", cmdPathAndFileName)
		om.SetFatalError(s, err, 203)
		return FileHelper{}, om
	}

	if !fh2.AbsolutePathIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Populate Absolute Path for Cmd File: %v.", cmdPath)
		e:= fmt.Errorf("Absolute App Path is NOT Populated! App Path File Name = %v", cmdPathAndFileName)
		om.SetFatalError(s, e, 204)
		return FileHelper{}, om
	}

	if !fh2.FileNameExtIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Populate File Name and Extension Path for Cmd File: %v.", cmdPath)
		e:= fmt.Errorf("File Name and Extension is NOT Populated! App Path File Name = %v", cmdPathAndFileName)
		om.SetFatalError(s, e, 205)
		return FileHelper{}, om
	}

	om.SetSuccessfulCompletionMessage("Finished AssembleCmdPath", 209)

	return fh2, om
}

func (sUp *StartupParameters) AssembleLogPath(lgPath string, parent []OpsMsgContextInfo) (FileHelper, OpsMsgDto) {

	om := sUp.startUpBaseErrConfig(parent, "AssembleLogPath")

	fh := FileHelper{}

	lgPath = fh.AdjustPathSlash(lgPath)

	logPath, err := fh.MakeAbsolutePath(lgPath)

	if err != nil {
		s := fmt.Sprintf("sUp.AssembleCmdPath() Error. fh.MakeAbsolutePath() Failed! App File Path = %v", lgPath)
		om.SetFatalError(s, err, 301)
		return FileHelper{}, om
	}

	fh2, err := fh.GetPathFileNameElements(logPath)

	if err != nil {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Get Path lements for Log File Path: %v.", lgPath)
		om.SetFatalError(s, err, 302)
		return FileHelper{}, om
	}

	if !fh2.AbsolutePathIsPopulated {
		s:= fmt.Sprintf("sUp.AssembleCmdPath() Error - Failed to Populate Absolute Path for Cmd File: %v.", lgPath)
		e:= fmt.Errorf("Absolute App Path is NOT Populated! App Path File Name = %v", logPath)
		om.SetFatalError(s, e, 303)
		return FileHelper{}, om
	}

	om.SetSuccessfulCompletionMessage("Finished AssembleLogPath", 309)

	return fh2, om
}

// baseLogErrConfig - Used internally by LogJobGroup
// methods to set up error messages.
func (sUp *StartupParameters) startUpBaseErrConfig(parent []OpsMsgContextInfo, funcName string) OpsMsgDto {

	// bi := ErrBaseInfo{}.New(srcFileNameStartUpParms, funcName, errBlockNoStartUpParms)
	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNameStartUpParms,
							ParentObjectName: "StartupParameters",
							FuncName: funcName,
							BaseMessageId: errBlockNoStartUpParms,
						}

	return OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)
}
