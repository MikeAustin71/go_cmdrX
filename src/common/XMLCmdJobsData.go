package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	srcFileNameXMLCmdJobsData = "XMLCmdJobsData.go"
	errBlockNoXMLCmdJobsData  = int64(9230610000)
	logOpsTimeFmt = "2006-01-02 15:04:05.000000000 -0700 MST"
)

// CommandBatch - Xml Root and Parent Element
type CommandBatch struct {
	CmdJobsHdr CommandJobsHdr  `xml:"CommandJobsHeader"`
	CmdJobs    CommandJobArray `xml:"CommandJobs"`
}

func (cmdBatch *CommandBatch) FormatCmdParameters() {
	cmdBatch.assembleTimeFormats()
	cmdBatch.assembleCmdElements()
}

func (cmdBatch *CommandBatch) assembleTimeFormats() {

	tzu := TimeZoneUtility{}
	isValidTz, _, _ := tzu.IsValidTimeZone(cmdBatch.CmdJobsHdr.IanaTimeZone)

	if !isValidTz {
		cmdBatch.CmdJobsHdr.IanaTimeZone = "Local"
	}

	cmdBatch.CmdJobsHdr.StdTimeFormat = logOpsTimeFmt

}

// assembleCmdElements - Assembles
// Command and Command Arguments.
// These command elements are then
// stored in struct CombinedExeCommand
func (cmdBatch *CommandBatch) assembleCmdElements() {
	var exCmd string

	cmdBatch.CmdJobsHdr.NoOfCmdJobs = len(cmdBatch.CmdJobs.CmdJobArray)

	lJobs := len(cmdBatch.CmdJobs.CmdJobArray)

	for i := 0; i < lJobs; i++ {
		job := &cmdBatch.CmdJobs.CmdJobArray[i]

		exCmd = ""

		lCmdArgs := len(job.CmdArguments.CmdArgs)

		// Set Job No.
		job.CmdJobNo = i + 1

		// sync time zones
		job.IanaTimeZone = cmdBatch.CmdJobsHdr.IanaTimeZone

		// sync time formats
		job.CmdJobTimeFormat = cmdBatch.CmdJobsHdr.StdTimeFormat

		for k := 0; k < lCmdArgs; k++ {
			job.CmdArguments.CmdArgs[k] = strings.TrimRight(strings.TrimLeft(job.CmdArguments.CmdArgs[k], " "), " ")
			exCmd += job.CmdArguments.CmdArgs[k] + " "
		}

		job.ExeCommand = strings.TrimLeft(strings.TrimRight(job.ExeCommand, " "), " ")

		job.CombinedExeCommand =
			job.ExeCommand + " " + exCmd

		job.CombinedArguments = exCmd
	}

	return
}

// SetBatchStartTime - Sets the time at which jobs in this
// Command Batch began processing.
func (cmdBatch *CommandBatch) SetBatchStartTime(parent []ErrBaseInfo) SpecErr {

	bi := ErrBaseInfo{}.New(srcFileNameXMLCmdJobsData, "cmdBatch.SetBatchStartTime", errBlockNoXMLCmdJobsData)

	se:= SpecErr{}.InitializeBaseInfo(parent, bi)

	cmdBatch.CmdJobsHdr.CmdBatchStartUTC = time.Now().UTC()

	tzu, err := TimeZoneUtility{}.ConvertTz(cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz Error - Failed to convert UTC to local Time Zone. Start UTC: %v. Iana Time Zone: %v", cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.IanaTimeZone)
		return  se.New(s, err, true, 601)
	}

	cmdBatch.CmdJobsHdr.CmdBatchStartTime = tzu.TimeOut

	return se.SignalNoErrors()
}

