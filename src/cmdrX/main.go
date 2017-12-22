package main

import (
	"MikeAustin71/logopsgo/common"
	"fmt"
	"time"
)

const (
	srcFileNameLogOpsMain = "main.go"
	errBlockNoLogOpsMain  = int64(1000000)
	cmdPathFileName = "D:/go/work/src/go_cmdrX/src/cmdrX/cmdrXCmds.xml"
	appPathFileName = "D:/go/work/src/go_cmdrX/src/cmdrX/cmdrX.exe"
	appLogPathOnly = "D:/go/work/src/go_cmdrX/src/cmdrX/cmdrXLog/alog.log"
	appBanner1 = "===================================================================="
)

func main() {


	parent := common.OpsMsgContextInfo{
		SourceFileName: srcFileNameLogOpsMain,
		ParentObjectName: "",
		FuncName: "main",
		BaseMessageId: errBlockNoLogOpsMain,
	}

	om := common.OpsMsgDto{}.InitializeWithMessageContext(parent)

	fh := common.FileHelper{}

	currDir, err := fh.GetCurrentDir()

	if err!=nil {
		om.SetFatalError("fh.GetCurrentDir() FAILED!\n", err, 1)
		om.PrintToConsole()
		return
	}

	su := common.StringUtility{}
	lBanner1 := len(appBanner1)
	fmt.Println("\n\n"+ appBanner1)
	strx, _ := su.StrCenterInStr("CmdrX Now Running Commands", lBanner1 - 2)
	fmt.Println("=" + strx + "=")
	fmt.Println(appBanner1)
	fmt.Println("Current Directory: ", currDir)

	exeFilePath, err := fh.GetExecutablePathFileName()

	if err!=nil {
		om.SetFatalError("fh.GetExecutablePathFileName() FAILED!\n", err, 2)
		om.PrintToConsole()
		return
	}

	fmt.Println("Executable Directory: ", exeFilePath)

	parentHistory := []common.OpsMsgContextInfo{parent}

	lg := common.LogJobGroup{}

	cmds, om3 := startUp(&lg, parentHistory)

	if om3.IsFatalError() {
		panic(om3)
	}

	for _, cmdJob := range cmds.CmdJobs.CmdJobArray {

		om4 := executeJob(&cmdJob, &lg, parentHistory)

		if cmdJob.CmdJobIsCompleted && !om4.IsError() {
			lg.NoOfJobsCompleted++
		}

		lg.NoOfJobGroupMsgs += cmdJob.CmdJobNoOfMsgs

		if om4.IsError() && lg.KillAllJobsOnFirstError {
			om4.LogMsgToFile(lg.FilePtr)
			doLogWrapUp(&lg, cmds, parentHistory)
			panic(om4)
		}
	}

	doLogWrapUp(&lg, cmds, parentHistory)
}

func doLogWrapUp(lg *common.LogJobGroup, cmds common.CommandBatch, parent []common.OpsMsgContextInfo) common.OpsMsgDto {

	msgCtx := common.OpsMsgContextInfo{
							SourceFileName:"main.go",
							ParentObjectName: "main()",
							FuncName:" doLogWrapUp",
							BaseMessageId: errBlockNoLogOpsMain,
						}

	om1 := common.OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)

	// Closes lg.FilePtr
	om2 := lg.WriteJobGroupFooterToLog(cmds, om1.GetNewParentHistory())

	if om2.IsFatalError(){
		om2.PrintToConsole()
		return om2
	}

	fmt.Println(appBanner1)
	fmt.Println("See Log File:")
	fmt.Println(lg.AppLogPathFileName)

	fmt.Println(appBanner1)

	return om2
}

