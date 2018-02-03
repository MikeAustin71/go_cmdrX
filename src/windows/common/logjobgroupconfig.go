package common

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*  'logjobgroupconfig.go' is located in source code
		repository:

				https://github.com/MikeAustin71/logopsgo.git

 */

const (
	// LogGroupConfigSrcFile - Log Group Cfg source
	// code file name.
	LogGroupConfigSrcFile = "logjobgroupconfig.go"

	// LogGroupConfigErrBlockNo - The block of numbers
	// assigned as error codes for this source file.
	LogGroupConfigErrBlockNo = int64(647000000)
)

// LogJobGroup - holds logging configuration for the
// current group of jobs
type LogJobGroup struct {
	LogMode                 LoggingMode
	StartTimeUTC            time.Time
	StartTime               time.Time
	EndTimeUTC              time.Time
	EndTime                 time.Time
	Duration                time.Time
	Dtfmt                   *DateTimeFormatUtility
	CommandFileName         string
	AppName                 string
	AppExeFileName          string
	AppVersion              string
	AppPath                 string
	AppLogFileName          string
	AppLogDir               string
	AppLogPathFileName      string
	AppLogDirWalkInfo       DirWalkInfo
	FilePtr                 *os.File
	Banner1                 string
	Banner2                 string
	Banner3                 string
	Banner4                 string
	Banner5                 string
	Banner6                 string
	Banner7                 string
	BannerLen               int
	LeftTab                 string
	BaseStartDir            string
	AppNoOfJobs             int
	NoOfJobGroupMsgs        int
	NoOfJobsCompleted       int
	LogFileRetentionInDays  int
	NoOfLogFilesPurged      int
	CmdExeDir               string
	KillAllJobsOnFirstError bool
	IanaTimeZone            string
}

// New - Initializes key
// elements of a Logging Configuration
func (logOps *LogJobGroup) New(startParams StartupParameters,
																	parent []OpsMsgContextInfo) OpsMsgDto {

	var err error

	om := logOps.baseLogErrConfig(parent, "New()")

	// Assumes CreateAllFormatsInMemory() has
	// already been called.
	logOps.Dtfmt = startParams.Dtfmt

	logOps.LogMode = startParams.LogMode
	logOps.StartTimeUTC = startParams.StartTimeUTC
	logOps.StartTime = startParams.StartTime
	logOps.CommandFileName = startParams.CommandFileName
	logOps.AppName = startParams.AppName
	logOps.AppExeFileName = startParams.AppExeFileName
	logOps.AppPath = startParams.AppPath
	logOps.BaseStartDir = startParams.BaseStartDir
	logOps.AppVersion = startParams.AppVersion
	logOps.AppNoOfJobs = startParams.NoOfJobs
	logOps.LogFileRetentionInDays = startParams.LogFileRetentionInDays
	logOps.CmdExeDir = startParams.CmdExeDir
	logOps.KillAllJobsOnFirstError = startParams.KillAllJobsOnFirstError
	logOps.IanaTimeZone = startParams.IanaTimeZone

	fh := FileHelper{}

	if startParams.AppLogFileName == "" {
		logOps.AppLogFileName = logOps.AppName
	}

	dt := DateTimeUtility{}
	logOps.AppLogFileName = logOps.AppLogFileName + "_" + dt.GetDateTimeStr(logOps.StartTime) + ".log"

	logOps.AppLogDir, err = fh.MakeAbsolutePath(startParams.AppLogPath)

	if err != nil {
		s := fmt.Sprintf("MakeAbsolutePath Failed for AppLogPath '%v'", startParams.AppLogPath)
		om.SetFatalError(s, err, 101)
		return om
	}

	logOps.AppLogPathFileName = fh.JoinPathsAdjustSeparators(logOps.AppLogDir, logOps.AppLogFileName)

	logOps.BannerLen = 80

	logOps.Banner1 = strings.Repeat("#", logOps.BannerLen) + "\n"

	logOps.Banner2 = strings.Repeat("=", logOps.BannerLen) + "\n"

	logOps.Banner3 = strings.Repeat("*", logOps.BannerLen) + "\n"

	logOps.Banner4 = strings.Repeat("-", logOps.BannerLen) + "\n"

	logOps.Banner5 = strings.Repeat("!", logOps.BannerLen) + "\n"

	logOps.Banner6 = strings.Repeat("&", logOps.BannerLen) + "\n"

	logOps.Banner7 = strings.Repeat("+", logOps.BannerLen) + "\n"

	logOps.LeftTab = strings.Repeat(" ", 2)

	logOps.BaseStartDir, err = fh.GetAbsCurrDir()

	if err != nil {
		s := "GetAbsCurrDir() Failed!"
		om.SetFatalError(s, err, 102)
		return om
	}

	return logOps.InitializeLogFile(om.GetNewParentHistory())

}