// SetBatchEndTime - Sets the time at which all jobs
// in this Command Batch ended and, in addition,
// computes the elapsed time to complete all jobs in
// this Command Batch.
func (cmdBatch *CommandBatch) SetBatchEndTime(parent []ErrBaseInfo) SpecErr {

	bi := ErrBaseInfo{}.New(srcFileNameXMLCmdJobsData, "cmdBatch.SetBatchEndTime", errBlockNoXMLCmdJobsData)

	se:= SpecErr{}.InitializeBaseInfo(parent, bi)

	cmdBatch.CmdJobsHdr.CmdBatchEndUTC = time.Now().UTC()

	tzu, err := TimeZoneUtility{}.ConvertTz(cmdBatch.CmdJobsHdr.CmdBatchEndUTC, cmdBatch.CmdJobsHdr.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz Error - Failed to convert UTC to local Time Zone. End UTC: %v. Iana Time Zone: %v", cmdBatch.CmdJobsHdr.CmdBatchEndUTC, cmdBatch.CmdJobsHdr.IanaTimeZone)
		return  se.New(s, err, true, 701)
	}

	cmdBatch.CmdJobsHdr.CmdBatchEndTime = tzu.TimeOut

	dutil := DurationUtility{}
	err = dutil.SetStartEndTimes(cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.CmdBatchEndUTC)

	if err != nil {
		s:= fmt.Sprintf("dutil.SetStartEndTimes() Error - Start Time End Time = Elapsed Time Caclulation Failed. Start UTC: %v. End UTC: %v", cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.CmdBatchEndUTC)
		return  se.New(s, err, true, 702)
	}

	cmdBatch.CmdJobsHdr.CmdBatchDuration = dutil.TimeDuration

	elapsedTime, err := dutil.GetYearMthDaysTimeAbbrv()

	if err != nil {
		s:= fmt.Sprintf("dutil.GetYearMthDaysTimeAbbrv() Error - Failed to Duration Get Year Mth Days Time. Start UTC: %v. End UTC: %v", cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.CmdBatchEndUTC)
		return  se.New(s, err, true, 703)
	}

	cmdBatch.CmdJobsHdr.CmdBatchElapsedTime = elapsedTime.DisplayStr

	return se.SignalNoErrors()

}

// CommandJobsHdr - Holds base info related to
// command jobs
type CommandJobsHdr struct {
	CmdFileVersion          string `xml:"CmdFileVersion"`
	LogFileRetentionInDays  int    `xml:"LogFileRetentionInDays"`
	KillAllJobsOnFirstError bool   `xml:"KillAllJobsOnFirstError"`
	IanaTimeZone            string `xml:"IanaTimeZone"`
	NoOfCmdJobs             int
	StdTimeFormat           string
	CmdBatchStartTime       time.Time
	CmdBatchStartUTC        time.Time
	CmdBatchEndTime         time.Time
	CmdBatchEndUTC          time.Time
	CmdBatchDuration        time.Duration
	CmdBatchElapsedTime     string
}

// CommandJobArray - Holds individual
// CommandJob structs
type CommandJobArray struct {
	CmdJobArray []CmdJob `xml:"CommandJob"`
}

// CmdJob - Command Job information
type CmdJob struct {
	CmdDisplayName             string `xml:"CommandDisplayName"`
	CmdDescription             string `xml:"CommandDescription"`
	CmdJobNo									 int
	CmdType                    string `xml:"CommandType"`
	ExeCmdInDir                string `xml:"ExecuteCmdInDir"`
	DelayCmdStartSeconds       string `xml:"DelayCmdStartSeconds"`
	DelayCmdStartDuration      time.Duration
	DelayStartCmdDateTime      string `xml:"DelayStartCmdDateTime"`
	DelayStartCmdDateTimeValue time.Time
	DelayStartCmdDateTimeUTC   time.Time
	CommandTimeOutInSeconds    float64 `xml:"CmdTimeOutInSeconds"`
	CommandTimeOutDuration     time.Duration
	CombinedExeCommand         string
	ExeCommand                 string `xml:"ExeCommand"`
	CombinedArguments          string
	CmdArguments               CommandArgumentsArray `xml:"CmdArguments"`
	IanaTimeZone               string
	CmdJobTimeFormat           string
	CmdJobStartTimeValue       time.Time
	CmdJobStartUTC             time.Time
	CmdJobEndTimeValue         time.Time
	CmdJobEndUTC               time.Time
	CmdJobDuration             time.Duration
	CmdJobElapsedTime          string
	CmdJobNoOfMsgs						 int
	CmdJobIsCompleted					 bool
}

