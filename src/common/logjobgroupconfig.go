package common

import (
	"fmt"
	"os"
	"strings"
	"time"
)

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
	parent []ErrBaseInfo) SpecErr {

	var err error
	var isPanic bool

	se := logOps.baseLogErrConfig(parent, "New()")

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
		isPanic = true
		return se.New(s, err, isPanic, 101)
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
		isPanic = true
		return se.New(s, err, isPanic, 102)
	}

	return logOps.InitializeLogFile(se.AddBaseToParentInfo())

}

// InitializeLogFile - Creates the log directory and
// a new log file. Opens the Log File. New(..) MUST
// be called before this method!
func (logOps *LogJobGroup) InitializeLogFile(parent []ErrBaseInfo) SpecErr {

	var err error
	var isPanic bool
	se := logOps.baseLogErrConfig(parent, "InitializeLogFile()")

	fh := FileHelper{}

	if !fh.DoesFileExist(logOps.AppLogDir) {

		err = fh.MakeDirAll(logOps.AppLogDir)

		if err != nil {
			s := fmt.Sprintf("MakeDirAll() Failed on Dir: %v", logOps.AppLogDir)
			isPanic = true
			return se.New(s, err, isPanic, 201)
		}

	}

	// At this point, logOps.AppLogPath exists.
	// Check to determine if the log file exists.
	// If log file already exists, delete it.
	if fh.DoesFileExist(logOps.AppLogPathFileName) {

		err = fh.DeleteDirFile(logOps.AppLogPathFileName)

		if err != nil {
			s := fmt.Sprintf("DeleteDirFile() Failed on File: %v", logOps.AppLogPathFileName)
			isPanic = true
			return se.New(s, err, isPanic, 202)
		}
	}

	// Create a new log file
	logOps.FilePtr, err = fh.CreateFile(logOps.AppLogPathFileName)

	if err != nil {
		s := fmt.Sprintf("CreateFile() Failed on File: %v Error: %v", logOps.AppLogPathFileName, err.Error())
		isPanic = true
		return se.New(s, err, isPanic, 203)
	}

	si := logOps.purgeOldLogFiles(se.AddBaseToParentInfo())

	if si.IsErr {
		si.IsPanic = true
		return si
	}

	return logOps.writeJobGroupHeaderToLog(se.AddBaseToParentInfo())
}

// purgeOldLogFiles - Deletes log files which are older than
// logOps.LogFileRetentionInDays
func (logOps *LogJobGroup) purgeOldLogFiles(parent []ErrBaseInfo) SpecErr {
	se := logOps.baseLogErrConfig(parent, "purgeOldLogFiles()")

	fh := FileHelper{}

	if logOps.LogFileRetentionInDays < 0 {
		return se.SignalNoErrors()
	}

	logDur := time.Duration(logOps.LogFileRetentionInDays*24*-1) * time.Hour
	du := DurationUtility{}
	du.SetStartTimeDuration(time.Now(), logDur)
	thresholdTime := du.StartDateTime

	if thresholdTime.IsZero() {
		return se.SignalNoErrors()
	}

	logOps.AppLogDirWalkInfo = DirWalkInfo{}

	logOps.AppLogDirWalkInfo.StartPath = logOps.AppLogDir
	logOps.AppLogDirWalkInfo.PatternMatch = "*.log"
	logOps.AppLogDirWalkInfo.DeleteFilesOlderThan = thresholdTime
	err := fh.WalkDirPurgeFilesOlderThan(&logOps.AppLogDirWalkInfo)
	if err != nil {
		s := fmt.Sprintf("WalkDirPurgeFilesOlderThan() Failed on File: %v Error: %v", logOps.AppLogPathFileName, err.Error())
		return se.New(s, err, true, 1801)
	}

	logOps.NoOfLogFilesPurged = len(logOps.AppLogDirWalkInfo.DeletedFiles)

	return se.SignalNoErrors()

}

