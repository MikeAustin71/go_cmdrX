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
	StartTime               time.Time
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
	BannerLen               int
	LeftTab                 string
	BaseStartDir            string
	AppNoOfJobs             int
	NoOfJobGroupMsgs        int
	NoOfJobMsgs             int
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

	logOps.BannerLen = 79

	logOps.Banner1 = strings.Repeat("#", logOps.BannerLen) + "\n"

	logOps.Banner2 = strings.Repeat("=", logOps.BannerLen) + "\n"

	logOps.Banner3 = strings.Repeat("*", logOps.BannerLen) + "\n"

	logOps.Banner4 = strings.Repeat("-", logOps.BannerLen) + "\n"

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

	return logOps.writeFileGroupHeaderToLog(se.AddBaseToParentInfo())
}

// purgeOldLogFiles - Deletes log files which are older than
// logOps.LogFileRetentionInDays
func (logOps *LogJobGroup) purgeOldLogFiles(parent []ErrBaseInfo) SpecErr {
	se := logOps.baseLogErrConfig(parent, "purgeOldLogFiles()")

	fh := FileHelper{}

	if logOps.LogFileRetentionInDays < 0 {
		return se.SignalNoErrors()
	}

	logDur := time.Duration(logOps.LogFileRetentionInDays*24) * time.Hour
	du := DurationUtility{}
	thresholdTime := du.GetTimeMinusDuration(time.Now(), logDur)

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

// writeFileGroupHeaderToLog - Writes the Job Group
// Header info to the log file. This is a one-time
// operation for each log file. Method InitializeLogFile(..)
// MUST be called before first use of this method.
// The Header is always the first element in the Log.
func (logOps *LogJobGroup) writeFileGroupHeaderToLog(parent []ErrBaseInfo) SpecErr {
	var err error
	var isPanic bool
	var str, stx string

	se := logOps.baseLogErrConfig(parent, "writeFileGroupHeaderToLog()")

	thisParentInfo := se.AddBaseToParentInfo()

	if logOps.FilePtr == nil {
		s := "logOps.FilePtr was not correctly initialized! logOps.FilePtr *os.File pointer is nil!"
		isPanic = true
		return se.New(s, err, isPanic, 301)
	}

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	su := StringUtility{}
	stx = fmt.Sprintf("App Name: %v  AppVersion: %v ", logOps.AppName, logOps.AppVersion)
	str, err = su.StrCenterInStr(stx, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on AppName AppVersion"
		isPanic = true
		return se.New(s, err, isPanic, 302)
	}

	logOps.writeFileStr(str, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	str = fmt.Sprintf("Execution Job Group Command File: %v", logOps.CommandFileName)

	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command File Name"
		isPanic = true
		return se.New(s, err, isPanic, 303)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("Starting Execution of %v Jobs.", logOps.AppNoOfJobs)

	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Number of Jobs Executed"
		isPanic = true
		return se.New(s, err, isPanic, 304)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	dt := DateTimeUtility{}

	str = fmt.Sprintf("Job Execution Start Time: %v", dt.GetDateTimeEverything(logOps.StartTime))

	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Number of Job Execution Start Time"
		isPanic = true
		return se.New(s, err, isPanic, 305)
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	str = fmt.Sprintf("Initial Application Path: %v", logOps.AppPath)

	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("This Log File: %v", logOps.AppLogPathFileName)

	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Base Start Directory: %v", logOps.BaseStartDir)

	logOps.writeTabFileStr(str, 1, thisParentInfo)

	str = fmt.Sprintf("Number Of Old Log Files Deleted: %v", logOps.NoOfLogFilesPurged)

	logOps.writeTabFileStr(str, 1, thisParentInfo)

	if logOps.NoOfLogFilesPurged > 0 {

		str = "!!!!!! Log Files Deleted !!!!!!"
		stx, _ = su.StrCenterInStr(str, logOps.BannerLen)

		logOps.writeTabFileStr(stx, 1, thisParentInfo)
		logOps.writeFileStr(logOps.Banner4, thisParentInfo)
		for i := 0; i < logOps.NoOfLogFilesPurged; i++ {
			fi := logOps.AppLogDirWalkInfo.DeletedFiles[i]

			str = fmt.Sprintf("%v. File Date: %v   File Name: %v",
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
func (logOps *LogJobGroup) WriteJobGroupFooterToLog(parent []ErrBaseInfo) SpecErr {

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

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	su := StringUtility{}

	str = "Completed Job Group Execution"
	stx, err = su.StrCenterInStr(stx, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Job Group Execution Title"
		isPanic = true
		return se.New(s, err, isPanic, 802)
	}

	logOps.writeFileStr(stx, thisParentInfo)

	str = fmt.Sprintf("App Name: %v  AppVersion: %v ", logOps.AppName, logOps.AppVersion)
	stx, err = su.StrCenterInStr(stx, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on AppName AppVersion"
		isPanic = true
		return se.New(s, err, isPanic, 803)
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	str = fmt.Sprintf("Execution Job Group Command File: %v", logOps.CommandFileName)

	stx, err = su.StrCenterInStr(str, logOps.BannerLen)
	if err != nil {
		s := "StrCenterInStr threw error on Command File Name"
		isPanic = true
		return se.New(s, err, isPanic, 804)
	}

	logOps.writeFileStr(stx, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	stx = logOps.LeftTab + fmt.Sprintf("  Number of Jobs Executed: %v ", logOps.NoOfJobsCompleted)

	logOps.writeFileStr(stx, thisParentInfo)

	stx = logOps.LeftTab + fmt.Sprintf("Number of Messages Logged: %v ", logOps.NoOfJobGroupMsgs)

	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner3, thisParentInfo)

	dt := DateTimeUtility{}
	dur := DurationUtility{}
	str = dt.GetDateTimeNanoSecText(logOps.StartTime)
	stx = logOps.LeftTab + fmt.Sprintf("JobGroup   Start Time: %v ", str)
	logOps.writeFileStr(stx, thisParentInfo)

	str = dt.GetDateTimeNanoSecText(logOps.EndTime)
	stx = logOps.LeftTab + fmt.Sprintf("JobGroup     End Time: %v ", str)
	logOps.writeFileStr(stx, thisParentInfo)

	ed, err2 := dur.GetElapsedTime(logOps.StartTime, logOps.EndTime)

	if err2 != nil {
		s := fmt.Sprintf("DateTimeUtility:GetElapsedTime threw error on Start '%v' and End Time '%v'", logOps.StartTime, logOps.EndTime)
		isPanic = true
		return se.New(s, err, isPanic, 805)

	}

	stx = logOps.LeftTab + fmt.Sprintf("JobGroup Elapsed Time: %v ", ed.NanosecStr)
	logOps.writeFileStr(stx, thisParentInfo)

	logOps.writeFileStr(logOps.Banner1, thisParentInfo)
	logOps.writeFileStr(logOps.Banner1, thisParentInfo)

	// Signal Successful Completion
	return se.SignalNoErrors()
}

func (logOps *LogJobGroup) writeTabFileStr(s string, noOfTabs int, parent []ErrBaseInfo) {

	stx := ""

	for i:=0; i < noOfTabs; i++ {
		stx+= logOps.LeftTab
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