// SetDelayCmdStartTime - Sets the date time at which the command
// job can begin execution. User has the option to input a delay
// factor expressed in seconds or to input a specific time at which
// the command will begin execution. If no start time is signaled
// by the user, the command will begin execution immediately.
func (job *CmdJob) SetDelayCmdStartTime(dtf *DateTimeFormatUtility, parent []ErrBaseInfo) SpecErr {

	bi := ErrBaseInfo{}.New(srcFileNameXMLCmdJobsData, "CmdJob.SetDelayCmdStartTime", errBlockNoXMLCmdJobsData)

	se:= SpecErr{}.InitializeBaseInfo(parent, bi)

	durUtil := DurationUtility{}
	tDto := TimeDto{Years: 3}
	durUtil.SetStartTimeMinusTime(time.Now().UTC(), tDto)

	job.DelayStartCmdDateTimeUTC = durUtil.StartDateTime

	tzuUTC, err := TimeZoneUtility{}.ConvertTz(job.DelayStartCmdDateTimeUTC, job.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz() Error - Time Zone Conversion Failure. Delay Start UTC: %v. Local Time Zone: %v", job.DelayStartCmdDateTimeUTC, job.IanaTimeZone)
		return  se.New(s, err, true, 1201)
	}

	job.DelayStartCmdDateTimeValue = tzuUTC.TimeOut

	if job.DelayCmdStartSeconds != "" {

		dSecs, err := strconv.ParseInt(job.DelayCmdStartSeconds, 10, 64)

		if err == nil {

			dur := durUtil.GetDurationFromSeconds(dSecs)
			job.DelayStartCmdDateTimeUTC = time.Now().UTC().Add(dur)

			tzu, err := TimeZoneUtility{}.ConvertTz(job.DelayStartCmdDateTimeUTC, job.IanaTimeZone)

			if err == nil {
				job.DelayStartCmdDateTimeValue = tzu.TimeOut
				return se.SignalNoErrors()
			}
		}
	}

	if job.DelayStartCmdDateTime != "" {

		tStart, err := dtf.ParseDateTimeString(job.DelayStartCmdDateTime, "")

		if err != nil {
			s:= fmt.Sprintf("ParseDateTimeString Error. job.DelayStartCmdDateTime: %v", job.DelayStartCmdDateTime)
			return  se.New(s, err, true, 1205)
		}

		tzu := TimeZoneUtility{}

		tStartTz, err := tzu.ReclassifyTimeWithNewTz(tStart, job.IanaTimeZone)

		if err != nil {
			s:= fmt.Sprintf(" REclassifyTimeWithNewTz Error. tStart: %v. Local Time Zone: %v", tStart, job.IanaTimeZone)
			return  se.New(s, err, true, 1206)
		}

		job.DelayStartCmdDateTimeUTC = tStartTz.UTC()
		job.DelayStartCmdDateTimeValue = tStartTz
	}

	return se.SignalNoErrors()
}

// SetCmdJobActualStartTime - Computes and saves the date time
// at which job execution commenced.
func (job *CmdJob) SetCmdJobActualStartTime(parent []ErrBaseInfo) SpecErr {

	bi := ErrBaseInfo{}.New(srcFileNameXMLCmdJobsData, "CmdJob.SetDelayCmdStartTime", errBlockNoXMLCmdJobsData)

	se:= SpecErr{}.InitializeBaseInfo(parent, bi)

	job.CmdJobStartUTC = time.Now().UTC()
	tzu, err := TimeZoneUtility{}.ConvertTz(job.CmdJobStartUTC, job.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz - Error converting Start Time UTC to local time zone. Start Time %v. Local Time Zone %v.", job.CmdJobStartUTC, job.IanaTimeZone)
		return  se.New(s, err, true, 1401)
	}

	job.CmdJobStartTimeValue = tzu.TimeOut

	return se.SignalNoErrors()
}

// SetCmdJobActualEndTime - Sets the time at which this job completed processing.
// In addition, this method also computes the elapsed time required to
// complete processing of this command job.

func (job *CmdJob) SetCmdJobActualEndTime(parent []ErrBaseInfo) SpecErr {

	bi := ErrBaseInfo{}.New(srcFileNameXMLCmdJobsData, "CmdJob.SetCmdJobActualEndTime", errBlockNoXMLCmdJobsData)

	se:= SpecErr{}.InitializeBaseInfo(parent, bi)

	job.CmdJobEndUTC = time.Now().UTC()

	tzu, err := TimeZoneUtility{}.ConvertTz(job.CmdJobEndUTC, job.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz - Error converting end time to local time zone. End Time %v. Local Time Zone %v.", job.CmdJobEndUTC, job.IanaTimeZone)
		return  se.New(s, err, true, 1401)
	}

	job.CmdJobEndTimeValue = tzu.TimeOut

	dutil := DurationUtility{}

	err = dutil.SetStartEndTimes(job.CmdJobStartUTC, job.CmdJobEndUTC)

	if err != nil {
		s:= fmt.Sprintf("dutil.SetStartEndTimes - Error calculating duration from start time and end time. Start Time %v. End Time %v.", job.CmdJobEndUTC, job.IanaTimeZone)
		return  se.New(s, err, true, 1402)
	}

	job.CmdJobDuration = dutil.TimeDuration

	elapsedTime, err := dutil.GetYearMthDaysTimeAbbrv()

	if err != nil {
		s:= fmt.Sprintf("dutil.GetYearMthDaysTimeAbbrv - Duration Calculation Error. Start Time %v. End Time %v.", job.CmdJobEndUTC, job.IanaTimeZone)
		return  se.New(s, err, true, 1403)
	}

	job.CmdJobElapsedTime = elapsedTime.DisplayStr

	return se.SignalNoErrors()
}

// CommandArgumentsArray - Holds CmdElement structures
type CommandArgumentsArray struct {
	CmdArgs []string `xml:"CmdArg"`
}
