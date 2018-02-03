package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*  'loggingmode.go' is located in source code
		repository:

		https://github.com/MikeAustin71/go_cmdrX.git

 */


const (
	srcFileNameXMLCmdJobsData = "xmlcmdjobsdata.go"
	errBlockNoXMLCmdJobsData  = int64(9230610000)
	logOpsTimeFmt = "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
)

// CommandBatch - Xml Root and Parent Element
type CommandBatch struct {
	CmdJobsHdr CommandJobsHdr  `xml:"CommandJobsHeader"`
	CmdJobs    CommandJobArray `xml:"CommandJobs"`
}

func (cmdBatch *CommandBatch) FormatCmdParameters(parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
		SourceFileName: srcFileNameXMLCmdJobsData,
		ParentObjectName: "CommandBatch",
		FuncName: "FormatCmdParameters",
		BaseMessageId: errBlockNoXMLCmdJobsData,
	}

	om := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	om2 := cmdBatch.assembleTimeFormats(om.GetNewParentHistory())

	if om2.IsError() {
		return om2
	}

	om3 := cmdBatch.assembleCmdElements(om.GetNewParentHistory())

	if om3.IsError() {
		return om3
	}

	om.SetSuccessfulCompletionMessage("Completed FormatCmdParameters", 109)

	return om
}

func (cmdBatch *CommandBatch) assembleTimeFormats(parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
		SourceFileName: srcFileNameXMLCmdJobsData,
		ParentObjectName: "CommandBatch",
		FuncName: "assembleTimeFormats",
		BaseMessageId: errBlockNoXMLCmdJobsData,
	}

	om := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	if cmdBatch.CmdJobsHdr.IanaTimeZone == "" {
		cmdBatch.CmdJobsHdr.IanaTimeZone = "Local"
	}

	tzu := TimeZoneUtility{}
	isValidTz, _, _ := tzu.IsValidTimeZone(cmdBatch.CmdJobsHdr.IanaTimeZone)

	if !isValidTz {
		cmdBatch.CmdJobsHdr.IanaTimeZone = "Local"
	}

	cmdBatch.CmdJobsHdr.StdTimeFormat = logOpsTimeFmt

	om.SetSuccessfulCompletionMessage("Completed assembleTimeFormats", 209)

	return om
}

// assembleCmdElements - Assembles
// Command and Command Arguments.
// These command elements are then
// stored in struct CombinedExeCommand
func (cmdBatch *CommandBatch) assembleCmdElements(parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
		SourceFileName: srcFileNameXMLCmdJobsData,
		ParentObjectName: "CommandBatch",
		FuncName: "assembleTimeFormats",
		BaseMessageId: errBlockNoXMLCmdJobsData,
	}

	om := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	cmdBatch.CmdJobsHdr.NoOfCmdJobs = len(cmdBatch.CmdJobs.CmdJobArray)

	lJobs := len(cmdBatch.CmdJobs.CmdJobArray)

	for i := 0; i < lJobs; i++ {
		job := &cmdBatch.CmdJobs.CmdJobArray[i]

		// Set Job No.
		job.CmdJobNo = i + 1

		// sync time zones
		job.IanaTimeZone = cmdBatch.CmdJobsHdr.IanaTimeZone

		// sync time formats
		job.CmdJobTimeFormat = cmdBatch.CmdJobsHdr.StdTimeFormat

		omx := cmdBatch.assembleCmdArgs(job, om.GetNewParentHistory())

		if omx.IsError() {
			return omx
		}

		omy := cmdBatch.assembleInputArgs(job, om.GetNewParentHistory())

		if omy.IsError() {
			return omy
		}

		job.ExeCommand = strings.TrimLeft(strings.TrimRight(job.ExeCommand, " "), " ")

		if len(job.CombinedArguments) > 0 {
			job.CombinedExeCommand =
				job.ExeCommand + " " + job.CombinedArguments
		} else {
			job.CombinedExeCommand = job.ExeCommand
		}

	}

	om.SetSuccessfulCompletionMessage("Finished assembleCmdElements", 309)

	return om
}

