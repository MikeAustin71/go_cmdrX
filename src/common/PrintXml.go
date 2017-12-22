package common

/*  'PrintXml.go' is located in source code
		repository:

		https://github.com/MikeAustin71/go_cmdrX.git

*/

import "fmt"

const (
	srcFileNamePrintXml = "PrintXml.go"
	errBlockNoPrintXml  = int64(11230000)
)

// PrintXML - Prints Commands generated
// by reading XML file
func PrintXML(cmds CommandBatch) {

	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNamePrintXml,
							ParentObjectName: "",
							FuncName: "PrintXML",
							BaseMessageId: errBlockNoPrintXml,
						}

  parentHistory := []OpsMsgContextInfo{msgCtx}

	fmt.Println("=======================================")
	fmt.Println("Command Data from XML File")

	om1 := PrintCmdJobsHdr(cmds.CmdJobsHdr, parentHistory)

	if om1.IsFatalError() {
		panic(om1)
	}

	om2 := PrintCmdJobs(cmds.CmdJobs, parentHistory)

	if om2.IsFatalError() {
		panic(om2)
	}

	return
}

// PrintCmdJobsHdr - Prints the Command
// Jobs Header info from CommandBatch
// structure
func PrintCmdJobsHdr(hdr CommandJobsHdr, parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNamePrintXml,
							ParentObjectName: "",
							FuncName: "PrintCmdJobsHdr",
							BaseMessageId: errBlockNoPrintXml,
						}

	om := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	fmt.Println("=======================================")
	fmt.Println("CmdJobsHdr")
	fmt.Println("=======================================")
	fmt.Println("Command File Version:", hdr.CmdFileVersion)
	fmt.Println("LogFileRetentionInDays:", hdr.LogFileRetentionInDays)
	fmt.Println("KillAllJobsOnFirstError:", hdr.KillAllJobsOnFirstError)
	fmt.Println("IanaTimeZone:", hdr.IanaTimeZone)
	fmt.Println("No Of Command Jobs:", hdr.NoOfCmdJobs)
	fmt.Println("Command Batch Start Time: ", hdr.CmdBatchStartTime.Format(hdr.StdTimeFormat))

	om.SetSuccessfulCompletionMessage("", 9)
	return om
}

// PrintCmdJobs - Prints All Command Jobs
func PrintCmdJobs(cmdJobs CommandJobArray, parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNamePrintXml,
							ParentObjectName: "",
							FuncName: "PrintCmdJobs",
							BaseMessageId: errBlockNoPrintXml,
						}

	om := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	thisParentHistory := om.GetNewParentHistory()

	fmt.Println("=======================================")

	fmt.Println("Printing Command Jobs")
	fmt.Println("=======================================")

	lJobs := len(cmdJobs.CmdJobArray)

	for i := 0; i < lJobs; i++ {

		om2 := cmdJobs.CmdJobArray[i].SetCmdJobActualStartTime(thisParentHistory)

		if om2.IsFatalError() {
			panic(om2)
		}

		fmt.Println("Display Name:", cmdJobs.CmdJobArray[i].CmdDisplayName)
		fmt.Println("Command Desc:", cmdJobs.CmdJobArray[i].CmdDescription)
		fmt.Println("Command Type:", cmdJobs.CmdJobArray[i].CmdType)
		fmt.Println("ExecuteCmdInDir:", cmdJobs.CmdJobArray[i].ExeCmdInDir)
		fmt.Println("DelayCmdStartSeconds:", cmdJobs.CmdJobArray[i].DelayCmdStartSeconds)
		fmt.Println("DelayStartCmdDateTime:", cmdJobs.CmdJobArray[i].DelayStartCmdDateTime)
		fmt.Println("CommandTimeOutInSeconds:", cmdJobs.CmdJobArray[i].CommandTimeOutInSeconds)
		fmt.Println("Job Start Time:", cmdJobs.CmdJobArray[i].CmdJobStartTimeValue.Format(cmdJobs.CmdJobArray[i].CmdJobTimeFormat))
		fmt.Println("Job Time Zone:", cmdJobs.CmdJobArray[i].IanaTimeZone)
		fmt.Println("+++++++++++++++++++++++++++++++++++++++")
		fmt.Println("      Combined Exe Command             ")
		fmt.Println("+++++++++++++++++++++++++++++++++++++++")
		fmt.Println("Combined Exe Command:", cmdJobs.CmdJobArray[i].CombinedExeCommand)
		fmt.Println("---------------------------------------")
		fmt.Println("           Exe Command                 ")
		fmt.Println("---------------------------------------")
		fmt.Println("ExeCommand:", cmdJobs.CmdJobArray[i].ExeCommand)
		PrintCmdElements(cmdJobs.CmdJobArray[i].CmdArguments, thisParentHistory)
	}

	om.SetSuccessfulCompletionMessage("", 99)
	return om
}

// PrintCmdElements - Prints Command Elements Array
func PrintCmdElements(CmdArguments CommandArgumentsArray, parentHistory []OpsMsgContextInfo) OpsMsgDto {

	msgCtx := OpsMsgContextInfo{
							SourceFileName: srcFileNamePrintXml,
							ParentObjectName: "",
							FuncName: "PrintCmdJobs",
							BaseMessageId: errBlockNoPrintXml,
						}

	om := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgCtx)

	fmt.Println("---------------------------------------")
	fmt.Println("         Command Arguments             ")
	fmt.Println("---------------------------------------")
	for _, cmdArg := range CmdArguments.CmdArgs {
		fmt.Println("Cmd Argument:", cmdArg)
	}

	fmt.Println("=======================================")


	om.SetSuccessfulCompletionMessage("", 999)

	return om
}