// writeJobGroupHeaderToLog - Writes the Job Group
// Header info to the log file. This is a one-time
// operation for each log file. Method InitializeLogFile(..)
// MUST be called before first use of this method.
// The Header is always the first element in the Log.
func (logOps *LogJobGroup) writeJobGroupHeaderToLog(parent []ErrBaseInfo) SpecErr {
	var err error
	var isPanic bool
	var str, stx string

	se := logOps.baseLogErrConfig(parent, "writeJobGroupHeaderToLog()")

	thisParentInfo := se.AddBaseToParentInfo()

	if logOps.FilePtr == nil {
		s := "logOps.FilePtr was not correctly initialized! logOps.FilePtr *os.File pointer is nil!"
		isPanic = true
		return se.New(s, err, isPanic, 301)
	}

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	su := StringUtility{}
	stx = fmt.Sprintf("App Name: %v  AppVersion: %v \n", logOps.AppName, logOps.AppVersion)
	str, err = su.StrCenterInStr(stx, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on AppName AppVersion"
		isPanic = true
		return se.New(s, err, isPanic, 302)
	}

	logOps.writeFileStr(str, thisParentInfo)

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	str = fmt.Sprintf("Execution Job Group Command File: %v \n", logOps.CommandFileName)

	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command File Name"
		isPanic = true
		return se.New(s, err, isPanic, 303)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Starting Execution of %v Jobs. \n", logOps.AppNoOfJobs)

	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Number of Jobs Executed"
		isPanic = true
		return se.New(s, err, isPanic, 304)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	dt := DateTimeUtility{}

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("  Job Group Start Time UTC: %v \n", dt.GetDateTimeEverything(logOps.StartTimeUTC))

	logOps.writeTabFileStr(str, 0, parent)

	str = fmt.Sprintf("Job Group Start Time Local: %v \n", dt.GetDateTimeEverything(logOps.StartTime))

	logOps.writeTabFileStr(str, 0, parent)

	str = fmt.Sprintf(" Job Group Local Time Zone: %v \n", logOps.IanaTimeZone)

	logOps.writeTabFileStr(str, 0, parent)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)

	str = "Initial Application Path: \n"

	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = logOps.AppPath + "\n"

	logOps.writeTabFileStr(str, 2, thisParentInfo)

	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)

	str = "This Log File:\n"

	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = logOps.AppLogPathFileName + "\n"

	logOps.writeTabFileStr(str, 2, thisParentInfo)

	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)

	str = "Base Start Directory:\n"
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = logOps.BaseStartDir + "\n"
	logOps.writeTabFileStr(str, 2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	logOps.writeFileStr("\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("Number Of Old Log Files Deleted: %v \n", logOps.NoOfLogFilesPurged)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	if logOps.NoOfLogFilesPurged > 0 {

		str = "!!!!!! Log Files Deleted !!!!!!\n"
		stx, _ = su.StrCenterInStr(str, logOps.BannerLen)

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
	return se.SignalNoErrors()
}

// WriteJobGroupFooterToLog - Writes the trailing Job Group data
// to the Log File. This is the last entry in the Log. The File
// pointer is closed here.
func (logOps *LogJobGroup) WriteJobGroupFooterToLog(cmds CommandBatch, parent []ErrBaseInfo) SpecErr {

	var err error
	var isPanic bool
	var str, stx string

	se := logOps.baseLogErrConfig(parent, "WriteJobGroupFooterToLog()")

	thisParentInfo := se.AddBaseToParentInfo()

	if logOps.FilePtr == nil {
		s := "logOps.FilePtr was not correctly initialized! logOps.FilePtr *os.File pointer is nil!"
		isPanic = true
		return se.New(s, err, isPanic, 801)
	}

	defer logOps.FilePtr.Close()

	logOps.writeFileStr("\n\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	su := StringUtility{}

	str = "Completed Job Group Execution\n"
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Job Group Execution Title"
		isPanic = true
		return se.New(s, err, isPanic, 802)
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("App Name: %v  AppVersion: %v \n", logOps.AppName, logOps.AppVersion)
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on AppName AppVersion"
		isPanic = true
		return se.New(s, err, isPanic, 803)
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("Execution Job Group Command File: %v \n", logOps.CommandFileName)

	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command File Name"
		isPanic = true
		return se.New(s, err, isPanic, 804)
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	stx = fmt.Sprintf("  Number of Jobs Executed: %v \n", logOps.NoOfJobsCompleted)

	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	stx = logOps.LeftTab + fmt.Sprintf("Number of Messages Logged: %v \n", logOps.NoOfJobGroupMsgs)

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	dt := DateTimeUtility{}

	logOps.writeFileStr("\n", thisParentInfo)
	stx = "Job Group Execution Times:\n"
	logOps.writeTabFileStr(stx, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	str = dt.GetDateTimeEverything(logOps.StartTimeUTC)
	stx = fmt.Sprintf("JobGroup   Start Time UTC: %v \n", str)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.EndTimeUTC = cmds.CmdJobsHdr.CmdBatchEndUTC
	logOps.EndTime = cmds.CmdJobsHdr.CmdBatchEndTime
	str = dt.GetDateTimeEverything(logOps.EndTimeUTC)
	stx = fmt.Sprintf("JobGroup     End Time UTC: %v \n", str)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	str = dt.GetDateTimeEverything(logOps.StartTime)
	stx = fmt.Sprintf("JobGroup Start Time Local: %v \n", str)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	str = dt.GetDateTimeEverything(logOps.EndTime)
	stx = fmt.Sprintf("JobGroup   End Time Local: %v \n", str)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	stx = fmt.Sprintf("JobGroup  Local Time Zone: %v \n", logOps.IanaTimeZone)
	logOps.writeTabFileStr(stx, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)

	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	stx = "JobGroup Elapsed Time:\n"
	logOps.writeTabFileStr(stx, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	stx = cmds.CmdJobsHdr.CmdBatchElapsedTime + "\n"
	logOps.writeTabFileStr(stx, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	str = "End of Job Group Execution\n"
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on End of Job Group Execution"
		isPanic = true
		return se.New(s, err, isPanic, 805)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	// Signal Successful Completion
	return se.SignalNoErrors()
}

func (logOps *LogJobGroup) WriteCmdJobHeaderToLog(job *CmdJob, parent []ErrBaseInfo) SpecErr {

	se := logOps.baseLogErrConfig(parent, "WriteCmdJobHeaderToLog()")
	thisParentInfo := se.AddBaseToParentInfo()
	su := StringUtility{}
	isPanic := false
	str := "\n\n"

	logOps.writeFileStr(str, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	str = "Starting Command Job Execution\n"
	stx, err := su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Starting Command Job Execution"
		isPanic = true
		return se.New(s, err, isPanic, 2501)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Command Job Name: %v\n", job.CmdDisplayName)
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Starting Command Job Display Name"
		isPanic = true
		return se.New(s, err, isPanic, 2502)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Command Job Number: %v\n", job.CmdJobNo)
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Starting Command Job Number"
		isPanic = true
		return se.New(s, err, isPanic, 2503)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	logOps.writeFileStr("\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	str = fmt.Sprintf("Cmd Job Description: %v\n", job.CmdDescription)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	str = fmt.Sprintf("Cmd Type: %v\n", job.CmdDescription)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Execute Cmd In Directory: %v\n", job.ExeCmdInDir)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Delay Cmd Start Seconds: %v\n", job.DelayCmdStartSeconds)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Delay Cmd Start Target Time: %v\n", job.DelayStartCmdDateTime)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Delay Cmd Time Out in Seconds: %v\n", job.CommandTimeOutInSeconds)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)

	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	str = fmt.Sprintf("Execution Command: %v\n", job.ExeCommand)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Command Arguments: %v\n", job.CombinedArguments)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	logOps.writeFileStr("\n", thisParentInfo)

	str = "Command Job Start Times\n"
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command Job Start Times"
		isPanic = true
		return se.New(s, err, isPanic, 2504)
	}
	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	str = fmt.Sprintf("  Cmd Job Start Time UTC: %v\n", job.CmdJobStartUTC.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Cmd Job Start Time Local: %v\n", job.CmdJobStartTimeValue.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf(" Cmd Job Local Time Zone: %v\n", job.IanaTimeZone)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr("\n\n", thisParentInfo)

	return se.SignalNoErrors()
}

func (logOps *LogJobGroup) WriteCmdJobFooterToLog(job *CmdJob, parent []ErrBaseInfo) SpecErr {

	se := logOps.baseLogErrConfig(parent, "WriteCmdJobHeaderToLog()")
	thisParentInfo := se.AddBaseToParentInfo()
	su := StringUtility{}
	isPanic := false

	logOps.writeFileStr("\n\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner5, thisParentInfo)
	str := "Completed Command Job Execution\n"
	stx, err := su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Completed Command Job Execution"
		isPanic = true
		return se.New(s, err, isPanic, 2601)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Command Job Name: %v\n", job.CmdDisplayName)
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Completing Command Job Display Name"
		isPanic = true
		return se.New(s, err, isPanic, 2602)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Command Job Number: %v\n", job.CmdJobNo)
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Completing Command Job Number"
		isPanic = true
		return se.New(s, err, isPanic, 2603)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner5, thisParentInfo)
	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	logOps.writeFileStr("\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	str = fmt.Sprintf("Cmd Job Description: %v\n", job.CmdDescription)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	str = fmt.Sprintf("Execution Command: %v\n", job.ExeCommand)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Command Arguments: %v\n", job.CombinedArguments)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	logOps.writeFileStr(logOps.Banner3, thisParentInfo)
	str = "Command Job Execution Times\n"
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command Job Execution Times"
		isPanic = true
		return se.New(s, err, isPanic, 2604)
	}
	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeTabFileStr("UTC Start\\End Times:\n", 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	str = fmt.Sprintf(" Cmd Job Start Time UTC: %v\n", job.CmdJobStartUTC.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("   Cmd Job End Time UTC: %v\n", job.CmdJobEndUTC.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	logOps.writeFileStr("\n", thisParentInfo)

	logOps.writeTabFileStr("Local Start\\End Times:\n", 1, thisParentInfo)

	logOps.writeFileStr(logOps.Banner4, thisParentInfo)

	str = fmt.Sprintf("Cmd Job Start Time Local: %v\n", job.CmdJobStartTimeValue.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("  Cmd Job End Time Local: %v\n", job.CmdJobEndTimeValue.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf(" Cmd Job Local Time Zone: %v\n", job.IanaTimeZone)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)

	logOps.writeFileStr(logOps.Banner6, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)

	str = "Cmd Job Execution Elapsed Time:\n"
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner6, thisParentInfo)

	str = fmt.Sprintf("%v\n", job.CmdJobElapsedTime)
	logOps.writeTabFileStr(str, 2, thisParentInfo)
	logOps.writeFileStr(logOps.Banner6, thisParentInfo)
	logOps.writeFileStr("\n\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	str = fmt.Sprintf("Number of Messages: %v\n", job.CmdJobNoOfMsgs)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner4, thisParentInfo)
	logOps.writeFileStr("\n", thisParentInfo)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)

	str = fmt.Sprintf("End of Command Job Number: %v\n", job.CmdJobNo)
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on End of Command Job Number"
		isPanic = true
		return se.New(s, err, isPanic, 2605)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner2, thisParentInfo)
	logOps.writeFileStr("\n\n", thisParentInfo)

	job.CmdJobIsCompleted = true

	return se.SignalNoErrors()

}

func (logOps *LogJobGroup) WriteOpsMsgToLog(opsMsg OpsMsgDto, job *CmdJob, parent []ErrBaseInfo) SpecErr {

	job.CmdJobNoOfMsgs++

	se := logOps.baseLogErrConfig(parent, "WriteOpsMsgToLog()")
	thisParentInfo := se.AddBaseToParentInfo()
	su := StringUtility{}
	isPanic := false
	opsMsg.SetTime(job.IanaTimeZone)

	logOps.writeFileStr("\n\n", thisParentInfo)
	logOps.writeFileStr(logOps.Banner7, thisParentInfo)
	logOps.writeFileStr(logOps.Banner7, thisParentInfo)

	str := "Job Execution Message\n"
	stx, err := su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command Job Execution Message"
		isPanic = true
		return se.New(s, err, isPanic, 2701)
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner5, thisParentInfo)
	str = fmt.Sprintf("Command Job Name: %v\n", job.CmdDisplayName)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	str = fmt.Sprintf("Command Job Number: %v\n", job.CmdJobNo)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Command Message Number: %v\n", job.CmdJobNoOfMsgs)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Command Message Type: %v\n", opsMsg.MsgType)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Command Message Level: %v\n", opsMsg.MsgLevel)
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("  Command Message Time UTC: %v\n", opsMsg.MsgTimeUTC.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Command Message Time Local: %v\n", opsMsg.MsgTimeLocal.Format(job.CmdJobTimeFormat))
	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf(" Command Message Time Zone: %v\n", opsMsg.LocalTimeZone)
	logOps.writeTabFileStr(str, 1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner5, thisParentInfo)

	lenMsgs := len(opsMsg.Message)

	for i := 0; i < lenMsgs; i++ {
		if i == 0 {
			str = "Command Message:\n"
			logOps.writeTabFileStr(str, 1, thisParentInfo)
			logOps.writeFileStr(logOps.Banner6, thisParentInfo)
			logOps.writeFileStr("\n", thisParentInfo)
		}

		str = opsMsg.Message[i]
		logOps.writeTabFileStr(str, 2, thisParentInfo)

		if i == (lenMsgs - 1) {
			logOps.writeFileStr("\n", thisParentInfo)
			logOps.writeFileStr(logOps.Banner6, thisParentInfo)
			logOps.writeFileStr("\n", thisParentInfo)
		}

	}

	logOps.writeFileStr(logOps.Banner7, thisParentInfo)
	str = fmt.Sprintf("End Of Job Execution Message Number: %v\n",	job.CmdJobNoOfMsgs)
	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command Job Execution Message"
		isPanic = true
		return se.New(s, err, isPanic, 2701)
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner7, thisParentInfo)
	logOps.writeFileStr("\n\n", thisParentInfo)

	return se.SignalNoErrors()
}

func (logOps *LogJobGroup) writeTabFileStr(s string, noOfTabs int, parent []ErrBaseInfo) {

	stx := ""

	for i := 0; i < noOfTabs; i++ {
		stx += logOps.LeftTab
	}

	stx += s

	_, err := logOps.FilePtr.WriteString(stx)

	if err != nil {
		s := fmt.Sprintf("FilePtr.WriteString threw error on string: '%v'", stx)
		se :=
			logOps.baseLogErrConfig(parent, "writeTabFileStr()").
				New(s, err, true, 1001)
		panic(se)
	}

}

func (logOps *LogJobGroup) writeFileStr(s string, parent []ErrBaseInfo) {

	_, err := logOps.FilePtr.WriteString(s)

	if err != nil {
		s := fmt.Sprintf("FilePtr.WriteString threw error on string: '%v'", s)
		se :=
			logOps.baseLogErrConfig(parent, "writeFileStr()").
				New(s, err, true, 901)
		panic(se)
	}

}

// baseLogErrConfig - Used internally by LogJobGroup
// methods to set up error messages.
func (logOps *LogJobGroup) baseLogErrConfig(parent []ErrBaseInfo, funcName string) SpecErr {

	bi := ErrBaseInfo{}.New(LogGroupConfigSrcFile, funcName, LogGroupConfigErrBlockNo)

	return SpecErr{}.InitializeBaseInfo(parent, bi)
}