// assembleCmdArgs - Processes Command Arguments located in job.CmdArguments.CmdArgs
func (cmdBatch *CommandBatch) assembleCmdArgs(job *CmdJob, parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
		SourceFileName: srcFileNameXMLCmdJobsData,
		ParentObjectName: "CommandBatch",
		FuncName: "assembleCmdArgs",
		BaseMessageId: errBlockNoXMLCmdJobsData,
	}

	om := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	job.CombinedArguments = ""

	lCmdArgs := len(job.CmdArguments.CmdArgs)

	if lCmdArgs == 0 {
		om.SetSuccessfulCompletionMessage("Finished assembleCmdArgs", 408)
		return om
	}

	cmdArgs := ""
	doCmdArgsExist := false

	for i := 0; i < lCmdArgs; i++ {
		job.CmdArguments.CmdArgs[i] = strings.TrimRight(strings.TrimLeft(job.CmdArguments.CmdArgs[i], " "), " ")

		if job.CmdArguments.CmdArgs[i] != "" {

			if i==0 {
				cmdArgs += job.CmdArguments.CmdArgs[i]
			} else {
				cmdArgs += " " + job.CmdArguments.CmdArgs[i]
			}
			doCmdArgsExist = true
		}
	}

	if doCmdArgsExist {
		job.CombinedArguments = cmdArgs
	} else {
		job.CmdArguments.CmdArgs = make([]string, 0, 0)
	}

	om.SetSuccessfulCompletionMessage("Finished assembleCmdArgs", 409)
	return om
}

// assembleInputArgs -
func (cmdBatch *CommandBatch) assembleInputArgs(job *CmdJob, parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
		SourceFileName: srcFileNameXMLCmdJobsData,
		ParentObjectName: "CommandBatch",
		FuncName: "assembleInputArgs",
		BaseMessageId: errBlockNoXMLCmdJobsData,
	}

	om := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	job.CombinedInputArguments = ""

	lCmdInputs := len(job.CmdInputs.InputArgs)

	if lCmdInputs == 0 {
		om.SetSuccessfulCompletionMessage("Finished assembleInputArgs", 508)
		return om
	}

	cmdInputs := ""
	doCmdInputsExist := false

	for i:=0; i < lCmdInputs; i++ {
		job.CmdInputs.InputArgs[i] = strings.TrimRight(strings.TrimLeft(job.CmdInputs.InputArgs[i], " "), " ")

		if job.CmdInputs.InputArgs[i] != "" {

			if i==0 {
				cmdInputs += job.CmdInputs.InputArgs[i]
			} else {
				cmdInputs += " " + job.CmdInputs.InputArgs[i]
			}

			doCmdInputsExist = true
		}

	}

	if doCmdInputsExist {
		job.CombinedInputArguments = cmdInputs
	} else {
		job.CmdInputs.InputArgs = make([]string, 0, 0)
	}

	om.SetSuccessfulCompletionMessage("Finished assembleInputArgs", 509)
	return om
}

// SetBatchStartTime - Sets the time at which jobs in this
// Command Batch began processing.
func (cmdBatch *CommandBatch) SetBatchStartTime(parent []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNameXMLCmdJobsData,
							ParentObjectName: "CommandBatch",
							FuncName: "SetBatchStartTime",
							BaseMessageId: errBlockNoXMLCmdJobsData,
						}

	om := OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)

	cmdBatch.CmdJobsHdr.CmdBatchStartUTC = time.Now().UTC()

	tzu, err := TimeZoneUtility{}.ConvertTz(cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz Error - Failed to convert UTC to local Time Zone. Start UTC: %v. Iana Time Zone: %v", cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.IanaTimeZone)
		om.SetFatalError(s, err, 601)
		return  om
	}

	cmdBatch.CmdJobsHdr.CmdBatchStartTime = tzu.TimeOut

	om.SetSuccessfulCompletionMessage("Finished SetBatchStartTime", 609)

	return om
}