// InitializeLogFile - Creates the log directory and
// a new log file. Opens the Log File. New(..) MUST
// be called before this method!
func (logOps *LogJobGroup) InitializeLogFile(parent []OpsMsgContextInfo) OpsMsgDto {

	var err error
	om := logOps.baseLogErrConfig(parent, "InitializeLogFile()")

	fh := FileHelper{}

	if !fh.DoesFileExist(logOps.AppLogDir) {

		err = fh.MakeDirAll(logOps.AppLogDir)

		if err != nil {
			s := fmt.Sprintf("MakeDirAll() Failed on Dir: %v", logOps.AppLogDir)
			om.SetFatalError(s, err, 201)
			return om
		}

	}

	// At this point, logOps.AppLogPath exists.
	// Check to determine if the log file exists.
	// If log file already exists, delete it.
	if fh.DoesFileExist(logOps.AppLogPathFileName) {

		err = fh.DeleteDirFile(logOps.AppLogPathFileName)

		if err != nil {
			s := fmt.Sprintf("DeleteDirFile() Failed on File: %v", logOps.AppLogPathFileName)
			om.SetFatalError(s, err, 202)
			return om
		}
	}

	// Create a new log file
	logOps.FilePtr, err = fh.CreateFile(logOps.AppLogPathFileName)

	if err != nil {
		s := fmt.Sprintf("CreateFile() Failed on File: %v Error: %v", logOps.AppLogPathFileName, err.Error())
		om.SetFatalError(s, err, 203)
		return om
	}

	om2 := logOps.purgeOldLogFiles(om.GetNewParentHistory())

	if om2.IsFatalError() {

		return om2
	}

	return logOps.writeJobGroupHeaderToLog(om.GetNewParentHistory())
}

// purgeOldLogFiles - Deletes log files which are older than
// logOps.LogFileRetentionInDays
func (logOps *LogJobGroup) purgeOldLogFiles(parent []OpsMsgContextInfo) OpsMsgDto {

	om := logOps.baseLogErrConfig(parent, "purgeOldLogFiles()")

	fh := FileHelper{}

	if logOps.LogFileRetentionInDays < 0 {
		om.SetSuccessfulCompletionMessage("Finished purgeOldLogFiles - No Log Files To Delete.", 1801)
		return om
	}

	logDur := time.Duration(logOps.LogFileRetentionInDays*24*-1) * time.Hour
	du := DurationUtility{}
	du.SetStartTimeDuration(time.Now(), logDur)
	thresholdTime := du.StartDateTime

	if thresholdTime.IsZero() {
		om.SetSuccessfulCompletionMessage("Finished purgeOldLogFiles - thresholdTime is Zero.", 1802)
		return om
	}

	logOps.AppLogDirWalkInfo = DirWalkInfo{}

	logOps.AppLogDirWalkInfo.StartPath = logOps.AppLogDir
	logOps.AppLogDirWalkInfo.PatternMatch = "*.log"
	logOps.AppLogDirWalkInfo.DeleteFilesOlderThan = thresholdTime
	err := fh.WalkDirPurgeFilesOlderThan(&logOps.AppLogDirWalkInfo)
	if err != nil {
		s := fmt.Sprintf("WalkDirPurgeFilesOlderThan() Failed on File: %v Error: %v", logOps.AppLogPathFileName, err.Error())
		om.SetFatalError(s, err, 1803)
		return om
	}

	logOps.NoOfLogFilesPurged = len(logOps.AppLogDirWalkInfo.DeletedFiles)
	s:= fmt.Sprintf("Finished purgeOldLogFiles - Successfully Purged %v old Log Files.", logOps.NoOfLogFilesPurged)
	om.SetSuccessfulCompletionMessage(s,1809 )
	return om

}

