package common

import "fmt"

// PrintXML - Prints Commands generated
// by reading XML file
func PrintXML(cmds CommandsBatch) {

	fmt.Println("=======================================")
	fmt.Println("Command Data from XML File")

	PrintCmdJobsHdr(cmds.CmdJobsHdr)
	PrintCmdJobs(cmds.CmdJobs)

	return
}

// PrintCmdJobsHdr - Prints the Command
// Jobs Header info from CommandsBatch
// structure
func PrintCmdJobsHdr(hdr CommandJobsHdr) {

	fmt.Println("=======================================")
	fmt.Println("CmdJobsHdr")
	fmt.Println("=======================================")
	fmt.Println("Version:", hdr.Version)
	fmt.Println("LogFileRetentionInDays:", hdr.LogFileRetentionInDays)
	fmt.Println("CommandExeDirectory:", hdr.CmdExeDir)
	fmt.Println("LogPath:", hdr.LogPath)
	fmt.Println("LogFileName:", hdr.LogFileName)
	fmt.Println("KillAllJobsOnFirstError:", hdr.KillAllJobsOnFirstError)
	fmt.Println("IanaTimeZone:", hdr.IanaTimeZone)
	fmt.Println("No Of Command Jobs:", hdr.NoOfCmdJobs)

	return
}

// PrintCmdJobs - Prints All Command Jobs
func PrintCmdJobs(cmdJobs CommandJobArray) {
	fmt.Println("=======================================")

	fmt.Println("Printing Command Jobs")
	fmt.Println("=======================================")

	for _, cmdJob := range cmdJobs.CmdJobArray {
		fmt.Println("Display Name:", cmdJob.CmdDisplayName)
		fmt.Println("Command Desc:", cmdJob.CmdDescription)
		fmt.Println("Command Type:", cmdJob.CmdType)
		fmt.Println("ExecuteCmdInDir:", cmdJob.ExeCmdInDir)
		fmt.Println("StartCmdDateTime:", cmdJob.StartCmdDateTime)
		fmt.Println("KillJobsOnExitCodeGreaterThan:", cmdJob.KillJobsOnExitCodeGreaterThan)
		fmt.Println("KillJobsOnExitCodeLessThan:", cmdJob.KillJobsOnExitCodeLessThan)
		fmt.Println("CommandTimeOutInMinutes:", cmdJob.CommandTimeOutInMinutes)
		fmt.Println("ExeCommand:", cmdJob.ExeCommand)
		PrintCmdElements(cmdJob.CmdElements)
	}
}

// PrintCmdElements - Prints Command Elements Array
func PrintCmdElements(CmdElements CommandElementsArray) {
	fmt.Println("---------------------------------------")
	fmt.Println("          Command Elements             ")
	fmt.Println("---------------------------------------")
	for _, cmdElement := range CmdElements.CmdFragments {
		fmt.Println("CmdElement:", cmdElement)
	}

	fmt.Println("=======================================")

}