// SetBatchEndTime - Sets the time at which all jobs
// in this Command Batch ended and, in addition,
// computes the elapsed time to complete all jobs in
// this Command Batch.
func (cmdBatch *CommandBatch) SetBatchEndTime(parent []OpsMsgContextInfo) OpsMsgDto {


	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNameXMLCmdJobsData,
							ParentObjectName: "CommandBatch",
							FuncName: "SetBatchEndTime",
							BaseMessageId: errBlockNoXMLCmdJobsData,
	}

	om:= OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)

	cmdBatch.CmdJobsHdr.CmdBatchEndUTC = time.Now().UTC()

	tzu, err := TimeZoneUtility{}.ConvertTz(cmdBatch.CmdJobsHdr.CmdBatchEndUTC, cmdBatch.CmdJobsHdr.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz Error - Failed to convert UTC to local Time Zone. End UTC: %v. Iana Time Zone: %v", cmdBatch.CmdJobsHdr.CmdBatchEndUTC, cmdBatch.CmdJobsHdr.IanaTimeZone)
		om.SetFatalError(s, err, 701)
		return  om
	}

	cmdBatch.CmdJobsHdr.CmdBatchEndTime = tzu.TimeOut

	dutil := DurationUtility{}
	err = dutil.SetStartEndTimes(cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.CmdBatchEndUTC)

	if err != nil {
		s:= fmt.Sprintf("dutil.SetStartEndTimes() Error - Start Time End Time = Elapsed Time Caclulation Failed. Start UTC: %v. End UTC: %v", cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.CmdBatchEndUTC)
		om.SetFatalError(s, err, 702)
		return  om
	}

	cmdBatch.CmdJobsHdr.CmdBatchDuration = dutil.TimeDuration

	elapsedTime, err := dutil.GetYearMthDaysTimeAbbrv()

	if err != nil {
		s:= fmt.Sprintf("dutil.GetYearMthDaysTimeAbbrv() Error - Failed to Duration Get Year Mth Days Time. Start UTC: %v. End UTC: %v", cmdBatch.CmdJobsHdr.CmdBatchStartUTC, cmdBatch.CmdJobsHdr.CmdBatchEndUTC)
		om.SetFatalError(s, err, 703)
		return  om
	}

	cmdBatch.CmdJobsHdr.CmdBatchElapsedTime = elapsedTime.DisplayStr

	om.SetSuccessfulCompletionMessage("Finished SetBatchEndTime", 709)
	return om

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
	CmdInputs									 CommandInputsArray	`xml:"CmdInputs"`
	CombinedInputArguments		 string
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
func (job *CmdJob) SetDelayCmdStartTime(dtf *DateTimeFormatUtility, parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNameXMLCmdJobsData,
							ParentObjectName: "CmdJob",
							FuncName: "SetDelayCmdStartTime",
							BaseMessageId: errBlockNoXMLCmdJobsData,
						}

	//se:= SpecErr{}.InitializeBaseInfo(parent, bi)
	om:= OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	durUtil := DurationUtility{}
	tDto := TimeDto{Years: 3}
	durUtil.SetStartTimeMinusTime(time.Now().UTC(), tDto)

	job.DelayStartCmdDateTimeUTC = durUtil.StartDateTime

	tzuUTC, err := TimeZoneUtility{}.ConvertTz(job.DelayStartCmdDateTimeUTC, job.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz() Error - Time Zone Conversion Failure. Delay Start UTC: %v. Local Time Zone: %v", job.DelayStartCmdDateTimeUTC, job.IanaTimeZone)
		om.SetFatalError(s, err, 1201)
		return om
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
				om.SetSuccessfulCompletionMessage("Finished SetDelayCmdStartTime", 1298)
				return om
			}
		}
	}

	if job.DelayStartCmdDateTime != "" {

		tStart, err := dtf.ParseDateTimeString(job.DelayStartCmdDateTime, "")

		if err != nil {
			s:= fmt.Sprintf("ParseDateTimeString Error. job.DelayStartCmdDateTime: %v", job.DelayStartCmdDateTime)
			om.SetFatalError(s, err, 1205)
			return  om
		}

		tzu := TimeZoneUtility{}

		tStartTz, err := tzu.ReclassifyTimeWithNewTz(tStart, job.IanaTimeZone)

		if err != nil {
			s:= fmt.Sprintf(" REclassifyTimeWithNewTz Error. tStart: %v. Local Time Zone: %v", tStart, job.IanaTimeZone)
			om.SetFatalError(s, err, 1206)
			return  om
		}

		job.DelayStartCmdDateTimeUTC = tStartTz.UTC()
		job.DelayStartCmdDateTimeValue = tStartTz
	}

	om.SetSuccessfulCompletionMessage("Finished SetDelayCmdStartTime", 1299)
	return om
}