func startUp(lg *common.LogJobGroup,parent []common.OpsMsgContextInfo) (common.CommandBatch, common.OpsMsgDto) {

	om := baseLogMsgConfigMain(parent, "startUp")
	parentHistory := om.GetNewParentHistory()
	parms := common.StartupParameters{}

	appFileParms, om2 := parms.AssembleAppPath(appPathFileName, parentHistory)

	if om2.IsFatalError() {
		panic(om2)
	}

	cmdFileParms, om3 := parms.AssembleCmdPath(cmdPathFileName, parentHistory)

	if om3.IsFatalError() {
		panic(om3)
	}

	appLogPathParms, om4 := parms.AssembleLogPath(appLogPathOnly, parentHistory)

	if om4.IsFatalError() {
		panic(om4)
	}

	cmds, om5 := common.ParseXML(cmdFileParms.AbsolutePathFileName, parentHistory)

	if om5.IsFatalError() {
		panic(om5)
	}

	om6 := cmds.FormatCmdParameters(parentHistory)

	if om6.IsFatalError() {
		panic(om6)
	}


	om7 := cmds.SetBatchStartTime(parentHistory)
	if om7.IsFatalError() {
		panic(om7)
	}

	parms.IanaTimeZone = 	cmds.CmdJobsHdr.IanaTimeZone
	parms.KillAllJobsOnFirstError = cmds.CmdJobsHdr.KillAllJobsOnFirstError
	parms.StartTimeUTC = cmds.CmdJobsHdr.CmdBatchStartUTC
	parms.StartTime = cmds.CmdJobsHdr.CmdBatchStartTime
	dtf := common.DateTimeFormatUtility{}
	dtf.CreateAllFormatsInMemory()


	parms.AppVersion = "2.0.0"
	parms.LogMode = common.LogVERBOSE
	parms.AppLogPath = appLogPathParms.AbsolutePath
	parms.AppName = appFileParms.FileName
	parms.AppExeFileName = appFileParms.FileNameExt
	parms.AppPath = appFileParms.AbsolutePath
	parms.BaseStartDir = appFileParms.AbsolutePath
	parms.CommandFileName = cmdFileParms.FileNameExt
	parms.NoOfJobs = cmds.CmdJobsHdr.NoOfCmdJobs
	parms.Dtfmt = &dtf

	om8 := lg.New(parms, parentHistory)

	if om8.IsFatalError() {
		panic(om8)
	}

	dur := common.DurationUtility{}

	time.Sleep(dur.GetDurationFromSeconds(10))

	om9 := cmds.SetBatchEndTime(parentHistory)

	if om9.IsFatalError() {
		panic(om9)
	}

	om.SetSuccessfulCompletionMessage("Finished startUp()", 79)
	return cmds, om
}


func executeJob(job *common.CmdJob, logOps *common.LogJobGroup, parent []common.OpsMsgContextInfo) common.OpsMsgDto {

	om := baseLogMsgConfigMain(parent, "executeJob")

	thisParentInfo := om.GetNewParentHistory()

	job.SetCmdJobActualStartTime(thisParentInfo)

	om2 := logOps.WriteCmdJobHeaderToLog(job, thisParentInfo)

	if om2.IsFatalError() {
		logOps.WriteOpsMsgToLog(om2, job, thisParentInfo)
		om2.PrintToConsole()
		return om2
	}

	executeJobCommand(job,logOps,om.GetNewParentHistory())

	executeJobCommand(job,logOps, om.GetNewParentHistory())

	logOps.WriteCmdJobFooterToLog(job, thisParentInfo)

	return om.SignalSuccessfulCompletion(619)
}

func executeJobCommand(job *common.CmdJob, logOps *common.LogJobGroup, parent []common.OpsMsgContextInfo) common.OpsMsgDto {

	om := baseLogMsgConfigMain(parent, "executeJobCommand")

	time.Sleep(time.Duration(5) * time.Second)

	s := fmt.Sprintf("Completed Job: %v. No Errors!", job.CmdDisplayName)
	opsMsg := common.OpsMsgDto{}

	job.CmdJobNoOfMsgs++

	opsMsg.SetTimeZone(job.IanaTimeZone)

	opsMsg.SetInfoMessage(s, int64((job.CmdJobNo * 10000000) + job.CmdJobNoOfMsgs))

	om1 := logOps.WriteOpsMsgToLog(opsMsg, job, om.GetNewParentHistory())

	if om1.IsFatalError() {
		panic(om1)
	}

	om2 := job.SetCmdJobActualEndTime(om.GetNewParentHistory())

	if om2.IsFatalError() {
		panic(om2)
	}

	om.SetSuccessfulCompletionMessage("Completed executeJobCommand", 629)

	return om
}

// baseLogErrConfig - Used internally by LogJobGroup
// methods to set up error messages.
func baseLogMsgConfigMain(parent []common.OpsMsgContextInfo, funcName string) common.OpsMsgDto {

	opsContext := common.OpsMsgContextInfo{
									SourceFileName: srcFileNameLogOpsMain,
									ParentObjectName: "",
									FuncName: funcName,
									BaseMessageId: errBlockNoLogOpsMain,
	}


	return common.OpsMsgDto{}.InitializeAllContextInfo(parent, opsContext)
}