// writeJobGroupHeaderToLog - Writes the Job Group
// Header info to the log file. This is a one-time
// operation for each log file. Method InitializeLogFile(..)
// MUST be called before first use of this method.
// The Header is always the first element in the Log.
func (logOps *LogJobGroup) writeJobGroupHeaderToLog(parent []OpsMsgContextInfo) OpsMsgDto {
	var err error
	var str, stx string

	om := logOps.baseLogErrConfig(parent, "writeJobGroupHeaderToLog()")

	thisParentInfo := om.GetNewParentHistory()

	if logOps.FilePtr == nil {
		s := "logOps.FilePtr was not correctly initialized! logOps.FilePtr *os.File pointer is nil!"
		om.SetFatalErrorMessage(s, 301)
		return om
	}

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	su := StringUtility{}
	stx = fmt.Sprintf("App Name: %v  AppVersion: %v \n", logOps.AppName, logOps.AppVersion)
	str, err = su.StrCenterInStrLeft(stx, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStrLeft threw error on AppName AppVersion"
		om.SetFatalError(s, err, 302)
		return om
	}

	logOps.writeFileStr(str, thisParentInfo)

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	str = fmt.Sprintf("Execution Job Group Command File: %v \n", logOps.CommandFileName)

	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command File Name"
		om.SetFatalError(s, err, 303)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Starting Execution of %v Jobs. \n", logOps.AppNoOfJobs)

	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Number of Jobs Executed"
		om.SetFatalError(s, err, 304)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	dt := DateTimeUtility{}

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("  Job Group Start Time UTC: %v \n", dt.GetDateTimeYMDAbbrvDowNano(logOps.StartTimeUTC))

	logOps.writeTabFileStr(str, 0, parent)

	str = fmt.Sprintf("Job Group Start Time Local: %v \n", dt.GetDateTimeYMDAbbrvDowNano(logOps.StartTime))

	logOps.writeTabFileStr(str, 0, parent)

	localZone, _ := logOps.StartTime.Zone()

	str = fmt.Sprintf(" Job Group Local Time Zone: %v - %v \n", logOps.IanaTimeZone, localZone)

	logOps.writeTabFileStr(str, 0, parent)

	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	str = "Initial Application Path: \n"

	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = logOps.AppPath + "\n\n"

	logOps.writeTabFileStr(str, 2, thisParentInfo)


	str = "This Log File:\n"

	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = logOps.AppLogPathFileName + "\n\n"

	logOps.writeTabFileStr(str, 2, thisParentInfo)

	str = "Base Start Directory:\n"
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = logOps.BaseStartDir + "\n"
	logOps.writeTabFileStr(str, 2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	str = fmt.Sprintf("Number Of Old Log Files Deleted: %v \n", logOps.NoOfLogFilesPurged)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	if logOps.NoOfLogFilesPurged > 0 {

		str = "!!!!!! Log Files Deleted !!!!!!\n"
		stx, _ = su.StrCenterInStrLeft(str, logOps.BannerLen)

		logOps.writeTabFileStr(stx, 1, thisParentInfo)
		logOps.writeFileStr(logOps.Banner4, thisParentInfo)
		for i := 0; i < logOps.NoOfLogFilesPurged; i++ {
			fi := logOps.AppLogDirWalkInfo.DeletedFiles[i]

			str = fmt.Sprintf("%v. File Date: %v   File Name: %v \n",
				i+1, fi.Info.ModTime().Format(FmtDateTimeSecText), fi.Info.Name())

			logOps.writeTabFileStr(str, 2, thisParentInfo)
		}

		logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	}

	logOps.AppLogDirWalkInfo = DirWalkInfo{}
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	// Signal Successful Completion
	om.SetSuccessfulCompletionMessage("Finished writeJobGroupHeaderToLog", 309)
	return om
}

// WriteJobGroupFooterToLog - Writes the trailing Job Group data
// to the Log File. This is the last entry in the Log. The File
// pointer is closed here.
func (logOps *LogJobGroup) WriteJobGroupFooterToLog(cmds CommandBatch, parent []OpsMsgContextInfo) OpsMsgDto {

	var err error
	var str, stx string

	om := logOps.baseLogErrConfig(parent, "WriteJobGroupFooterToLog()")

	thisParentInfo := om.GetNewParentHistory()

	if logOps.FilePtr == nil {
		s := "logOps.FilePtr was not correctly initialized! logOps.FilePtr *os.File pointer is nil!"
		om.SetFatalErrorMessage(s, 801)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	defer logOps.FilePtr.Close()

	logOps.writeFileStr("\n\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	su := StringUtility{}

	str = "Completed Job Group Execution\n"
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Job Group Execution Title"
		om.SetFatalError(s, err, 802)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("App Name: %v  AppVersion: %v \n", logOps.AppName, logOps.AppVersion)
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on AppName AppVersion"
		om.SetFatalError(s, err, 803)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("Execution Job Group Command File: %v \n", logOps.CommandFileName)

	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command File Name"
		om.SetFatalError(s, err, 804)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	stx = fmt.Sprintf("  Number of Jobs Executed: %v \n", logOps.NoOfJobsCompleted)

	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	stx =  fmt.Sprintf("Number of Messages Logged: %v \n", logOps.NoOfJobGroupMsgs)

	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	dt := DateTimeUtility{}

	stx = "Job Group Execution Times:\n"
	logOps.writeTabFileStr(stx, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	str = dt.GetDateTimeYMDAbbrvDowNano(logOps.StartTimeUTC)
	stx = fmt.Sprintf("JobGroup   Start Time UTC: %v \n", str)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.EndTimeUTC = cmds.CmdJobsHdr.CmdBatchEndUTC
	logOps.EndTime = cmds.CmdJobsHdr.CmdBatchEndTime
	str = dt.GetDateTimeYMDAbbrvDowNano(logOps.EndTimeUTC)
	stx = fmt.Sprintf("JobGroup     End Time UTC: %v \n", str)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	str = dt.GetDateTimeYMDAbbrvDowNano(logOps.StartTime)
	stx = fmt.Sprintf("JobGroup Start Time Local: %v \n", str)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	str = dt.GetDateTimeYMDAbbrvDowNano(logOps.EndTime)
	stx = fmt.Sprintf("JobGroup   End Time Local: %v \n", str)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	tzLocal, _ := logOps.EndTime.Zone()

	stx = fmt.Sprintf("JobGroup  Local Time Zone: %v - %v\n", logOps.IanaTimeZone, tzLocal)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	stx = "JobGroup Elapsed Time:\n"
	logOps.writeTabFileStr(stx, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner6, thisParentInfo)
	stx = cmds.CmdJobsHdr.CmdBatchElapsedTime + "\n"
	logOps.writeTabFileStr(stx, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner6, thisParentInfo)

	logOps.writeFileStr("\n\n", thisParentInfo)

	str = "End of Job Group Execution\n"
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on End of Job Group Execution"
		om.SetFatalError(s, err, 805)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	// Signal Successful Completion
	om.SetSuccessfulCompletionMessage("Finished WriteJobGroupFooterToLog", 809)

	return om
}

func (logOps *LogJobGroup) WriteCmdJobHeaderToLog(job *CmdJob, parent []OpsMsgContextInfo) OpsMsgDto {

	om := logOps.baseLogErrConfig(parent, "WriteCmdJobHeaderToLog()")

	thisParentInfo := om.GetNewParentHistory()

	su := StringUtility{}

	str := "\n\n"

	logOps.writeFileStr(str, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	str = "Starting Command Job Execution\n"
	stx, err := su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStrLeft threw error on Starting Command Job Execution"
		om.SetFatalError(s, err, 2501)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Command Job Name: %v\n", job.CmdDisplayName)
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Starting Command Job Display Name"
		om.SetFatalError(s, err, 2502)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Command Job Number: %v\n", job.CmdJobNo)
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStrLeft threw error on Starting Command Job Number"
		om.SetFatalError(s, err, 2503)
		om.LogMsgToFile(logOps.FilePtr)

		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("   Command Job Description: %v\n", job.CmdDescription)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	str = fmt.Sprintf("               Command Type: %v\n", job.CmdDescription)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("   Execute Cmd In Directory: %v\n", job.ExeCmdInDir)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("    Delay Cmd Start Seconds: %v\n", job.DelayCmdStartSeconds)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Delay Cmd Start Target Time: %v\n", job.DelayStartCmdDateTime)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("    Cmd Time Out in Seconds: %v\n", job.CommandTimeOutInSeconds)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("          Execution Command: %v\n", job.ExeCommand)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	largs := len(job.CmdArguments.CmdArgs)

	if largs < 1 {
    str = "          Command Arguments: NO COMMAND ARGUMENTS PROVIDED\n"
		logOps.writeTabFileStr(str, 1, thisParentInfo)

	}	else if len(job.CombinedArguments) < 50 {
		str = fmt.Sprintf("          Command Arguments: %v\n", job.CombinedArguments)
		logOps.writeTabFileStr(str, 1, thisParentInfo)
	} else {
		logOps.writeFileStr(logOps.Banner4, thisParentInfo)
		str = "Command Arguments:\n"
		logOps.writeTabFileStr(str, 1, thisParentInfo)
		logOps.writeFileStr(logOps.Banner4, thisParentInfo)
		for i:= 0; i < largs; i++ {
		 	str = job.CmdArguments.CmdArgs[i] + "\n"
			logOps.writeTabFileStr(str, 3, thisParentInfo)
		}

		logOps.writeFileStr(logOps.Banner4 + "\n", thisParentInfo)
	}

	largs = len(job.CmdInputs.InputArgs)

	if largs < 1 {
		str = "            Input Arguments: NO INPUT ARGUMENTS PROVIDED\n"
		logOps.writeTabFileStr(str, 1, thisParentInfo)

	} else if len(job.CombinedInputArguments) < 50 {
		str = fmt.Sprintf("            Input Arguments: %v\n", job.CombinedInputArguments)
		logOps.writeTabFileStr(str, 1, thisParentInfo)

	} else {

		str = "Input Arguments:\n"
		logOps.writeTabFileStr(str, 1, thisParentInfo)
		logOps.writeFileStr(logOps.Banner4, thisParentInfo)

		for i:= 0; i < largs; i++ {
			str = job.CmdInputs.InputArgs[i] + "\n"
			logOps.writeTabFileStr(str, 3, thisParentInfo)
		}

		logOps.writeTabFileStr("\n", 1, thisParentInfo)
	}


	logOps.writeFileStr(logOps.Banner3, thisParentInfo)


	str = "Command Job Start Times\n"
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command Job Start Times"
		om.SetFatalError(s, err, 2504)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	str = fmt.Sprintf("     Cmd Job Start Time UTC: %v\n", job.CmdJobStartUTC.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("   Cmd Job Start Time Local: %v\n", job.CmdJobStartTimeValue.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	tzLocal, _ := job.CmdJobStartTimeValue.Zone()
	str = fmt.Sprintf("    Cmd Job Local Time Zone: %v - %v\n", job.IanaTimeZone, tzLocal)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)


	om.SetSuccessfulCompletionMessage("Finished WriteCmdJobHeaderToLog", 2509)
	return om
}

func (logOps *LogJobGroup) WriteCmdJobFooterToLog(job *CmdJob, parent []OpsMsgContextInfo) OpsMsgDto {

	om := logOps.baseLogErrConfig(parent, "WriteCmdJobHeaderToLog()")
	thisParentInfo := om.GetNewParentHistory()
	su := StringUtility{}

	logOps.writeFileStr("\n\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	str := "Completed Command Job Execution\n"
	stx, err := su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Completed Command Job Execution"
		om.SetFatalError(s, err, 2601)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}
	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner5, thisParentInfo)
	str = "Job Execution Results\n"
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Completed Command Job Execution"
		om.SetFatalError(s, err, 2602)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Command Job Name: %v\n", job.CmdDisplayName)
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Completing Command Job Display Name"
		om.SetFatalError(s, err, 2603)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Command Job Number: %v\n", job.CmdJobNo)
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Completing Command Job Number"
		om.SetFatalError(s, err, 2604)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner5, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)


	str = fmt.Sprintf("Cmd Job Description: %v\n", job.CmdDescription)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("  Execution Command: %v\n", job.ExeCommand)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	largs := len(job.CmdArguments.CmdArgs)

	if largs < 1 {
		str = "          Command Arguments: NO COMMAND ARGUMENTS PROVIDED\n"
		logOps.writeTabFileStr(str, 1, thisParentInfo)

	}	else if len(job.CombinedArguments) < 50 {
		str = fmt.Sprintf("          Command Arguments: %v\n", job.CombinedArguments)
		logOps.writeTabFileStr(str, 1, thisParentInfo)
	} else {
		logOps.writeFileStr(logOps.Banner4, thisParentInfo)
		str = "Command Arguments:\n"
		logOps.writeTabFileStr(str, 1, thisParentInfo)
		logOps.writeFileStr(logOps.Banner4, thisParentInfo)
		for i:= 0; i < largs; i++ {
			str = job.CmdArguments.CmdArgs[i] + "\n"
			logOps.writeTabFileStr(str, 3, thisParentInfo)
		}

		logOps.writeFileStr(logOps.Banner4 + "\n", thisParentInfo)
	}

	largs = len(job.CmdInputs.InputArgs)

	if largs < 1 {
		str = "            Input Arguments: NO INPUT ARGUMENTS PROVIDED\n"
		logOps.writeTabFileStr(str, 1, thisParentInfo)

	} else if len(job.CombinedInputArguments) < 50 {
		str = fmt.Sprintf("            Input Arguments: %v\n", job.CombinedInputArguments)
		logOps.writeTabFileStr(str, 1, thisParentInfo)

	} else {

		str = "Input Arguments:\n"
		logOps.writeTabFileStr(str, 1, thisParentInfo)
		logOps.writeFileStr(logOps.Banner4, thisParentInfo)

		for i:= 0; i < largs; i++ {
			str = job.CmdInputs.InputArgs[i] + "\n"
			logOps.writeTabFileStr(str, 3, thisParentInfo)
		}

		logOps.writeTabFileStr("\n", 1, thisParentInfo)
	}

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	str = "Command Job Execution Times\n"
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command Job Execution Times"
		om.SetFatalError(s, err, 2605)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	logOps.writeTabFileStr("UTC Start\\End Times:\n", 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	str = fmt.Sprintf(" Cmd Job Start Time UTC: %v\n", job.CmdJobStartUTC.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("   Cmd Job End Time UTC: %v\n", job.CmdJobEndUTC.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	logOps.writeTabFileStr("Local Start\\End Times:\n", 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	str = fmt.Sprintf("Cmd Job Start Time Local: %v\n", job.CmdJobStartTimeValue.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("  Cmd Job End Time Local: %v\n", job.CmdJobEndTimeValue.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	tzLocal, _ := job.CmdJobEndTimeValue.Zone()
	str = fmt.Sprintf(" Cmd Job Local Time Zone: %v - %v\n", job.IanaTimeZone, tzLocal)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	str = "Cmd Job Execution Elapsed Time:\n"
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner6, thisParentInfo)

	str = fmt.Sprintf("%v\n", job.CmdJobElapsedTime)
	logOps.writeTabFileStr(str, 2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner6 + "\n", thisParentInfo)

	str = fmt.Sprintf("Number of Messages: %v\n", job.CmdJobNoOfMsgs)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("End of Command Job Number: %v\n", job.CmdJobNo)
	stx, err = su.StrCenterInStrLeft(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on End of Command Job Number"
		om.SetFatalError(s, err, 2606)
		om.LogMsgToFile(logOps.FilePtr)
		return om
	}

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr("\n\n", thisParentInfo)

	job.CmdJobIsCompleted = true

	om.SetSuccessfulCompletionMessage("Finished WriteCmdJobFooterToLog", 2609)
	return om

}

func (logOps *LogJobGroup) WriteOpsMsgToLog(opsMsgToLog OpsMsgDto, job *CmdJob, parent []OpsMsgContextInfo) OpsMsgDto {


	om := logOps.baseLogErrConfig(parent, "WriteOpsMsgToLog()")

	thisParentInfo := om.GetNewParentHistory()

	logOps.writeTabFileStr(opsMsgToLog.String(), 2, thisParentInfo)

	om.SetSuccessfulCompletionMessage("Finished WriteOpsMsgToLog", 2709)

	return om
}

func (logOps *LogJobGroup) writeTabFileStr(s string, noOfTabs int, parent []OpsMsgContextInfo) {

	om := logOps.baseLogErrConfig(parent, "writeTabFileStr()")

	stx := ""

	for i := 0; i < noOfTabs; i++ {
		stx += logOps.LeftTab
	}

	stx += s

	_, err := logOps.FilePtr.WriteString(stx)

	if err != nil {
		s := fmt.Sprintf("FilePtr.WriteString threw error on string: '%v'", stx)
		om.SetFatalError(s, err, 1001)
		panic(om)
	}

}

func (logOps *LogJobGroup) writeFileStr(s string, parent []OpsMsgContextInfo) {

	_, err := logOps.FilePtr.WriteString(s)

	om := logOps.baseLogErrConfig(parent, "writeFileStr()")

	if err != nil {
		s := fmt.Sprintf("FilePtr.WriteString threw error on string: '%v'", s)
		om.SetFatalError(s, err, 901)
		panic(om)
	}

}

// baseLogErrConfig - Used internally by LogJobGroup
// methods to set up error messages.
func (logOps *LogJobGroup) baseLogErrConfig(parent []OpsMsgContextInfo, funcName string) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
							SourceFileName: LogGroupConfigSrcFile,
							ParentObjectName: "LogJobGroup",
							FuncName: funcName,
							BaseMessageId: LogGroupConfigErrBlockNo,
						}

	return OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)
}