// SetCmdJobActualStartTime - Computes and saves the date time
// at which job execution commenced.
func (job *CmdJob) SetCmdJobActualStartTime(parent []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo {
								SourceFileName: srcFileNameXMLCmdJobsData,
								ParentObjectName: "CmdJob",
								FuncName: "SetCmdJobActualStartTime",
								BaseMessageId: errBlockNoXMLCmdJobsData,
						}

	om := OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)

	job.CmdJobStartUTC = time.Now().UTC()

	tzu, err := TimeZoneUtility{}.ConvertTz(job.CmdJobStartUTC, job.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz - Error converting Start Time UTC to local time zone. Start Time %v. Local Time Zone %v.", job.CmdJobStartUTC, job.IanaTimeZone)
		om.SetFatalError(s, err, 1401)
		return  om
	}

	job.CmdJobStartTimeValue = tzu.TimeOut
	om.SetSuccessfulCompletionMessage("Finished SetCmdJobActualStartTime", 1409)
	return om
}

// SetCmdJobActualEndTime - Sets the time at which this job completed processing.
// In addition, this method also computes the elapsed time required to
// complete processing of this command job.

func (job *CmdJob) SetCmdJobActualEndTime(parent []OpsMsgContextInfo) OpsMsgDto {

	//bi := ErrBaseInfo{}.New(srcFileNameXMLCmdJobsData, "CmdJob.SetCmdJobActualEndTime", errBlockNoXMLCmdJobsData)
	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNameXMLCmdJobsData,
							ParentObjectName: "CmdJob",
							FuncName: "SetCmdJobActualEndTime",
							BaseMessageId: errBlockNoXMLCmdJobsData,
					}

	//se:= SpecErr{}.InitializeBaseInfo(parent, bi)
	om := OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)

	job.CmdJobEndUTC = time.Now().UTC()

	tzu, err := TimeZoneUtility{}.ConvertTz(job.CmdJobEndUTC, job.IanaTimeZone)

	if err != nil {
		s:= fmt.Sprintf("TimeZoneUtility{}.ConvertTz - Error converting end time to local time zone. End Time %v. Local Time Zone %v.", job.CmdJobEndUTC, job.IanaTimeZone)
		om.SetFatalError(s, err, 1401)
		return om
	}

	job.CmdJobEndTimeValue = tzu.TimeOut

	dutil := DurationUtility{}

	err = dutil.SetStartEndTimes(job.CmdJobStartUTC, job.CmdJobEndUTC)

	if err != nil {
		s:= fmt.Sprintf("dutil.SetStartEndTimes - Error calculating duration from start time and end time. Start Time %v. End Time %v.", job.CmdJobEndUTC, job.IanaTimeZone)
		om.SetFatalError(s, err, 1402)
		return om
	}

	job.CmdJobDuration = dutil.TimeDuration

	elapsedTime, err := dutil.GetYearMthDaysTimeAbbrv()

	if err != nil {
		s:= fmt.Sprintf("dutil.GetYearMthDaysTimeAbbrv - Duration Calculation Error. Start Time %v. End Time %v.", job.CmdJobEndUTC, job.IanaTimeZone)
		om.SetFatalError(s, err, 1403)
		return  om
	}

	job.CmdJobElapsedTime = elapsedTime.DisplayStr

	om.SetSuccessfulCompletionMessage("Finished InitializeAllContextInfo", 1409)
	return om
}

// CommandArgumentsArray - Holds CmdElement structures
type CommandArgumentsArray struct {
	CmdArgs []string `xml:"CmdArg"`
}

type CommandInputsArray struct {
	InputArgs [] string `xml:"InputArg"`
}
