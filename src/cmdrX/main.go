package main

import (
	"errors"
	"fmt"
	"go_cmdrX/src/common"
	"time"
)

const (
	srcFileNameCmdrX = "main.go"
	errBlockNoCmdrX  = int64(100000)
)

func main() {

	fh := common.FileHelper{}
	xmlFilePathName := "D:\\go\\work\\src\\bitbucket.org\\xmlgo\\03_ReadXmlCmds\\common\\cmdrXCmds.xml"

	s, err := fh.MakeAbsolutePath(xmlFilePathName)

	if err != nil {
		panic(fmt.Errorf("main() - FileHelper{}.GetPathFileNameElements Failed! Xml File: %v - Error: %v", xmlFilePathName, err.Error()))
	}

	xmlFile := fh.AdjustPathSlash(s)

	parent := common.ErrBaseInfo{}.GetNewParentInfo(srcFileNameCmdrX, "main", errBlockNoCmdrX)
	cmds, se := common.ParseXML(xmlFile, parent)

	if se.IsErr {
		panic(se)
	}

	pathElements, err := fh.GetPathFileNameElements(xmlFile)

	if err != nil {
		panic(fmt.Errorf("main() - FileHelper{}.GetPathFileNameElements Failed! Xml File: %v - Error: %v", xmlFile, err.Error()))
	}

	if !fh.AbsolutePathIsPopulated {
		panic(fmt.Errorf("main() - FileHelper{}.GetPathFileNameElements Failed! Xml File: %v - Error: Could Not Locate Absolute Path", xmlFile))
	}

	noOfCmdJobs := len(cmds.CmdJobs.CmdJobArray)

	if noOfCmdJobs < 1 {
		panic(errors.New("main() - No command jobs were found!"))
	}

	dtf := common.DateTimeFormatUtility{}
	dtf.CreateAllFormatsInMemory()
	parms := common.StartupParameters{}

	parms.StartTime = time.Now()
	parms.AppVersion = cmds.CmdJobsHdr.CmdFileVersion
	parms.LogMode = common.LogVERBOSE
	parms.AppLogFileName = "cmdrXRun"
	parms.AppLogPath = "./cmdrX"
	parms.AppName = "cmdrX"
	parms.AppExeFileName = "cmdrX.exe"
	parms.NoOfJobs = noOfCmdJobs
	parms.CommandFileName = pathElements.FileNameExt
	parms.AppPath = pathElements.AbsolutePath
	parms.Dtfmt = &dtf
	parms.KillAllJobsOnFirstError = cmds.CmdJobsHdr.KillAllJobsOnFirstError
	parms.IanaTimeZone = cmds.CmdJobsHdr.IanaTimeZone
	parms.LogFileRetentionInDays = cmds.CmdJobsHdr.LogFileRetentionInDays

	parms.BaseStartDir = fh.AbsolutePath

	logOps := common.LogJobGroup{}

	se = logOps.New(parms, parent)

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

func executeJob(job common.CmdJob, logOps *common.LogJobGroup, parent []common.ErrBaseInfo) common.SpecErr {

	bi := common.ErrBaseInfo{}.New(srcFileNameCmdrX, "executeJob", errBlockNoCmdrX)

	se := common.SpecErr{}.InitializeBaseInfo(parent, bi)

	return se.SignalNoErrors()
}
