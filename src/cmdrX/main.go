package main

import (
"go_cmdrX/src/common"
"fmt"
	"time"
	"errors"
)

const (
	mainSrcFileName = "main.go"
	mainErrBlockNo  = int64(1000)
)
func main()  {
	xmlFile := "./cmdrXCmds.xml"

	fh, err := common.FileHelper{}.GetPathFileNameElements(xmlFile)

	if err != nil {
		panic(fmt.Errorf("main() - FileHelper{}.GetPathFileNameElements Failed! Xml File: %v - Error: %v",	xmlFile, err.Error()))
	}

	if !fh.AbsolutePathIsPopulated{
		panic(fmt.Errorf("main() - FileHelper{}.GetPathFileNameElements Failed! Xml File: %v - Error: Could Not Locate Absolute Path",	xmlFile))
	}

	cmds := common.ParseXML(xmlFile)

	noOfCmdJobs :=  len(cmds.CmdJobs.CmdJobArray)

	if noOfCmdJobs < 1 {
		panic(errors.New("main() - No command jobs were found!"))
	}

	common.AssembleCmdElements(&cmds)

	dtf := common.DateTimeFormatUtility{}
	dtf.CreateAllFormatsInMemory()
	parms := common.StartupParameters{}

	parms.StartTime = time.Now()
	parms.AppVersion = cmds.CmdJobsHdr.Version
	parms.LogMode = common.LogVERBOSE
	parms.AppLogFileName = "cmdrXRun"
	parms.AppLogPath = "./cmdrX"
	parms.AppName = "cmdrX"
	parms.AppExeFileName = "cmdrX.exe"
	parms.NoOfJobs = noOfCmdJobs
	parms.CommandFileName = fh.FileNameExt
	parms.AppPath = fh.AbsolutePath
	parms.Dtfmt = &dtf
	parms.KillAllJobsOnFirstError = cmds.CmdJobsHdr.KillAllJobsOnFirstError
	parms.IanaTimeZone = cmds.CmdJobsHdr.IanaTimeZone
	parms.LogFileRetentionInDays = cmds.CmdJobsHdr.LogFileRetentionInDays

	parms.BaseStartDir = fh.AbsolutePath

	parent := common.ErrBaseInfo{}.GetNewParentInfo(mainSrcFileName, "main", mainErrBlockNo)

	logOps := common.LogJobGroup{}

	se := logOps.New(parms, parent)

	if se.IsErr {
		panic(se.Error)
	}

	for _, cmdJob := range cmds.CmdJobs.CmdJobArray {

		se2 := executeJob(cmdJob, &logOps, parent)

		if se2.IsErr && logOps.KillAllJobsOnFirstError {
			panic(se2)
		}
	}

}

func executeJob( job common.CmdJob, logOps *common.LogJobGroup, parent []common.ErrBaseInfo ) common.SpecErr {
	se := common.SpecErr{}


	return se.SignalNoErrors()
}

